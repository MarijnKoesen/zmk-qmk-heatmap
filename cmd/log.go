package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tarm/goserial"
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

		var logFile *os.File
		if outputParam != "" {
			logFile, err = os.OpenFile(outputParam, os.O_WRONLY|os.O_CREATE, 0o644)
			if err != nil {
				log.Fatalln("cannot open output file: " + err.Error())
			}

			// logFile = bufio.NewWriter(file)
			defer logFile.Sync()
		}

		c := &serial.Config{Name: keyboardPath, Baud: 115200}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatalln(err)
		}

		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			if outputParam != "" {
				_, err = logFile.WriteString(scanner.Text() + "\n")
				if err != nil {
					log.Fatalln("cannot write to file ", outputParam, ":"+err.Error())
				}

			} else {
				log.Println(scanner.Text())
			}
		}
		if scanner.Err() != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringVarP(&keyboardParam, "keyboard", "k", "auto", "e.g. /dev/tty.usbmodem144001")
	logCmd.Flags().StringVarP(&outputParam, "output", "o", "", "e.g. testdata/zmk/my.log")
}
