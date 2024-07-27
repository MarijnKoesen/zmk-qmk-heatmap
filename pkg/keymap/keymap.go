package keymap

import "fmt"

type Keymap struct {
	NumberOfSensors int
	Layers          []*Layer
	Combos          []*Combo
}

type Layer struct {
	Name string
	Rows []*Row
}

type Row struct {
	Keys []*Key
}

type Key struct {
	Tap     string
	Hold    string
	Shifted string
}

type Combo struct {
	Keys   []int
	Layers []string
	Key    *Key
}

func New(numberOfSensors int) *Keymap {
	return &Keymap{
		NumberOfSensors: numberOfSensors,
		Layers:          make([]*Layer, 0),
		Combos:          make([]*Combo, 0),
	}
}

func (k *Keymap) Key(layer int, position int) *Key {
	if layer > len(k.Layers)-1 {
		return nil
	}

	l := k.Layers[layer]
	for r, row := range l.Rows {
		if position > len(row.Keys) {
			position -= len(row.Keys)
			continue
		}

		return k.Layers[layer].Rows[r].Keys[position]
	}

	return nil
}

func (k *Keymap) NumberOfKeys() int {
	numberOfKeys := 0
	if len(k.Layers) == 0 || len(k.Layers[0].Rows) == 0 {
		return numberOfKeys
	}

	for _, row := range k.Layers[0].Rows {
		numberOfKeys += len(row.Keys)
	}

	return numberOfKeys
}

func (k *Keymap) AddLayer(name string, rows []*Row) {
	k.Layers = append(k.Layers, &Layer{
		Name: name,
		Rows: rows,
	})
}

func (r *Row) Append(k *Key) {
	r.Keys = append(r.Keys, k)
}

// ComboByPosition returns the combo from the ZMK position
func (k *Keymap) ComboByPosition(pos int) (*Combo, error) {
	idx := pos - k.NumberOfKeys() - k.NumberOfSensors
	if idx < 0 || idx > len(k.Combos)-1 {
		return nil, fmt.Errorf("cannot find combo")
	}

	return k.Combos[idx], nil
}
