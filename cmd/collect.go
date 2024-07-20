package cmd

import (
	"bufio"
	"errors"
	"github.com/spf13/cobra"
	"github.com/tarm/goserial"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	log_parser "zmk-heatmap/pkg/collector"
	"zmk-heatmap/pkg/heatmap"
	"zmk-heatmap/pkg/keymap"
)

var (
	keyboardParam string
	outputParam   string
	keymapParam   string
)

func init() {
	rootCmd.AddCommand(collectCmd)

	collectCmd.Flags().StringVarP(&keyboardParam, "keyboard", "k", "auto", "e.g. /dev/tty.usbmodem144001")
	collectCmd.Flags().StringVarP(&outputParam, "output", "o", "heatmap.json", "e.g. ~/heatmap.json")
	collectCmd.Flags().StringVarP(&keymapParam, "keymap", "m", "keymap.yaml", "e.g. ~/keymap.yaml")
	collectCmd.MarkFlagRequired("keymap")
}

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Process the keystrokes from your keyboard",
	Long:  "Process the keystrokes from your keyboard and save the aggregated result in a file that can be used to generate the heatmap",
	Run: func(cmd *cobra.Command, args []string) {
		// Remove the timestamp from the log messages
		log.SetFlags(log.Flags() &^ (log.Ldate))

		keyboardPath, err := findKeyboard(keyboardParam)
		if err != nil {
			log.Fatalln(err.Error())
		}

		_, err = os.Stat(keyboardPath)
		if err != nil {
			log.Fatalln("Cannot connect to the keyboardPath:", err)
		}

		log.Println("Connecting to the keyboardPath at:", keyboardPath)

		c := &serial.Config{Name: keyboardPath, Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatalln(err)
		}

		heatmap, err := loadHeatMap(outputParam)
		if err != nil {
			log.Fatalln(err)
		}
		if heatmap.GetPressCount() > 0 {
			log.Println("Loaded", outputParam, "with", heatmap.GetPressCount(), "key presses")
		}

		keymapFile := keymapParam
		keymapp, err := keymap.Load(keymapFile)
		if err != nil {
			log.Fatalln("Cannot load the keymap:", err)
		}
		log.Println("Loading keymap", keymapFile)

		parser := log_parser.NewZmkLogParser(keymapp)

		// Store the collected keystrokes every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		go storeKeyStrokes(ticker, heatmap)

		// Start the key scanner
		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			_ = parser.Parse(scanner.Text(), heatmap)
		}
		if scanner.Err() != nil {
			log.Fatal(err)
		}
	},
}

func loadHeatMap(heatmapFile string) (heatmap_ *heatmap.Heatmap, err error) {
	if _, err := os.Stat(heatmapFile); os.IsNotExist(err) {
		return heatmap.New(), nil
	}

	return heatmap.Load(heatmapFile)
}

// Scan /dev/tty* for possible keyboards and returns the path for the keyboard if one and only one is connected
// returns an error otherwise
func findKeyboard(keyboardPath string) (k string, err error) {
	if keyboardPath != "" && keyboardPath != "auto" {
		return keyboardPath, nil
	}

	matches, _ := filepath.Glob("/dev/tty.usbmodem*")
	if len(matches) == 0 {
		return k, errors.New("No keyboardPath found. Please make sure that the keyboardPath is connected via USB and that the firmware has USB_LOGGING enabled. See: https://zmk.dev/docs/development/usb-logging")
	}

	if len(matches) > 1 {
		return k, errors.New("Multiple keyboards found: " + (strings.Join(matches, ", ")) + ". Please specify the wanted keyboardPath with the -k flag.")
	}

	return matches[0], nil
}

func storeKeyStrokes(ticker *time.Ticker, heatmap *heatmap.Heatmap) {
	for {
		select {
		case <-ticker.C:
			log.Println("Collected", heatmap.GetPressCount(), "key presses")
			heatmap.Save(outputParam)

			//case <- quit:
			//	ticker.Stop()
			//	return
		}
	}
}
