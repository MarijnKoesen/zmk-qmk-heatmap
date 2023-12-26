package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	color2 "image/color"
	"log"
	"sort"
	log_parser "zmk-heatmap/pkg/collector"
	"zmk-heatmap/pkg/heatmap"
)

var inputParam string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the heatmap",
	Long:  "Generate the heatmap",
	Run: func(cmd *cobra.Command, args []string) {
		// Remove the timestamp from the log messages
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

		p, err := log_parser.LoadParser(inputParam)
		if err != nil {
			log.Fatalln("Cannot read heatmap:", err)
		}

		sort.Slice(p.KeyPresses, func(i, j int) bool { return p.KeyPresses[i].Count < p.KeyPresses[j].Count })
		maxPress := p.KeyPresses[len(p.KeyPresses)-1].Count
		sort.Slice(p.ComboPresses, func(i, j int) bool { return p.ComboPresses[i].Count < p.ComboPresses[j].Count })
		maxCombo := p.ComboPresses[len(p.ComboPresses)-1].Count

		max := maxPress
		if maxCombo > max {
			max = maxCombo
		}

		for _, press := range p.KeyPresses {
			perc := (float32(press.Count) / float32(max))
			luminance := uint8(255 * perc)
			color := heatmap.RgbaToRgb(color2.RGBA{R: 255, G: 0, B: 0, A: luminance}, heatmap.RGB{R: 255, G: 255, B: 255})

			/*
					as an alternative you can do this, which draws a border around the text, this is useful when you want to
					distinguish between normal and shifted presses
				        stroke: red;
					    stroke-width: 0.5em;
					    fill: white;
					    paint-order: stroke;
					    stroke-linejoin: round;
			*/

			fmt.Printf("svg > g:nth-of-type(%d) .keypos-%d rect { fill: #%02X%02X%02X; }\n", press.Layer+1, press.Position, color.R, color.G, color.B)
		}

		for _, press := range p.ComboPresses {
			perc := (float32(press.Count) / float32(max))
			luminance := uint8(255 * perc)
			color := heatmap.RgbaToRgb(color2.RGBA{R: 255, G: 0, B: 0, A: luminance}, heatmap.RGB{R: 255, G: 255, B: 255})
			fmt.Printf(".combopos-%d rect { stroke: #c9cccf; stroke-width: 1; fill: #%02X%02X%02X; }\n", press.Number-p.KeyMap.NumberOfKeys, color.R, color.G, color.B)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputParam, "input", "i", "heatmap.json", "e.g. ~/heatmap.json")
}
