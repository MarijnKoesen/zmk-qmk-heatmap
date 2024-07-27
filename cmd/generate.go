package cmd

import (
	"fmt"
	"image/color"
	"log"
	"slices"
	"sort"

	"github.com/spf13/cobra"

	heatmappkg "zmk-heatmap/pkg/heatmap"
	"zmk-heatmap/pkg/keymap"
)

var inputParam string

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputParam, "input", "i", "heatmap.json", "e.g. ~/heatmap.json")
	generateCmd.Flags().StringVarP(&keymapParam, "keymap", "m", "keymap.yaml", "e.g. ~/keymap.yaml")
	generateCmd.Flags().IntVarP(&numSensors, "sensors", "s", 0, "the number of sensors (encoders) of the keyboard, this value is needed to properly collect combos")
}

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

		keymapFile := keymapParam
		keymapp, err := keymap.Load(keymapFile, numSensors)
		if err != nil {
			log.Fatalln("Cannot load the keymap:", err)
		}
		log.Println("Loading keymap", keymapFile)

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
			comboIndex, err := getComboIndex(keymapp, press)
			if err != nil {
				panic(err)
			}

			perc := (float32(press.GetTotalPressCounts()) / float32(max))
			luminance := uint8(255 * perc)
			color := heatmappkg.RgbaToRgb(color.RGBA{R: 255, G: 0, B: 0, A: luminance}, heatmappkg.RGB{R: 255, G: 255, B: 255})
			fmt.Printf("svg > g:nth-of-type(%d) .combopos-%d rect { stroke: #c9cccf; stroke-width: 1; fill: #%02X%02X%02X; }\n", press.Layer+1, comboIndex, color.R, color.G, color.B)
		}
	},
}

func getComboIndex(keymap *keymap.Keymap, press heatmappkg.ComboPress) (int, error) {
	layerName := keymap.Layers[press.Layer].Name

	index := 0
	for _, combo := range keymap.Combos {
		if !slices.Contains(combo.Layers, layerName) {
			// We must not increase the index in case the combo is not active on this layer, as it will be filtered
			// out by the keymap tool so the svg numbers would be mismatched if we'd increase them here
			continue
		}

		if slices.Equal(combo.Keys, press.Keys) {
			return index, nil
		}

		index++
	}

	return 0, fmt.Errorf("cannot find combo")
}
