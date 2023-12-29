package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"github.com/tarm/goserial"
	"log"
	"os"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Output the raw keyboard log",
	Long:  "Output the raw keyboard log",
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

		log.Println("Connecting to the keyboardPath at:", keyboardPath)

		c := &serial.Config{Name: keyboardPath, Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatalln(err)
		}

		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
		if scanner.Err() != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringVarP(&keyboardParam, "keyboard", "k", "auto", "e.g. /dev/tty.usbmodem144001")
}
