package heatmap

import "strconv"

type PressType int

const (
	Tap PressType = iota
	Hold
	Shifted
)

type ComboPress struct {
	Layer  int
	Keys   []int
	Taps   int
	Holds  int
	Shifts int
}

type KeyPress struct {
	Layer    int
	Position int
	Taps     int
	Holds    int
	Shifts   int
}

func (p *ComboPress) RegisterPress(pressType PressType) {
	switch pressType {
	case Tap:
		p.Taps++
		return
	case Hold:
		p.Holds++
		return
	case Shifted:
		p.Shifts++
		return
	}
}

func (p *KeyPress) RegisterPress(pressType PressType) {
	switch pressType {
	case Tap:
		p.Taps++
		return
	case Hold:
		p.Holds++
		return
	case Shifted:
		p.Shifts++
		return
	}
}

func (p *KeyPress) GetTotalPressCounts() (pressCount int) {
	return p.Taps + p.Shifts + p.Holds
}

func (p *ComboPress) GetTotalPressCounts() (pressCount int) {
	return p.Taps + p.Shifts + p.Holds
}

func (t PressType) String() string {
	switch t {
	case Tap:
		return "Tap"
	case Hold:
		return "Hold"
	case Shifted:
		return "Shifted"
	default:
		return "Undefined"
	}
}

func (k KeyPress) String() string {
	return "KeyPress[Layer: " + strconv.Itoa(k.Layer) + ", Position: " + strconv.Itoa(k.Position) + ", Taps: " + strconv.Itoa(k.Taps) + ", Holds: " + strconv.Itoa(k.Holds) + ", Shifts: " + strconv.Itoa(k.Shifts) + "]"
}
