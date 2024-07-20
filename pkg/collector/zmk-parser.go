package collector

import (
	"strconv"
	"strings"
	. "zmk-heatmap/pkg/heatmap"
	"zmk-heatmap/pkg/keycodes"
	"zmk-heatmap/pkg/keymap"
)

type ZmkParser struct {
	KeyMap               *keymap.Keymap
	numberOfKeysInKeyMap int
	pressed              map[string]bool
	shiftIsDown          bool
}

func (p *ZmkParser) Parse(logLine string, heatmap *Heatmap) (err error) {
	logLine = strings.Trim(logLine, "\n")
	tokens := strings.Split(logLine, " ")

	if len(tokens) < 5 {
		return
	}

	messageType := strings.TrimRight(tokens[3], ":")
	messageTokens := tokens[4:]

	switch messageType {
	case "zmk_hid_register_mod", "zmk_hid_unregister_mod":
		if !strings.Contains(logLine, "Modifiers set to") {
			break
		}

		modHexString := messageTokens[len(messageTokens)-1]
		mods, err := strconv.ParseInt(modHexString[2:], 16, 16)
		if err != nil {
			return err
		}

		p.shiftIsDown = (mods&keycodes.ModifierLeftShift > 0) || (mods&keycodes.ModifierRightShift > 0)
		break

	case "zmk_keymap_apply_position_state":
		layer, position, err := p.parseKeyPress(messageTokens[1], messageTokens[3])
		if err != nil {
			return err
		}

		// Keypresses show up twice in the apply Position state, once as the
		// key down and once as a key up, but that info is not in the logLine,
		// so we need to keep track of that.
		pressString := messageTokens[1] + messageTokens[3]
		if p.pressed[pressString] {
			p.pressed[pressString] = false
			heatmap.RegisterKeyPress(layer, position, getPressType(p.shiftIsDown))
		} else {
			p.pressed[pressString] = true
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

		if virtualKeyNumber < p.KeyMap.NumberOfKeys() {
			break
		}

		heatmap.RegisterComboPress(virtualKeyNumber-p.KeyMap.NumberOfKeys(), getPressType(p.shiftIsDown))

		//fmt.Println(messageTokens)
		// key := p.KeyMap.Combos[virtualKeyNumber-p.KeyMap.NumberOfKeys].Key
		// fmt.Println("ComboPress press:", virtualKeyNumber, "(", key, ")")
		break
	}

	return
}

func getPressType(shiftIsDown bool) PressType {
	if shiftIsDown {
		return Shifted
	}

	return Tap
}

func (p ZmkParser) parseKeyPress(layerString string, positionString string) (layer int, position int, err error) {
	layer, err = strconv.Atoi(layerString)
	if err != nil {
		return
	}

	position, err = strconv.Atoi(strings.TrimRight(positionString, ":,"))
	if err != nil {
		return
	}

	return
}

func NewZmkLogParser(keyMap *keymap.Keymap) *ZmkParser {
	return &ZmkParser{
		KeyMap:               keyMap,
		pressed:              make(map[string]bool),
		shiftIsDown:          false,
		numberOfKeysInKeyMap: keyMap.NumberOfKeys(),
	}
}
