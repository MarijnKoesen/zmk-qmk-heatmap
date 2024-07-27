package heatmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPresses(t *testing.T) {
	heatmap := Heatmap{}

	heatmap.RegisterKeyPress(1, 1, Tap)
	heatmap.RegisterKeyPress(1, 1, Tap)
	heatmap.RegisterKeyPress(2, 1, Tap)
	heatmap.RegisterKeyPress(1, 2, Tap)
	heatmap.RegisterKeyPress(1, 1, Tap)
	heatmap.RegisterKeyPress(1, 2, Hold)
	heatmap.RegisterKeyPress(1, 2, Hold)

	heatmap.RegisterKeyPress(1, 3, Hold)
	heatmap.RegisterKeyPress(1, 3, Hold)
	heatmap.RegisterKeyPress(1, 3, Tap)
	heatmap.RegisterKeyPress(1, 3, Shifted)
	heatmap.RegisterKeyPress(1, 3, Shifted)
	heatmap.RegisterKeyPress(1, 3, Shifted)

	assert.Equal(t,
		[]KeyPress{
			{Layer: 1, Position: 1, Taps: 3, Holds: 0, Shifts: 0},
			{Layer: 2, Position: 1, Taps: 1, Holds: 0, Shifts: 0},
			{Layer: 1, Position: 2, Taps: 1, Holds: 2, Shifts: 0},
			{Layer: 1, Position: 3, Taps: 1, Holds: 2, Shifts: 3},
		},
		heatmap.KeyPresses,
	)
	assert.Empty(t, heatmap.ComboPresses)
	assert.Equal(t, 13, heatmap.GetPressCount())
}

func TestComboPresses(t *testing.T) {
	heatmap := Heatmap{}

	heatmap.RegisterComboPress(1, []int{1, 2}, Tap)
	heatmap.RegisterComboPress(1, []int{1, 2}, Tap)
	heatmap.RegisterComboPress(2, []int{1, 2}, Tap)
	heatmap.RegisterComboPress(1, []int{1, 2}, Tap)
	heatmap.RegisterComboPress(2, []int{1, 2}, Hold)
	heatmap.RegisterComboPress(2, []int{1, 2}, Hold)

	heatmap.RegisterComboPress(3, []int{3, 4}, Hold)
	heatmap.RegisterComboPress(3, []int{1, 2}, Hold)
	heatmap.RegisterComboPress(3, []int{5, 6}, Tap)
	heatmap.RegisterComboPress(3, []int{5, 6}, Shifted)
	heatmap.RegisterComboPress(3, []int{7, 8}, Shifted)
	heatmap.RegisterComboPress(3, []int{1, 2}, Shifted)
	heatmap.RegisterComboPress(3, []int{1, 2}, Shifted)

	assert.Equal(t,
		[]ComboPress{
			{Layer: 1, Keys: []int{1, 2}, Taps: 3, Holds: 0, Shifts: 0},
			{Layer: 2, Keys: []int{1, 2}, Taps: 1, Holds: 2, Shifts: 0},
			{Layer: 3, Keys: []int{3, 4}, Taps: 0, Holds: 1, Shifts: 0},
			{Layer: 3, Keys: []int{1, 2}, Taps: 0, Holds: 1, Shifts: 2},
			{Layer: 3, Keys: []int{5, 6}, Taps: 1, Holds: 0, Shifts: 1},
			{Layer: 3, Keys: []int{7, 8}, Taps: 0, Holds: 0, Shifts: 1},
		},
		heatmap.ComboPresses,
	)
	assert.Empty(t, heatmap.KeyPresses)
	assert.Equal(t, 13, heatmap.GetPressCount())
}

func TestJsonSerialization(t *testing.T) {
	heatmap := &Heatmap{}
	heatmap.RegisterKeyPress(1, 1, Tap)
	heatmap.RegisterKeyPress(1, 1, Tap)
	heatmap.RegisterKeyPress(2, 2, Tap)

	heatmap.RegisterComboPress(1, []int{1, 2}, Tap)
	heatmap.RegisterComboPress(1, []int{3, 4}, Tap)
	heatmap.RegisterComboPress(1, []int{1, 2}, Tap)
	heatmap.RegisterComboPress(2, []int{1, 2}, Hold)
	heatmap.RegisterComboPress(2, []int{2, 6}, Hold)

	jsonBytes, err := heatmap.ToJson()
	assert.Nil(t, err)

	expectedJson := []byte("{\"KeyPresses\":[{\"Layer\":1,\"Position\":1,\"Taps\":2,\"Holds\":0,\"Shifts\":0},{\"Layer\":2,\"Position\":2,\"Taps\":1,\"Holds\":0,\"Shifts\":0}],\"ComboPresses\":[{\"Layer\":1,\"Keys\":[1,2],\"Taps\":2,\"Holds\":0,\"Shifts\":0},{\"Layer\":1,\"Keys\":[3,4],\"Taps\":1,\"Holds\":0,\"Shifts\":0},{\"Layer\":2,\"Keys\":[1,2],\"Taps\":0,\"Holds\":1,\"Shifts\":0},{\"Layer\":2,\"Keys\":[2,6],\"Taps\":0,\"Holds\":1,\"Shifts\":0}]}")
	assert.Equal(t, expectedJson, jsonBytes)

	heatMapFromJson, err := FromJson(jsonBytes)
	assert.Nil(t, err)
	assert.Equal(t, heatmap, heatMapFromJson)
}
