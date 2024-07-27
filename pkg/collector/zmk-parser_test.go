package collector

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "zmk-heatmap/pkg/heatmap"
	"zmk-heatmap/pkg/keymap"
)

func TestParser(t *testing.T) {
	countTests := []struct {
		name            string
		logfile         string
		keymap          string
		numberOfSensors int
		expectedKeys    []KeyPress
		expectedCombos  []ComboPress
	}{
		{
			"primary Layer, simple key press on left side",
			"testdata/zmk/logs/ferris-sweep/q.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 0, Taps: 1},
			},
			[]ComboPress{},
		},
		{
			"primary Layer, simple key press on left side",
			"testdata/zmk/logs/ferris-sweep/w.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 1, Taps: 1},
			},
			[]ComboPress{},
		},
		{
			"primary Layer, simple key press on right side",
			"testdata/zmk/logs/ferris-sweep/l.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 6, Taps: 1},
			},
			[]ComboPress{},
		},
		{
			"secondary Layer on left side",
			"testdata/zmk/logs/ferris-sweep/semicolon.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 1, Position: 13, Taps: 1},  // semicolon
				{Layer: 0, Position: 30, Holds: 1}, // layer
			},
			[]ComboPress{},
		},
		{
			"secondary Layer on left side",
			"testdata/zmk/logs/ferris-sweep/parentheses-close.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 1, Position: 11, Taps: 1},  // parentheses
				{Layer: 0, Position: 30, Holds: 1}, // layer
			},
			[]ComboPress{},
		},
		{
			"secondary Layer on right side",
			"testdata/zmk/logs/ferris-sweep/1.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 1, Position: 26, Taps: 1},  // 1
				{Layer: 0, Position: 30, Holds: 1}, // layer
			},
			[]ComboPress{},
		},
		{
			"home row mod left side",
			"testdata/zmk/logs/ferris-sweep/ctrl-c.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 22, Taps: 1},  // c
				{Layer: 0, Position: 17, Holds: 1}, // ctrl
			},
			[]ComboPress{},
		},
		{
			"home row mod left side",
			"testdata/zmk/logs/ferris-sweep/command-t.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 13, Taps: 1},  // t
				{Layer: 0, Position: 10, Holds: 1}, // cmd
			},
			[]ComboPress{},
		},
		{
			"home row mod left + right side",
			"testdata/zmk/logs/ferris-sweep/command-a.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 10, Taps: 1},  // a
				{Layer: 0, Position: 19, Holds: 1}, // cmd
			},
			[]ComboPress{},
		},
		{
			"multiple presses",
			"testdata/zmk/logs/ferris-sweep/multiple-presses.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 12, Taps: 3},
				{Layer: 0, Position: 13, Taps: 1},
				{Layer: 0, Position: 6, Taps: 2},
				{Layer: 1, Position: 7, Taps: 3},
				{Layer: 0, Position: 30, Holds: 1}, // layer
			},
			[]ComboPress{},
		},
		{
			"combo-test, once L, once N, then combo of L+N=+",
			"testdata/zmk/logs/ferris-sweep/combo-test-on-right-side.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 6, Taps: 1},
				{Layer: 0, Position: 16, Taps: 1},
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{6, 16}, Taps: 1},
			},
		},
		{
			"combo-test, once B, once G, then combo of B+G=%",
			"testdata/zmk/logs/ferris-sweep/combo-test-on-left-side.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 4, Taps: 1},
				{Layer: 0, Position: 14, Taps: 1},
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{4, 14}, Taps: 1},
			},
		},
		{
			"AAAaaAaaaa",
			"testdata/zmk/logs/ferris-sweep/AAAaaAaaaa.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 10, Taps: 6, Shifts: 4}, // A,
				{Layer: 0, Position: 33, Taps: 2, Holds: 0},  // shift (we only have a 'tap' binding on this key, hence 2 taps and not 2 holds)
			},
			[]ComboPress{},
		},
		{
			"combo first (, then <",
			"testdata/zmk/logs/ferris-sweep/combo_right_side_without_hold_then_with.log",
			"testdata/zmk/logs/ferris-sweep/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 33, Taps: 1}, // shift (we only have a 'tap' binding on this key, hence taps and not a hold)
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{16, 17}, Taps: 1, Shifts: 1},
			},
		},
		{
			"combo first (, then <",
			"testdata/zmk/logs/aurora-corne-44key/combo_right_side_without_shift_then_with.log",
			"testdata/zmk/logs/aurora-corne-44key/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 40, Taps: 1}, // shift (we only have a 'tap' binding on this key, hence taps and not a hold)
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{19, 20}, Taps: 1, Shifts: 1},
			},
		},
		{
			"combo first =, then +",
			"testdata/zmk/logs/aurora-corne-44key/combo_left_side_without_shift_then_with.log",
			"testdata/zmk/logs/aurora-corne-44key/keymap.yaml",
			0,
			[]KeyPress{
				{Layer: 0, Position: 40, Taps: 1}, // shift (we only have a 'tap' binding on this key, hence taps and not a hold)
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{6, 18}, Taps: 1, Shifts: 1},
			},
		},
		{
			"combo [t+p => $] on left side",
			"testdata/zmk/logs/aurora-corne-44key/combo-left-side.log",
			"testdata/zmk/logs/aurora-corne-44key/keymap.yaml",
			2,
			[]KeyPress{},
			[]ComboPress{
				{Layer: 0, Keys: []int{4, 16}, Taps: 1},
			},
		},
		{
			"hold combo shift+control, then r",
			"testdata/zmk/logs/aurora-corne-44key/hold-combo-same-side.log",
			"testdata/zmk/logs/aurora-corne-44key/keymap.yaml",
			2,
			[]KeyPress{
				{Layer: 0, Position: 14, Shifts: 1}, // r
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{15, 16}, Holds: 1}, // left shift+control
			},
		},
		{
			"hold combo shift+control, then r",
			"testdata/zmk/logs/aurora-corne-44key/combo-right-hold-shift-ctrl-i.log",
			"testdata/zmk/logs/aurora-corne-44key/keymap.yaml",
			2,
			[]KeyPress{
				{Layer: 0, Position: 21, Shifts: 1}, // i
			},
			[]ComboPress{
				{Layer: 0, Keys: []int{19, 20}, Holds: 1}, // right shift+control
			},
		},
		{
			"press combo N+E => (",
			"testdata/zmk/logs/aurora-corne-44key/combo-right-parenthesis-open.log",
			"testdata/zmk/logs/aurora-corne-44key/keymap.yaml",
			2,
			[]KeyPress{},
			[]ComboPress{
				{Layer: 0, Keys: []int{19, 20}, Taps: 1}, // n+e
			},
		},
	}

	for _, test := range countTests {
		t.Run(test.name, func(t *testing.T) {
			heatmap := parseTestFile(test.logfile, test.keymap, test.numberOfSensors)

			assert.Equal(t, test.expectedKeys, heatmap.KeyPresses)
			assert.Equal(t, test.expectedCombos, heatmap.ComboPresses)
		})
	}
}

func parseTestFile(logFile string, keyMapFile string, numberOfSensors int) (heatmap *Heatmap) {
	file, err := os.Open("../../" + logFile)
	if err != nil {
		log.Fatal(err)
	}

	keyMap, err := keymap.Load("../../"+keyMapFile, numberOfSensors)
	if err != nil {
		log.Fatal(err)
	}

	parser := NewZmkLogParser(keyMap)
	heatmap = &Heatmap{
		KeyPresses:   []KeyPress{},
		ComboPresses: []ComboPress{},
	}

	logScanner := bufio.NewScanner(file)
	for logScanner.Scan() {
		parser.Parse(logScanner.Text(), heatmap)
	}

	if err := logScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return heatmap
}
