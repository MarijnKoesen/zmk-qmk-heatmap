package collector

import (
	"zmk-heatmap/pkg/heatmap"
)

type LogProcessor interface {
	Process(logLine string, heatmap *heatmap.Heatmap) error
}
