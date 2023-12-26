package keymap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var keymapTests = []struct {
	filename             string
	expectedNumberOfKeys int
	expectedCombos       []Combo
}{
	{
		"testdata/keymaps/3x5.yaml",
		36,
		[]Combo{
			{Keys: []int{22, 23}, Key: "`", Layers: []string{"DEF"}},
			{Keys: []int{21, 22}, Key: "~", Layers: []string{"DEF"}},
			{Keys: []int{6, 7}, Key: ";", Layers: []string{"DEF"}},
			{Keys: []int{12, 13}, Key: "(", Layers: []string{"DEF"}},
			{Keys: []int{16, 17}, Key: ")", Layers: []string{"DEF"}},
			{Keys: []int{11, 12}, Key: "[", Layers: []string{"DEF"}},
			{Keys: []int{17, 18}, Key: "]", Layers: []string{"DEF"}},
			{Keys: []int{26, 27}, Key: "\\", Layers: []string{"DEF"}},
			{Keys: []int{27, 28}, Key: "|", Layers: []string{"DEF"}},
		},
	},
	{
		// no combos
		"testdata/keymaps/4x12.MIT.yaml",
		47,
		[]Combo{},
	},
	{
		// combos with lots of keys like arc_scale, but also combo's without a specified layer
		"testdata/keymaps/ardux.yaml",
		8,
		[]Combo{
			{Keys: []int{3, 7}, Key: "ENT", Layers: []string{}},
			{Keys: []int{5, 2}, Key: "DEL", Layers: []string{}},
			{Keys: []int{2, 6}, Key: "SFT", Layers: []string{}},
			{Keys: []int{7, 6}, Key: "C", Layers: []string{"ARDUX"}},
			{Keys: []int{5, 4}, Key: "N", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 2}, Key: "F", Layers: []string{"ARDUX"}},
			{Keys: []int{2, 1}, Key: "G", Layers: []string{"ARDUX"}},
			{Keys: []int{6, 5}, Key: "U", Layers: []string{"ARDUX"}},
			{Keys: []int{1, 0}, Key: "J", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 6}, Key: ".", Layers: []string{"ARDUX"}},
			{Keys: []int{1, 5}, Key: "!", Layers: []string{"ARDUX"}},
			{Keys: []int{2, 0}, Key: "V", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 2, 1, 0}, Key: "Z", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 1, 0}, Key: "Q", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 0}, Key: "W", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 2, 1}, Key: "D", Layers: []string{"ARDUX"}},
			{Keys: []int{2, 1, 0}, Key: "X", Layers: []string{"ARDUX"}},
			{Keys: []int{6, 4}, Key: "K", Layers: []string{"ARDUX"}},
			{Keys: []int{7, 5}, Key: "H", Layers: []string{"ARDUX"}},
			{Keys: []int{7, 6, 5, 4}, Key: "SPC", Layers: []string{"ARDUX"}},
			{Keys: []int{7, 5, 4}, Key: "P", Layers: []string{"ARDUX"}},
			{Keys: []int{7, 6, 5}, Key: "L", Layers: []string{"ARDUX"}},
			{Keys: []int{7, 4}, Key: "B", Layers: []string{"ARDUX"}},
			{Keys: []int{6, 5, 4}, Key: "M", Layers: []string{"ARDUX"}},
			{Keys: []int{3, 2}, Key: "7", Layers: []string{"Number"}},
			{Keys: []int{2, 1}, Key: "8", Layers: []string{"Number"}},
			{Keys: []int{7, 6}, Key: "9", Layers: []string{"Number"}},
			{Keys: []int{6, 5}, Key: "0", Layers: []string{"Number"}},
			{Keys: []int{7, 2}, Key: "BSPC", Layers: []string{"Number", "Symbol", "Paren", "Nav", "BT", "Custom"}},
			{Keys: []int{7, 6, 5, 4}, Key: "SPC", Layers: []string{"Number", "Symbol", "Paren", "Nav", "BT", "Custom"}},
		},
	},
	{
		// no combos
		"testdata/keymaps/combo_test.yaml",
		9,
		[]Combo{
			{Keys: []int{0, 1}, Key: "1AB", Layers: []string{}},
			{Keys: []int{4, 8}, Key: "2B3C", Layers: []string{}},
			{Keys: []int{0, 1, 2}, Key: "1ABC", Layers: []string{}},
			{Keys: []int{0, 2}, Key: "1AC", Layers: []string{}},
			{Keys: []int{6, 8}, Key: "3AC", Layers: []string{}},
			{Keys: []int{0, 3, 6}, Key: "123A", Layers: []string{}},
			{Keys: []int{2, 8}, Key: "13C", Layers: []string{}},
			{Keys: []int{0, 3}, Key: "12A", Layers: []string{}},
		},
	},
	//{
	//	// combo with holds
	//	"testdata/keymaps/keymap_hold_in_combo.yaml",
	//	9,
	//	[]Combo{
	//		{Keys: []int{0, 1}, Key: "1AB", Layers: []string{}},
	//		{Keys: []int{4, 8}, Key: "2B3C", Layers: []string{}},
	//		{Keys: []int{0, 1, 2}, Key: "1ABC", Layers: []string{}},
	//		{Keys: []int{0, 2}, Key: "1AC", Layers: []string{}},
	//		{Keys: []int{6, 8}, Key: "3AC", Layers: []string{}},
	//		{Keys: []int{0, 3, 6}, Key: "123A", Layers: []string{}},
	//		{Keys: []int{2, 8}, Key: "13C", Layers: []string{}},
	//		{Keys: []int{0, 3}, Key: "12A", Layers: []string{}},
	//	},
	//},
}

func TestParser(t *testing.T) {
	for _, test := range keymapTests {
		keymap, err := Load("../../" + test.filename)
		assert.Nilf(t, err, "Expected no error")
		assert.Equal(t, test.expectedNumberOfKeys, keymap.NumberOfKeys)
		assert.Equal(t, test.expectedCombos, keymap.Combos)
	}
}
