package collector

import (
	"fmt"
	"strconv"
	"strings"
	. "zmk-heatmap/pkg/heatmap"
	"zmk-heatmap/pkg/keycodes"
	"zmk-heatmap/pkg/keymap"
)

type state_ struct {
	keyState keyState
	layer    int
}

type keyState int

const (
	stateUp keyState = iota
	stateDown
	stateHold
	stateShifted
)

type ZmkParser struct {
	KeyMap               *keymap.Keymap
	numberOfKeysInKeyMap int
	currentLayer         int
	state                map[int]*state_
	shiftIsDown          bool
}

func (p *ZmkParser) Parse(logLine string, heatmap *Heatmap) (err error) {
	logLine = strings.Trim(logLine, "\n")
	logLine = strings.ReplaceAll(logLine, "\u001B[0m", "")
	tokens := strings.Split(logLine, " ")

	if len(tokens) < 5 {
		return
	}

	messageType := strings.TrimRight(tokens[3], ":")
	messageTokens := tokens[4:]

	switch messageType {
	case "zmk_hid_register_mod", "zmk_hid_unregister_mod":
		// Handle pressing of shift
		if !strings.Contains(logLine, "Modifiers set to") {
			break
		}

		modHexString := messageTokens[len(messageTokens)-1]
		mods, err := strconv.ParseInt(modHexString[2:], 16, 16)
		if err != nil {
			return err
		}

		p.shiftIsDown = (mods&keycodes.ModifierLeftShift > 0) || (mods&keycodes.ModifierRightShift > 0)

	case "set_layer_state":
		// Handle layer switching
		layer, err := strconv.Atoi(messageTokens[2])
		if err != nil {
			return err
		}

		if messageTokens[4] == "1" {
			p.currentLayer = layer
		} else {
			p.currentLayer = 0
		}

	case "decide_hold_tap":
		// Handle keys going into the 'hold' state
		position, err := strconv.Atoi(messageTokens[0])
		if err != nil {
			return err
		}

		if messageTokens[1] == "decided" && messageTokens[2] == "hold-timer" {
			p.state[position].keyState = stateHold
		}

	case "zmk_keymap_apply_position_state":
		// This is the normal/main handling of keys presses
		layer, position, err := p.parseApplyPositionStateKeyPress(messageTokens[1], messageTokens[3])
		if err != nil {
			return err
		}

		if p.state[position] == nil {
			p.registerPress(position, layer)
		} else {
			heatmap.RegisterKeyPress(
				p.state[position].layer,
				position,
				getPressType(p.state[position].keyState),
			)

			p.state[position] = nil
		}

	case "on_keymap_binding_pressed":
		// Handle combo presses
		var position int
		layer := p.currentLayer

		if messageTokens[0] == "layer" {
			// Handle combo presses on the peripheral side
			layer, position, err = p.parseApplyPositionStateKeyPress(messageTokens[1], messageTokens[3])
			if err != nil {
				return err
			}
		} else if messageTokens[0] == "position" {
			// Handle combo presses on the main side
			position, err = strconv.Atoi(messageTokens[1])
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unsupported on_keymap_binding_pressed line")
		}

		if position < p.KeyMap.NumberOfKeys() {
			break
		}

		if p.state[position] == nil {
			// With a combo we don't always get the zmk_keymap_apply_position_state, e.g. with combo-right-hold-shift-ctrl-i.log
			p.registerPress(position, layer)
		}

	case "on_keymap_binding_released":
		// Handle combo releases
		virtualKeyNumber, err := strconv.Atoi(messageTokens[1])
		if err != nil {
			return err
		}

		if virtualKeyNumber < p.KeyMap.NumberOfKeys() {
			break
		}

		combo, err := p.KeyMap.ComboByPosition(virtualKeyNumber)
		if err != nil {
			return err
		}

		heatmap.RegisterComboPress(p.currentLayer, combo.Keys, getPressType(p.state[virtualKeyNumber].keyState))
		p.state[virtualKeyNumber] = nil

	case "on_hold_tap_binding_pressed":
		// With a combo on the peripheral half we don't always get the zmk_keymap_apply_position_state
		// For an example see combo-right-hold-shift-ctrl-i.log, so we need to handle that here
		virtualKeyNumber, err := strconv.Atoi(messageTokens[0])
		if err != nil {
			return err
		}

		if p.state[virtualKeyNumber] == nil {
			p.registerPress(virtualKeyNumber, p.currentLayer)
		}

	case "on_hold_tap_binding_released":
		virtualKeyNumber, err := strconv.Atoi(messageTokens[0])
		if err != nil {
			return err
		}

		p.state[virtualKeyNumber] = nil
	}

	return
}

func (p *ZmkParser) registerPress(position int, layer int) {
	p.state[position] = &state_{
		keyState: stateDown,
		layer:    p.currentLayer,
	}

	if p.shiftIsDown {
		p.state[position].keyState = stateShifted
	}
}

func getPressType(state keyState) PressType {
	switch state {
	case stateShifted:
		return Shifted
	case stateHold:
		return Hold
	default:
		return Tap
	}
}

func (p ZmkParser) parseApplyPositionStateKeyPress(layerString string, positionString string) (layer int, position int, err error) {
	layer, err = strconv.Atoi(layerString)
	if err != nil {
		return
	}

	position, err = strconv.Atoi(strings.TrimRight(positionString, ":,"))
	return
}

func (p ZmkParser) parseKScanKeyPress(rowString string, colString string, positionString string, pressedString string) (row int, col int, position int, pressed bool, err error) {
	row, err = strconv.Atoi(strings.TrimRight(rowString, ","))
	if err != nil {
		return
	}

	col, err = strconv.Atoi(strings.TrimRight(colString, ":,"))
	if err != nil {

	}
	position, err = strconv.Atoi(strings.TrimRight(positionString, ":,"))
	if err != nil {
		return
	}

	pressed, err = strconv.ParseBool(strings.TrimRight(pressedString, ":,"))
	if err != nil {
		return
	}

	return row, col, position, pressed, nil
}

func NewZmkLogParser(keyMap *keymap.Keymap) *ZmkParser {
	return &ZmkParser{
		KeyMap:               keyMap,
		currentLayer:         0,
		state:                make(map[int]*state_),
		shiftIsDown:          false,
		numberOfKeysInKeyMap: keyMap.NumberOfKeys(),
	}
}
