package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"zmk-heatmap/pkg/keymap"

	"github.com/tarm/goserial"

	log_parser "zmk-heatmap/pkg/log-parser"
)

const save_file = "heatmap_data.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<keyboard-serial-device>")
		os.Exit(1)
	}

	_, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot conncect to the keyboard: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Connecting to the keyboard on:", os.Args[1])

	c := &serial.Config{Name: os.Args[1], Baud: 115200}
	s, err := serial.OpenPort(c)

	if err != nil {
		fmt.Println(err)
	}

	p, err := log_parser.LoadParser(save_file)
	if err != nil {
		keymapp, _ := keymap.Load("testdata/keymap.yaml")
		p = log_parser.NewParser(keymapp)
	}

	ticker := time.NewTicker(5 * time.Second)
	go oul(ticker, p)

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		p.Parse(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

func oul(ticker *time.Ticker, parser *log_parser.Parser) {
	for {
		select {
		case <-ticker.C:
			json, err := json.Marshal(parser)
			if err != nil {
				fmt.Println(err)
			}
			_ = os.WriteFile(save_file, json, 0644)

			//case <- quit:
			//	ticker.Stop()
			//	return
		}
	}
}
