package collector

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"zmk-heatmap/pkg/keymap"
)

type Comparer interface {
	Compare(b ComboPress) int
}

var countTests = []struct {
	filename       string
	expectedKeys   []KeyPress
	expectedCombos []ComboPress
}{
	{
		// primary Layer, simple key press on left side
		"testdata/zmk/q.log",
		[]KeyPress{
			{Layer: 0, Position: 0, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// primary Layer, simple key press on left side
		"testdata/zmk/w.log",
		[]KeyPress{
			{Layer: 0, Position: 1, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// primary Layer, simple key press on right side
		"testdata/zmk/l.log",
		[]KeyPress{
			{Layer: 0, Position: 6, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on left side
		"testdata/zmk/semicolon.log",
		[]KeyPress{
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 13, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on left side
		"testdata/zmk/parentheses-close.log",
		[]KeyPress{
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 11, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on right side
		"testdata/zmk/1.log",
		[]KeyPress{
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 26, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mode left side
		"testdata/zmk/ctrl-c.log",
		[]KeyPress{
			{Layer: 0, Position: 17, Count: 1},
			{Layer: 0, Position: 22, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mode left side
		"testdata/zmk/command-t.log",
		[]KeyPress{
			{Layer: 0, Position: 10, Count: 1},
			{Layer: 0, Position: 13, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mode left + right side
		"testdata/zmk/command-a.log",
		[]KeyPress{
			{Layer: 0, Position: 19, Count: 1},
			{Layer: 0, Position: 10, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// multiple presses
		"testdata/zmk/multiple-presses.log",
		[]KeyPress{
			{Layer: 0, Position: 12, Count: 3},
			{Layer: 0, Position: 13, Count: 1},
			{Layer: 0, Position: 6, Count: 2},
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 7, Count: 3},
		},
		[]ComboPress{},
	},
	{
		// combo-test, once L, once N, then combo of L+N=+
		"testdata/zmk/combo-test-on-right-side.log",
		[]KeyPress{
			{Layer: 0, Position: 6, Count: 2},
			{Layer: 0, Position: 16, Count: 2},
		},
		[]ComboPress{
			{
				Key:    "+",
				Number: 61,
				Count:  2,
			},
		},
	},
	{
		// combo-test, once B, once G, then combo of B+G=%
		"testdata/zmk/combo-test-on-left-side.log",
		[]KeyPress{
			{Layer: 0, Position: 4, Count: 1},
			{Layer: 0, Position: 14, Count: 1},
		},
		[]ComboPress{
			{
				Key:    "%",
				Number: 55,
				Count:  1,
			},
		},
	},
	{
		// combo first (, then <
		"testdata/zmk/combo_without_hold_then_with.log",
		[]KeyPress{
			{Layer: 0, Position: 33, Count: 1},
		},
		// TODO: this is wrong, we first press ( and then <
		[]ComboPress{
			{
				Key:    "(",
				Number: 44,
				Count:  2,
			},
		},
	},
}

func parseTestFile(filename string) Parser {
	file, err := os.Open("../../" + filename)

	if err != nil {
		log.Fatal(err)
	}

	keyMap, err := keymap.Load("../../testdata/zmk/keymap.yaml")
	if err != nil {
		log.Fatal(err)
	}

	parser := *NewParser(keyMap)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parser.Parse(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return parser
}

func TestParser(t *testing.T) {
	for _, test := range countTests {
		assert.Equal(t, test.expectedKeys, parseTestFile(test.filename).KeyPresses)
		assert.Equal(t, test.expectedCombos, parseTestFile(test.filename).ComboPresses)
	}
}
