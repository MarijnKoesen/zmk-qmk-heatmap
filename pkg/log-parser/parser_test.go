package log_parser

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
		"testdata/q.log",
		[]KeyPress{
			{Layer: 0, Position: 0, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// primary Layer, simple key press on left side
		"testdata/w.log",
		[]KeyPress{
			{Layer: 0, Position: 1, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// primary Layer, simple key press on right side
		"testdata/l.log",
		[]KeyPress{
			{Layer: 0, Position: 6, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on left side
		"testdata/semicolon.log",
		[]KeyPress{
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 13, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on left side
		"testdata/parentheses-close.log",
		[]KeyPress{
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 11, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// secondary Layer on right side
		"testdata/1.log",
		[]KeyPress{
			{Layer: 0, Position: 30, Count: 1},
			{Layer: 1, Position: 26, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mode left side
		"testdata/ctrl-c.log",
		[]KeyPress{
			{Layer: 0, Position: 17, Count: 1},
			{Layer: 0, Position: 22, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mode left side
		"testdata/command-t.log",
		[]KeyPress{
			{Layer: 0, Position: 10, Count: 1},
			{Layer: 0, Position: 13, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// home row mode left + right side
		"testdata/command-a.log",
		[]KeyPress{
			{Layer: 0, Position: 19, Count: 1},
			{Layer: 0, Position: 10, Count: 1},
		},
		[]ComboPress{},
	},
	{
		// multiple presses
		"testdata/multiple-presses.log",
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
		"testdata/combo-test-on-right-side.log",
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
		"testdata/combo-test-on-left-side.log",
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
}

func parseTestFile(filename string) Parser {
	file, err := os.Open("../../" + filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keyMap, err := keymap.Load("../../testdata/keymap.yaml")
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
