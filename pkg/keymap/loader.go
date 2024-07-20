package keymap

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

type rawKeymap struct {
	NumberOfKeys int
	Layers       map[string][]json.RawMessage
	Combos       []rawCombo
}

type rawCombo struct {
	Keys   []int           `json:"p"`
	Layers []string        `json:"l"`
	Key    json.RawMessage `json:"k"`
}

func Load(filename string) (*Keymap, error) {
	keymap := New()

	// First load the keymap YAML
	yamlBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read keymap from '%s': %w", filename, err)
	}

	var yamlObj interface{}
	if err := yaml.Unmarshal(yamlBytes, &yamlObj); err != nil {
		return nil, fmt.Errorf("cannot read yaml from '%s': %w", filename, err)
	}

	// Then transform it into json because with the json unmarshall we can unmarshall parts into json.RawMessage.
	// This is needed so we can figure out what type the fields are. With YAML we wouldn't be able to unmarshall
	// because for different types in the YAML (e.g. string, array representations of a key)
	jsonBytes, err := json.Marshal(yamlObj)
	if err != nil {
		return nil, fmt.Errorf("cannot convert '%s' yaml to json: %w", filename, err)
	}

	rawKeymap := &rawKeymap{}
	err = json.Unmarshal(jsonBytes, rawKeymap)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarhsal '%s' json: %w", filename, err)
	}

	for layerName, rows := range rawKeymap.Layers {
		layer := keymap.AddLayer(layerName, []*Row{})
		for _, row := range rows {
			layer.Rows = append(layer.Rows, mapRow(row))
		}
	}

	for _, combo := range rawKeymap.Combos {
		comboLayers := make([]string, len(combo.Layers))
		if combo.Layers != nil {
			comboLayers = combo.Layers
		}

		keymap.Combos = append(keymap.Combos, &Combo{
			Keys:   combo.Keys,
			Layers: comboLayers,
			Key:    mapKey(combo.Key),
		})
	}

	return keymap, nil
}

func mapRow(rawMessage json.RawMessage) *Row {
	jsonString, err := rawMessage.MarshalJSON()
	if err != nil {
		panic("invalid json: " + err.Error())
	}

	row := &Row{
		Keys: []*Key{},
	}

	switch string(jsonString[0]) {
	case "[":
		var rawKeys []json.RawMessage
		err = json.Unmarshal(rawMessage, &rawKeys)

		for _, rawKey := range rawKeys {
			row.AddKey(mapKey(rawKey))
		}
	default:
		row.AddKey(mapKey(rawMessage))
	}

	return row
}

func mapKey(rawMessage json.RawMessage) *Key {
	jsonString, err := rawMessage.MarshalJSON()
	if err != nil {
		panic("invalid json: " + err.Error())
	}

	switch string(jsonString[0]) {
	case "{":
		var keyObj map[string]string
		err = json.Unmarshal(rawMessage, &keyObj)
		if err == nil {
			return &Key{
				Tap:     keyObj["t"],
				Hold:    keyObj["h"],
				Shifted: keyObj["s"],
			}
		}

		panic("cannot map key from object: " + string(jsonString))
	default:
		// Try to decode as string first
		var key string
		err := json.Unmarshal(rawMessage, &key)
		if err == nil {
			return &Key{Tap: key}
		}

		// Try to decode as int
		var keyInt int
		err = json.Unmarshal(rawMessage, &keyInt)
		if err == nil {
			return &Key{Tap: strconv.Itoa(keyInt)}
		}

		panic("cannot map key: " + string(jsonString))
	}
}
