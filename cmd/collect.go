package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	log_parser "zmk-heatmap/pkg/collector"
	"zmk-heatmap/pkg/keymap"

	"github.com/tarm/goserial"
)

var keyboardParam string
var outputParam string

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect the keystrokes from your keyboard",
	Long:  "Collect the keystrokes from your keyboard and save the aggregated result in a file that can be used to generate the heatmap",
	Run: func(cmd *cobra.Command, args []string) {
		// Remove the timestamp from the log messages
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

		keyboardPath, err := findKeyboard(keyboardParam)
		if err != nil {
			log.Fatalln(err.Error())
		}

		_, err = os.Stat(keyboardPath)
		if err != nil {
			log.Fatalln("Cannot connect to the keyboardPath:", err)
		}

		fmt.Println("Connecting to the keyboardPath at:", keyboardPath)

		c := &serial.Config{Name: keyboardPath, Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatalln(err)
		}

		p, err := log_parser.LoadParser(outputParam)
		if err != nil {
			keymapp, _ := keymap.Load("testdata/keymap.yaml")
			p = log_parser.NewParser(keymapp)
		}

		// Store the collected keystrokes every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		go storeKeyStrokes(ticker, p)

		// Start the key scanner
		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			_ = p.Parse(scanner.Text())
		}
		if scanner.Err() != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(collectCmd)

	collectCmd.Flags().StringVarP(&keyboardParam, "keyboard", "k", "auto", "e.g. /dev/tty.usbmodem144001")
	collectCmd.Flags().StringVarP(&outputParam, "output", "o", "heatmap.json", "e.g. ~/heatmap.json")
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

func storeKeyStrokes(ticker *time.Ticker, parser *log_parser.Parser) {
	for {
		select {
		case <-ticker.C:
			json, err := json.Marshal(parser)
			if err != nil {
				fmt.Println(err)
			}
			_ = os.WriteFile(outputParam, json, 0644)

			//case <- quit:
			//	ticker.Stop()
			//	return
		}
	}
}
