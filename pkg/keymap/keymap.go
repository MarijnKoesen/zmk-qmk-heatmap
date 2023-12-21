package keymap

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Combo struct {
	Keys   []int    `yaml:"p"`
	Key    string   `yaml:"k"`
	Layers []string `yaml:"l"`
}

type Keymap struct {
	NumberOfKeys int
	Combos       []Combo
}

//func (c Combo) String() string {
//	return "Combo: " + c.Key
//}

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

	layers := obj["layers"].(map[string]interface{})
	firstLayer := layers[keys(layers)[0]].([]interface{})
	for _, row := range firstLayer {
		keymap.NumberOfKeys += len(row.([]interface{}))
	}

	combos := obj["combos"].([]interface{})
	for _, combo := range combos {
		combo2 := combo.(map[string]interface{})

		keymap.Combos = append(keymap.Combos, Combo{
			Keys:   convertSlice[int](combo2["p"].([]interface{})),
			Key:    combo2["k"].(string),
			Layers: convertSlice[string](combo2["l"].([]interface{})),
		})
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

	return keys
}
