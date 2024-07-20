package keymap

type Keymap struct {
	Layers []*Layer
	Combos []*Combo
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

func New() *Keymap {
	return &Keymap{
		Layers: make([]*Layer, 0),
		Combos: make([]*Combo, 0),
	}
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

func (k *Keymap) AddLayer(name string, rows []*Row) *Layer {
	layer := &Layer{
		Name: name,
		Rows: rows,
	}
	k.Layers = append(k.Layers, layer)

	return layer
}

func (r *Row) AddKey(key *Key) {
	r.Keys = append(r.Keys, key)
}
