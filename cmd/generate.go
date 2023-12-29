package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"image/color"
	"log"
	"sort"
	heatmappkg "zmk-heatmap/pkg/heatmap"
)

var inputParam string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate the heatmap",
	Long:  "Generate the heatmap",
	Run: func(cmd *cobra.Command, args []string) {
		// Remove the timestamp from the log messages
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

		heatmap, err := heatmappkg.Load(inputParam)
		if err != nil {
			log.Fatalln("Cannot read heatmap:", err)
		}

		sort.Slice(heatmap.KeyPresses, func(i, j int) bool {
			return heatmap.KeyPresses[i].GetTotalPressCounts() < heatmap.KeyPresses[j].GetTotalPressCounts()
		})
		maxPress := heatmap.KeyPresses[len(heatmap.KeyPresses)-1].GetTotalPressCounts()
		sort.Slice(heatmap.ComboPresses, func(i, j int) bool {
			return heatmap.ComboPresses[i].GetTotalPressCounts() < heatmap.ComboPresses[j].GetTotalPressCounts()
		})
		maxCombo := heatmap.ComboPresses[len(heatmap.ComboPresses)-1].GetTotalPressCounts()

		max := maxPress
		if maxCombo > max {
			max = maxCombo
		}

		for _, press := range heatmap.KeyPresses {
			perc := (float32(press.GetTotalPressCounts()) / float32(max))
			luminance := uint8(255 * perc)
			color := heatmappkg.RgbaToRgb(color.RGBA{R: 255, G: 0, B: 0, A: luminance}, heatmappkg.RGB{R: 255, G: 255, B: 255})

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

		for _, press := range heatmap.ComboPresses {
			perc := (float32(press.GetTotalPressCounts()) / float32(max))
			luminance := uint8(255 * perc)
			color := heatmappkg.RgbaToRgb(color.RGBA{R: 255, G: 0, B: 0, A: luminance}, heatmappkg.RGB{R: 255, G: 255, B: 255})
			fmt.Printf(".combopos-%d rect { stroke: #c9cccf; stroke-width: 1; fill: #%02X%02X%02X; }\n", press.Number, color.R, color.G, color.B)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputParam, "input", "i", "heatmap.json", "e.g. ~/heatmap.json")
}
