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
	rawKeymap := &rawKeymap{}

	yamlContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read keymap '%s': %w", filename, err)
	}

	var body interface{}
	if err := yaml.Unmarshal(yamlContent, &body); err != nil {
		return nil, fmt.Errorf("cannot read yaml from '%s': %w", filename, err)
	}

	body = convertYamlToJson(body)
	if b, err := json.Marshal(body); err != nil {
		return nil, fmt.Errorf("cannot concert yaml to json '%s': %w", filename, err)
	} else {
		err := json.Unmarshal(b, rawKeymap)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarhsal json '%s': %w", filename, err)
		}
	}

	for layerName, rows := range rawKeymap.Layers {
		layer := Layer{
			Name: layerName,
			Rows: make([]Row, 0),
		}

		for r, row := range rows {
			layer.Rows = append(layer.Rows, Row{
				Keys: make([]Key, 0),
			})

			// Try to decode as a list of keys
			var keys []json.RawMessage
			err = json.Unmarshal(row, &keys)
			if err == nil {
				for _, rawKey := range keys {
					layer.Rows[r].Keys = append(layer.Rows[r].Keys, mapKey(rawKey))
				}

				continue
			}

			layer.Rows[r].Keys = append(layer.Rows[r].Keys, mapKey(row))
		}

		keymap.Layers = append(keymap.Layers, layer)
	}

	for _, combo := range rawKeymap.Combos {
		layers := make([]string, len(combo.Layers))
		if combo.Layers != nil {
			layers = combo.Layers
		}

		keymap.Combos = append(keymap.Combos, Combo{
			Keys:   combo.Keys,
			Layers: layers,
			Key:    mapKey(combo.Key),
		})
	}

	return keymap, nil
}

func mapKey(rawMessage json.RawMessage) Key {
	// Try to decode as string first
	var key string
	err := json.Unmarshal(rawMessage, &key)
	if err == nil {
		return Key{Tap: key}
	}

	// Try to decode as int
	var keyInt int
	err = json.Unmarshal(rawMessage, &keyInt)
	if err == nil {
		return Key{Tap: strconv.Itoa(keyInt)}
	}

	// If that fails it must be an object
	var keyObj map[string]string
	err = json.Unmarshal(rawMessage, &keyObj)
	if err == nil {
		return Key{
			Tap:     keyObj["t"],
			Hold:    keyObj["h"],
			Shifted: keyObj["s"],
		}
	}

	jsonString, err := rawMessage.MarshalJSON()
	if err != nil {
		panic("invalid json: " + err.Error())
	}
	panic("cannot map key: " + string(jsonString))
}

func convertYamlToJson(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convertYamlToJson(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convertYamlToJson(v)
		}
	}
	return i
}
