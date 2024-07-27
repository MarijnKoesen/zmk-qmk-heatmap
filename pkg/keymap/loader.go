package keymap

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"gopkg.in/yaml.v3"
)

type rawKeymap struct {
	Layers map[string][]interface{}
	Combos []rawCombo `yaml:"combos"`
}

type rawYaml struct {
	Layers yaml.Node
}

type rawCombo struct {
	Keys   []int       `yaml:"p"`
	Layers []string    `yaml:"l"`
	Key    interface{} `yaml:"k"`
}

type rawKey struct {
	Tap     string `json:"t"`
	Hold    string `json:"h"`
	Shifted string `json:"s"`
}

func Load(filename string, numberOfSensors int) (*Keymap, error) {
	keymap := New(numberOfSensors)
	rawKeymap := &rawKeymap{}

	yamlContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read keymap '%s': %w", filename, err)
	}

	if err := yaml.Unmarshal(yamlContent, &rawKeymap); err != nil {
		return nil, fmt.Errorf("cannot read yaml from '%s': %w", filename, err)
	}

	for layerName, rawRows := range rawKeymap.Layers {
		layer := Layer{
			Name: layerName,
			Rows: make([]*Row, 0),
		}

		for _, rawRow := range rawRows {
			row := &Row{
				Keys: make([]*Key, 0),
			}
			layer.Rows = append(layer.Rows, row)

			switch a := rawRow.(type) {
			case string:
				row.Append(&Key{Tap: a})
			case []interface{}:
				for _, rawKey := range a {
					row.Append(mapKey(rawKey))
				}
			case map[string]interface{}:
				row.Append(mapKey(a))
			default:
				panic(fmt.Sprintf("unsupported keymap row(%T): %v", a, rawRow))
			}
		}

		keymap.Layers = append(keymap.Layers, &layer)
	}

	layerOrder := getLayerOrder(yamlContent)
	sort.Slice(keymap.Layers, func(i, j int) bool {
		return layerOrder[keymap.Layers[i].Name] < layerOrder[keymap.Layers[j].Name]
	})

	for _, combo := range rawKeymap.Combos {
		layers := make([]string, len(combo.Layers))
		if combo.Layers != nil {
			layers = combo.Layers
		}

		keymap.Combos = append(keymap.Combos, &Combo{
			Keys:   combo.Keys,
			Layers: layers,
			Key:    mapKey(combo.Key),
		})
	}

	return keymap, nil
}

func mapKey(keyData interface{}) *Key {
	switch key := keyData.(type) {
	case string:
		return &Key{Tap: key}
	case int:
		return &Key{Tap: strconv.Itoa(key)}
	case map[string]interface{}:
		jbytes, err := json.Marshal(key)
		if err != nil {
			panic(fmt.Sprintf("unsupported key in keymap: %v", keyData))
		}

		rawKeyObj := &rawKey{}
		err = json.Unmarshal(jbytes, rawKeyObj)
		if err != nil {
			panic(fmt.Sprintf("unsupported key in keymap: %v", keyData))
		}

		return &Key{
			Tap:     rawKeyObj.Tap,
			Hold:    rawKeyObj.Hold,
			Shifted: rawKeyObj.Shifted,
		}
	case nil:
		return &Key{}
	default:
		panic(fmt.Sprintf("unsupported key in keymap: %v", keyData))
	}
}

func getLayerOrder(yamlContent []byte) map[string]int {
	// In GoLang the order of maps is not guaranteed, but we must read the layers in order, because the keyboard
	// outputs key presses with the layer number. If we were to randomly get the layer we'd apply the wrong layer
	// in the heatmap. So use `yaml.Node` which actually makes it possible to unmarshall in order.
	order := make(map[string]int)

	obj := &rawYaml{}
	err := yaml.Unmarshal(yamlContent, obj)
	if err != nil {
		panic("cannot get layer order")
	}

	for i, node := range obj.Layers.Content {
		if node.Tag == "!!str" {
			order[node.Value] = i
		}
	}

	return order
}
