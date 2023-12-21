package log_parser

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
	"zmk-heatmap/pkg/keymap"
)

type ComboPress struct {
	Key    string
	Number int
	Count  int
}

type KeyPress struct {
	Layer    int
	Position int
	Count    int
}

type Parser struct {
	KeyMap       keymap.Keymap
	KeyPresses   []KeyPress
	ComboPresses []ComboPress
	pressed      map[string]bool
}

func (a ComboPress) Compare(b ComboPress) int {
	if a.Number > b.Number {
		return 1
	}
	if a.Number < b.Number {
		return -1
	}

	return 0
}

func (k KeyPress) String() string {
	return "KeyPress[Layer: " + strconv.Itoa(k.Layer) + ", Position: " + strconv.Itoa(k.Position) + ", Count: " + strconv.Itoa(k.Count) + "]"
}

func (p *Parser) Parse(line string) error {
	line = strings.Trim(line, "\n")
	tokens := strings.Split(line, " ")

	if len(tokens) < 5 {
		return nil
	}

	messageType := strings.TrimRight(tokens[3], ":")
	messageTokens := tokens[4:]

	switch messageType {
	case "zmk_keymap_apply_position_state":
		press, err := p.getKeyPress(messageTokens[1], messageTokens[3])
		if err != nil {
			return err
		}

		// Keypresses show up twice in the apply Position state, once as the
		// key down and once as a key up, but that info is not in the line,
		// so we need to keep track of that.
		if p.pressed[press.String()] {
			p.pressed[press.String()] = false
			press.Count += 1
		} else {
			p.pressed[press.String()] = true
		}

		// timestamp := strings.Trim(tokens[0], "[]")
		// message := strings.Join(messageTokens, " ")
		// fmt.Println(timestamp, ":", messageType, "=>", message, key)
		break

	case "on_keymap_binding_released":
		virtualKeyNumber, err := strconv.Atoi(messageTokens[1])
		if err != nil {
			return err
		}

		if virtualKeyNumber < p.KeyMap.NumberOfKeys {
			break
		}

		key := p.KeyMap.Combos[virtualKeyNumber-p.KeyMap.NumberOfKeys].Key

		//fmt.Println("ComboPress press:", virtualKeyNumber, "(", key, ")")
		p.handleComboPress(virtualKeyNumber, key)

		break
	}

	return nil
}

func (p *Parser) getKeyPress(layer string, position string) (keyPress *KeyPress, err error) {
	layerNum, err := strconv.Atoi(layer)
	if err != nil {
		return keyPress, err
	}

	pos, err := strconv.Atoi(strings.TrimRight(position, ":,"))
	if err != nil {
		return keyPress, err
	}

	for i := range p.KeyPresses {
		if p.KeyPresses[i].Layer == layerNum && p.KeyPresses[i].Position == pos {
			return &p.KeyPresses[i], err
		}
	}

	keyPress = &KeyPress{Layer: layerNum, Position: pos, Count: 0}
	p.KeyPresses = append(p.KeyPresses, *keyPress)

	return keyPress, err
}

func (p *Parser) handleComboPress(number int, key string) {
	for i := range p.ComboPresses {
		if p.ComboPresses[i].Number == number && p.ComboPresses[i].Key == key {
			p.ComboPresses[i].Count += 1
			return
		}
	}

	p.ComboPresses = append(
		p.ComboPresses,
		ComboPress{
			Key:    key,
			Number: number,
			Count:  1,
		},
	)
}

func NewParser(keyMap keymap.Keymap) *Parser {
	return &Parser{
		KeyMap:       keyMap,
		KeyPresses:   []KeyPress{},
		ComboPresses: []ComboPress{},
		pressed:      make(map[string]bool),
	}
}

func LoadParser(filename string) (*Parser, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	bytes, _ := io.ReadAll(jsonFile)
	var parser = &Parser{
		pressed: make(map[string]bool),
	}

	json.Unmarshal(bytes, parser)
	if err != nil {
		return nil, err
	}

	return parser, nil
}
