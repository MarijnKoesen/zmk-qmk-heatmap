package heatmap

import (
	"image/color"
)

type RGB struct {
	R uint8
	G uint8
	B uint8
}

func RgbaToRgb(color color.RGBA, background RGB) RGB {
	source := toFloat(color)
	bg := background.toFloat()
	luminance := float32(color.A) / 255

	target := floatColor{
		r: ((1 - luminance) * bg.r) + (luminance * source.r),
		g: ((1 - luminance) * bg.g) + (luminance * source.g),
		b: ((1 - luminance) * bg.b) + (luminance * source.b),
	}
	return target.toRGB()
}

type floatColor struct {
	r float32
	g float32
	b float32
}

func toFloat(rgba color.RGBA) floatColor {
	return floatColor{
		r: float32(rgba.R) / 255,
		g: float32(rgba.G) / 255,
		b: float32(rgba.B) / 255,
	}
}

func (rgb RGB) toFloat() floatColor {
	return floatColor{
		r: float32(rgb.R) / 255,
		g: float32(rgb.G) / 255,
		b: float32(rgb.B) / 255,
	}
}

func (c floatColor) toRGB() RGB {
	return RGB{
		R: uint8(255 * c.r),
		G: uint8(255 * c.g),
		B: uint8(255 * c.b),
	}
}
