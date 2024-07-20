package collector

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	. "zmk-heatmap/pkg/heatmap"
	"zmk-heatmap/pkg/keymap"
)

const zmkTestKeyMapNumberOfKeys int = 34

var countTests = []struct {
	filename       string
	expectedKeys   []KeyPress
	expectedCombos []ComboPress
}{
	{
		// primary Layer, simple key press on left side
		"testdata/zmk/logs/ferris-sweep/q.log",
		[]KeyPress{
			{Layer: 0, Position: 0, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// primary Layer, simple key press on left side
		"testdata/zmk/logs/ferris-sweep/w.log",
		[]KeyPress{
			{Layer: 0, Position: 1, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// primary Layer, simple key press on right side
		"testdata/zmk/logs/ferris-sweep/l.log",
		[]KeyPress{
			{Layer: 0, Position: 6, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on left side
		"testdata/zmk/logs/ferris-sweep/semicolon.log",
		[]KeyPress{
			{Layer: 1, Position: 13, Taps: 1},
			{Layer: 0, Position: 30, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on left side
		"testdata/zmk/logs/ferris-sweep/parentheses-close.log",
		[]KeyPress{
			{Layer: 1, Position: 11, Taps: 1},
			{Layer: 0, Position: 30, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on right side
		"testdata/zmk/logs/ferris-sweep/1.log",
		[]KeyPress{
			{Layer: 1, Position: 26, Taps: 1},
			{Layer: 0, Position: 30, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mod left side
		"testdata/zmk/logs/ferris-sweep/ctrl-c.log",
		[]KeyPress{
			{Layer: 0, Position: 22, Taps: 1},
			{Layer: 0, Position: 17, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mod left side
		"testdata/zmk/logs/ferris-sweep/command-t.log",
		[]KeyPress{
			{Layer: 0, Position: 13, Taps: 1},
			{Layer: 0, Position: 10, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mod left + right side
		"testdata/zmk/logs/ferris-sweep/command-a.log",
		[]KeyPress{
			{Layer: 0, Position: 10, Taps: 1},
			{Layer: 0, Position: 19, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// multiple presses
		"testdata/zmk/logs/ferris-sweep/multiple-presses.log",
		[]KeyPress{
			{Layer: 0, Position: 12, Taps: 3},
			{Layer: 0, Position: 13, Taps: 1},
			{Layer: 0, Position: 6, Taps: 2},
			{Layer: 1, Position: 7, Taps: 3},
			{Layer: 0, Position: 30, Taps: 1},
		},
		[]ComboPress{},
	},
	{
		// combo-test, once L, once N, then combo of L+N=+
		"testdata/zmk/logs/ferris-sweep/combo-test-on-right-side.log",
		[]KeyPress{
			{Layer: 0, Position: 6, Taps: 1},
			{Layer: 0, Position: 16, Taps: 1},
		},
		[]ComboPress{
			{Number: 61 - zmkTestKeyMapNumberOfKeys, Taps: 1},
		},
	},
	{
		// combo-test, once B, once G, then combo of B+G=%
		"testdata/zmk/logs/ferris-sweep/combo-test-on-left-side.log",
		[]KeyPress{
			{Layer: 0, Position: 4, Taps: 1},
			{Layer: 0, Position: 14, Taps: 1},
		},
		[]ComboPress{
			{Number: 55 - zmkTestKeyMapNumberOfKeys, Taps: 1},
		},
	},
	{
		"testdata/zmk/logs/ferris-sweep/AAAaaAaaaa.log",
		[]KeyPress{
			{Layer: 0, Position: 10, Taps: 6, Shifts: 4}, // A
			{Layer: 0, Position: 33, Taps: 0, Shifts: 2}, // shift
		},
		[]ComboPress{},
	},
	{
		// combo first (, then <
		"testdata/zmk/logs/ferris-sweep/combo_without_hold_then_with.log",
		[]KeyPress{
			{Layer: 0, Position: 33, Shifts: 1}, // shift
		},
		[]ComboPress{
			{Number: 44 - zmkTestKeyMapNumberOfKeys, Taps: 1, Shifts: 1},
		},
	},
}

func parseTestFile(filename string) (heatmap *Heatmap) {
	file, err := os.Open("../../" + filename)

	if err != nil {
		log.Fatal(err)
	}

	keyMap, err := keymap.Load("../../testdata/zmk/logs/ferris-sweep/keymap.yaml")
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

func TestParser(t *testing.T) {
	for _, test := range countTests {
		heatmap := parseTestFile(test.filename)

		assert.Equal(t, test.expectedKeys, heatmap.KeyPresses)
		assert.Equal(t, test.expectedCombos, heatmap.ComboPresses)
	}
}
