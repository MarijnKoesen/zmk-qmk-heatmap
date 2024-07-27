package heatmap

import (
	"encoding/json"
	"io"
	"os"
	"slices"
)

type Heatmap struct {
	KeyPresses   []KeyPress   `yaml:"keyPresses"`
	ComboPresses []ComboPress `yaml:"comboPresses"`
}

func (m *Heatmap) RegisterKeyPress(layer int, position int, pressType PressType) {
	for i := range m.KeyPresses {
		if m.KeyPresses[i].Position == position &&
			m.KeyPresses[i].Layer == layer {
			m.KeyPresses[i].RegisterPress(pressType)
			// fmt.Printf("key press: %s\n", m.KeyPresses[i])
			return
		}
	}

	press := KeyPress{
		Layer:    layer,
		Position: position,
	}
	press.RegisterPress(pressType)

	// fmt.Printf("key press: %s\n", press)
	m.KeyPresses = append(m.KeyPresses, press)
}

func (m *Heatmap) RegisterComboPress(layer int, keys []int, pressType PressType) {
	for i := range m.ComboPresses {
		if m.ComboPresses[i].Layer == layer && slices.Equal(m.ComboPresses[i].Keys, keys) {
			// fmt.Printf("key press: %s\n", m.KeyPresses[i])
			m.ComboPresses[i].RegisterPress(pressType)
			return
		}
	}

	press := ComboPress{
		Layer: layer,
		Keys:  keys,
	}
	press.RegisterPress(pressType)
	// fmt.Printf("key press: %s\n", press)

	m.ComboPresses = append(m.ComboPresses, press)
}

func (h *Heatmap) GetPressCount() (count int) {
	for i := range h.KeyPresses {
		count += h.KeyPresses[i].Taps + h.KeyPresses[i].Holds + h.KeyPresses[i].Shifts
	}

	for i := range h.ComboPresses {
		count += h.ComboPresses[i].Taps + h.ComboPresses[i].Holds + h.ComboPresses[i].Shifts
	}

	return
}

func (h *Heatmap) ToJson() (jsonData []byte, err error) {
	jsonData, err = json.Marshal(h)
	if err != nil {
		return
	}

	return
}

func FromJson(jsonData []byte) (heatmap *Heatmap, err error) {
	heatmap = &Heatmap{
		KeyPresses:   []KeyPress{},
		ComboPresses: []ComboPress{},
	}

	err = json.Unmarshal(jsonData, heatmap)
	return
}

func (h *Heatmap) Save(path string) (err error) {
	bytes, err := h.ToJson()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, bytes, 0o644)
	return
}

func Load(path string) (heatmap *Heatmap, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, _ := io.ReadAll(jsonFile)
	heatmap, err = FromJson(bytes)
	return
}

func New() (heatmap *Heatmap) {
	return &Heatmap{
		KeyPresses:   []KeyPress{},
		ComboPresses: []ComboPress{},
	}
}
