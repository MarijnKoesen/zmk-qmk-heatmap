package main

import (
	"fmt"
	"log"
	"sort"
	log_parser "zmk-heatmap/pkg/log-parser"
)

const save_file = "heatmap_data.json"

type Color struct {
	r uint8
	g uint8
	b uint8
}

type FloatColor struct {
	r float32
	g float32
	b float32
}

func (c Color) toFloat() FloatColor {
	return FloatColor{
		r: float32(c.r) / 255,
		g: float32(c.g) / 255,
		b: float32(c.b) / 255,
	}
}

func (c FloatColor) toColor() Color {
	return Color{
		r: uint8(255 * c.r),
		g: uint8(255 * c.g),
		b: uint8(255 * c.b),
	}
}

func rgba_to_rgb(color Color, a uint8, background Color) Color {
	source := color.toFloat()
	bg := background.toFloat()
	luminance := float32(a) / 255

	target := FloatColor{
		r: ((1 - luminance) * bg.r) + (luminance * source.r),
		g: ((1 - luminance) * bg.g) + (luminance * source.g),
		b: ((1 - luminance) * bg.b) + (luminance * source.b),
	}
	return target.toColor()

	//Source => Target = (BGColor + Source) =
	//Target.R = ((1 - Source.A) * BGColor.R) + (Source.A * Source.R)
	//Target.G = ((1 - Source.A) * BGColor.G) + (Source.A * Source.G)
	//Target.B = ((1 - Source.A) * BGColor.B) + (Source.A * Source.B)
}

func main() {
	fmt.Println("ok")

	p, err := log_parser.LoadParser(save_file)
	if err != nil {
		log.Fatal("Cannot read heatmap file:", save_file)
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
		luminande := uint8(255 * perc)
		//fmt.Println(press, " = ", luminande)
		//fmt.Printf("svg > g:nth-of-type(%d) .keypos-%d rect { fill: #FF0000%02X; }\n", press.Layer+1, press.Position, luminande)
		color := rgba_to_rgb(Color{r: 255, g: 0, b: 0}, luminande, Color{r: 255, g: 255, b: 255})
		fmt.Printf("svg > g:nth-of-type(%d) .keypos-%d rect { fill: #%02X%02X%02X; }\n", press.Layer+1, press.Position, color.r, color.g, color.b)
	}

	for _, press := range p.ComboPresses {
		perc := (float32(press.Count) / float32(max))
		luminande := uint8(255 * perc)
		//fmt.Println(press, " = ", luminande)
		color := rgba_to_rgb(Color{r: 255, g: 0, b: 0}, luminande, Color{r: 255, g: 255, b: 255})
		fmt.Printf(".combopos-%d rect { fill: #%02X%02X%02X; }\n", press.Number-p.KeyMap.NumberOfKeys, color.r, color.g, color.b)
	}
}
