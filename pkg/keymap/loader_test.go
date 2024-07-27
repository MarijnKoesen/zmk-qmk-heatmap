package keymap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	keymapTests := []struct {
		filename             string
		numberOfSensors      int
		expectedNumberOfKeys int
		expectedCombos       []*Combo
	}{
		{
			"testdata/keymaps/3x5.yaml",
			0,
			36,
			[]*Combo{
				{Keys: []int{22, 23}, Key: &Key{Tap: "`"}, Layers: []string{"DEF"}},
				{Keys: []int{21, 22}, Key: &Key{Tap: "~"}, Layers: []string{"DEF"}},
				{Keys: []int{6, 7}, Key: &Key{Tap: ";"}, Layers: []string{"DEF"}},
				{Keys: []int{12, 13}, Key: &Key{Tap: "("}, Layers: []string{"DEF"}},
				{Keys: []int{16, 17}, Key: &Key{Tap: ")"}, Layers: []string{"DEF"}},
				{Keys: []int{11, 12}, Key: &Key{Tap: "["}, Layers: []string{"DEF"}},
				{Keys: []int{17, 18}, Key: &Key{Tap: "]"}, Layers: []string{"DEF"}},
				{Keys: []int{26, 27}, Key: &Key{Tap: "\\"}, Layers: []string{"DEF"}},
				{Keys: []int{27, 28}, Key: &Key{Tap: "|"}, Layers: []string{"DEF"}},
			},
		},
		{
			// no combos
			"testdata/keymaps/4x12.MIT.yaml",
			0,
			47,
			[]*Combo{},
		},
		{
			// combos with lots of keys like arc_scale, but also combo's without a specified layer
			"testdata/keymaps/ardux.yaml",
			0,
			8,
			[]*Combo{
				{Keys: []int{3, 7}, Key: &Key{Tap: "ENT"}, Layers: []string{}},
				{Keys: []int{5, 2}, Key: &Key{Tap: "DEL"}, Layers: []string{}},
				{Keys: []int{2, 6}, Key: &Key{Tap: "SFT"}, Layers: []string{}},
				{Keys: []int{7, 6}, Key: &Key{Tap: "C"}, Layers: []string{"ARDUX"}},
				{Keys: []int{5, 4}, Key: &Key{Tap: "N"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 2}, Key: &Key{Tap: "F"}, Layers: []string{"ARDUX"}},
				{Keys: []int{2, 1}, Key: &Key{Tap: "G"}, Layers: []string{"ARDUX"}},
				{Keys: []int{6, 5}, Key: &Key{Tap: "U"}, Layers: []string{"ARDUX"}},
				{Keys: []int{1, 0}, Key: &Key{Tap: "J"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 6}, Key: &Key{Tap: "."}, Layers: []string{"ARDUX"}},
				{Keys: []int{1, 5}, Key: &Key{Tap: "!"}, Layers: []string{"ARDUX"}},
				{Keys: []int{2, 0}, Key: &Key{Tap: "V"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 2, 1, 0}, Key: &Key{Tap: "Z"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 1, 0}, Key: &Key{Tap: "Q"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 0}, Key: &Key{Tap: "W"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 2, 1}, Key: &Key{Tap: "D"}, Layers: []string{"ARDUX"}},
				{Keys: []int{2, 1, 0}, Key: &Key{Tap: "X"}, Layers: []string{"ARDUX"}},
				{Keys: []int{6, 4}, Key: &Key{Tap: "K"}, Layers: []string{"ARDUX"}},
				{Keys: []int{7, 5}, Key: &Key{Tap: "H"}, Layers: []string{"ARDUX"}},
				{Keys: []int{7, 6, 5, 4}, Key: &Key{Tap: "SPC"}, Layers: []string{"ARDUX"}},
				{Keys: []int{7, 5, 4}, Key: &Key{Tap: "P"}, Layers: []string{"ARDUX"}},
				{Keys: []int{7, 6, 5}, Key: &Key{Tap: "L"}, Layers: []string{"ARDUX"}},
				{Keys: []int{7, 4}, Key: &Key{Tap: "B"}, Layers: []string{"ARDUX"}},
				{Keys: []int{6, 5, 4}, Key: &Key{Tap: "M"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 2}, Key: &Key{Tap: "7"}, Layers: []string{"Number"}},
				{Keys: []int{2, 1}, Key: &Key{Tap: "8"}, Layers: []string{"Number"}},
				{Keys: []int{7, 6}, Key: &Key{Tap: "9"}, Layers: []string{"Number"}},
				{Keys: []int{6, 5}, Key: &Key{Tap: "0"}, Layers: []string{"Number"}},
				{Keys: []int{7, 2}, Key: &Key{Tap: "BSPC"}, Layers: []string{"Number", "Symbol", "Paren", "Nav", "BT", "Custom"}},
				{Keys: []int{7, 6, 5, 4}, Key: &Key{Tap: "SPC"}, Layers: []string{"Number", "Symbol", "Paren", "Nav", "BT", "Custom"}},
				{Keys: []int{3, 5}, Key: &Key{Tap: ","}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 4}, Key: &Key{Tap: "/"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 2, 4}, Key: &Key{Tap: "ESC"}, Layers: []string{}},
				{Keys: []int{2, 7, 5}, Key: &Key{Tap: "Nav"}, Layers: []string{}},
				{Keys: []int{3, 6, 5}, Key: &Key{Tap: "'"}, Layers: []string{"ARDUX"}},
				{Keys: []int{3, 2, 1, 4}, Key: &Key{Tap: "TAB"}, Layers: []string{}},
				{Keys: []int{7, 2, 1, 0}, Key: &Key{Tap: "LSHFT", Hold: "sticky"}, Layers: []string{}},
				{Keys: []int{3, 7, 0, 4}, Key: &Key{Tap: "BT"}, Layers: []string{}},
				{Keys: []int{2, 6, 1, 5}, Key: &Key{Tap: "BT CLR"}, Layers: []string{}},
				{Keys: []int{4, 5, 6, 3}, Key: &Key{Tap: "CAPS"}, Layers: []string{"ARDUX"}},
				{Keys: []int{7, 0}, Key: &Key{Tap: "LCTRL", Hold: "sticky"}, Layers: []string{}},
				{Keys: []int{6, 0}, Key: &Key{Tap: "LGUI", Hold: "sticky"}, Layers: []string{}},
				{Keys: []int{5, 0}, Key: &Key{Tap: "LALT", Hold: "sticky"}, Layers: []string{}},
			},
		},
		{
			// no combos
			"testdata/keymaps/combo_test.yaml",
			0,
			9,
			[]*Combo{
				{Keys: []int{0, 1}, Key: &Key{Tap: "1AB"}, Layers: []string{}},
				{Keys: []int{4, 8}, Key: &Key{Tap: "2B3C"}, Layers: []string{}},
				{Keys: []int{0, 1, 2}, Key: &Key{Tap: "1ABC"}, Layers: []string{}},
				{Keys: []int{0, 2}, Key: &Key{Tap: "1AC"}, Layers: []string{}},
				{Keys: []int{6, 8}, Key: &Key{Tap: "3AC"}, Layers: []string{}},
				{Keys: []int{0, 3, 6}, Key: &Key{Tap: "123A"}, Layers: []string{}},
				{Keys: []int{2, 8}, Key: &Key{Tap: "13C"}, Layers: []string{}},
				{Keys: []int{0, 3}, Key: &Key{Tap: "12A"}, Layers: []string{}},
			},
		},
		{
			// combo with holds
			"testdata/keymaps/corne.yaml",
			2,
			42,
			[]*Combo{
				{Keys: []int{2, 3}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "ESC"}},
				{Keys: []int{1, 13, 25}, Layers: []string{"nav"}, Key: &Key{Tap: "&bootloader"}},
				{Keys: []int{3, 4}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "RETURN"}},
				{Keys: []int{14, 15}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "TAB", Hold: "Alt+ LCTRL", Shifted: ""}},
				{Keys: []int{15, 16}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "&key_repeat", Hold: "Sft+ LCTRL", Shifted: ""}},
				{Keys: []int{26, 28}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "Ctl+ X", Hold: "", Shifted: ""}},
				{Keys: []int{26, 27}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "Ctl+ INS", Hold: "", Shifted: ""}},
				{Keys: []int{27, 28}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "Sft+ INS", Hold: "", Shifted: ""}},
				{Keys: []int{7, 8}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "BSPC", Hold: "", Shifted: ""}},
				{Keys: []int{8, 9}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "DEL", Hold: "", Shifted: ""}},
				{Keys: []int{19, 20}, Layers: []string{"default", "symbol"}, Key: &Key{Tap: "(", Hold: "Sft+ LCTRL", Shifted: "<"}},
				{Keys: []int{20, 21}, Layers: []string{"default", "symbol"}, Key: &Key{Tap: ")", Hold: "Alt+ LCTRL", Shifted: ">"}},
				{Keys: []int{19, 20}, Layers: []string{"nav"}, Key: &Key{Tap: "<", Hold: "", Shifted: ""}},
				{Keys: []int{20, 21}, Layers: []string{"nav"}, Key: &Key{Tap: ">", Hold: "", Shifted: ""}},
				{Keys: []int{31, 32}, Layers: []string{"default", "symbol"}, Key: &Key{Tap: "[", Hold: "", Shifted: ""}},
				{Keys: []int{32, 33}, Layers: []string{"default", "symbol"}, Key: &Key{Tap: "]", Hold: "", Shifted: ""}},
				{Keys: []int{31, 32}, Layers: []string{"nav"}, Key: &Key{Tap: "{", Hold: "", Shifted: ""}},
				{Keys: []int{32, 33}, Layers: []string{"nav"}, Key: &Key{Tap: "}", Hold: "", Shifted: ""}},
				{Keys: []int{2, 14}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "@", Hold: "", Shifted: ""}},
				{Keys: []int{3, 15}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "#", Hold: "", Shifted: ""}},
				{Keys: []int{4, 16}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "$", Hold: "", Shifted: ""}},
				{Keys: []int{5, 17}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "%", Hold: "", Shifted: ""}},
				{Keys: []int{14, 26}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "`", Hold: "", Shifted: ""}},
				{Keys: []int{15, 27}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "\\", Hold: "", Shifted: ""}},
				{Keys: []int{16, 28}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "=", Hold: "", Shifted: ""}},
				{Keys: []int{17, 29}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "~", Hold: "", Shifted: ""}},
				{Keys: []int{6, 18}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "^", Hold: "", Shifted: ""}},
				{Keys: []int{7, 19}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "+", Hold: "", Shifted: ""}},
				{Keys: []int{8, 20}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "*", Hold: "", Shifted: ""}},
				{Keys: []int{9, 21}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "&", Hold: "", Shifted: ""}},
				{Keys: []int{18, 30}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "_", Hold: "", Shifted: ""}},
				{Keys: []int{19, 31}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "-", Hold: "", Shifted: ""}},
				{Keys: []int{20, 32}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "/", Hold: "", Shifted: ""}},
				{Keys: []int{21, 33}, Layers: []string{"default", "nav", "symbol"}, Key: &Key{Tap: "|", Hold: "", Shifted: ""}},
			},
		},
	}

	for _, test := range keymapTests {
		keymap, err := Load("../../"+test.filename, test.numberOfSensors)
		require.NoError(t, err)

		t.Run("with "+test.filename, func(t *testing.T) {
			require.Equal(t, test.expectedNumberOfKeys, keymap.NumberOfKeys())
			require.Equal(t, test.expectedCombos, keymap.Combos)

			for _, layer := range keymap.Layers {
				numKeysInLayer := 0
				for _, row := range layer.Rows {
					numKeysInLayer += len(row.Keys)
				}

				require.Equal(t, keymap.NumberOfKeys(), numKeysInLayer)
			}
		})

		t.Run("ComboByPosition takes the sensors into account", func(t *testing.T) {
			if len(test.expectedCombos) == 0 {
				combo, err := keymap.ComboByPosition(0)
				require.Nil(t, combo)
				require.Error(t, err)
			} else {
				combo, err := keymap.ComboByPosition(test.numberOfSensors + test.expectedNumberOfKeys)
				require.NoError(t, err)
				require.Same(t, keymap.Combos[0], combo)
			}
		})
	}
}

func TestLoadingLoadsTheLayersInTheCorrectOrder(t *testing.T) {
	keymap, err := Load("../../testdata/keymaps/corne.yaml", 0)
	require.NoError(t, err)

	require.Equal(t, "default", keymap.Layers[0].Name)
	require.Equal(t, "symbol", keymap.Layers[1].Name)
	require.Equal(t, "nav", keymap.Layers[2].Name)
	require.Equal(t, "left", keymap.Layers[3].Name)
	require.Equal(t, "tri", keymap.Layers[4].Name)

	require.Equal(t, "Q", keymap.Key(0, 1).Tap)
	require.Equal(t, "W", keymap.Key(0, 2).Tap)
	require.Equal(t, "L", keymap.Key(0, 7).Tap)
	require.Equal(t, "A", keymap.Key(0, 13).Tap)

	require.Equal(t, "R", keymap.Key(0, 14).Tap)
	require.Equal(t, "LALT", keymap.Key(0, 14).Hold)

	require.Equal(t, ":", keymap.Key(1, 14).Tap)
}
