package keymap

import (
	"gopkg.in/yaml.v3"
	"os"
)

// t h type || string

type Combo struct {
	Keys   []int    `yaml:"p"`
	Key    string   `yaml:"k"`
	Layers []string `yaml:"l"`
}

type Keymap struct {
	NumberOfKeys int
	Combos       []Combo
}

func Load(filename string) (keymap Keymap, err error) {
	keymap = Keymap{Combos: []Combo{}}

	obj := make(map[string]interface{})

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		return
	}

	keymap.NumberOfKeys = countNumberOfKeysInLayers(obj["layers"].(map[string]interface{}))

	combos, comboExists := obj["combos"].([]interface{})
	if comboExists {
		for _, combo := range combos {
			combo2 := combo.(map[string]interface{})

			c := Combo{
				Keys:   convertSlice[int](combo2["p"].([]interface{})),
				Key:    combo2["k"].(string),
				Layers: []string{},
			}

			if comboLayers, ok := combo2["l"]; ok {
				c.Layers = convertSlice[string](comboLayers.([]interface{}))
			}

			keymap.Combos = append(keymap.Combos, c)
		}
	}

	return
}

func countNumberOfKeysInLayers(layers map[string]interface{}) (numberOfKeys int) {
	for _, row := range layers[keys(layers)[0]].([]interface{}) {
		switch row.(type) {
		case []interface{}:
			numberOfKeys += len(row.([]interface{}))
			break
		case interface{}:
			numberOfKeys++
			break
		}
	}

	return
}

func convertSlice[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, v.(E))
	}
	return
}

func keys(someMap map[string]interface{}) (keys []string) {
	for key, _ := range someMap {
		keys = append(keys, key)
	}

	return
}
