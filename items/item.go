package items

import "encoding/json"

var allItems map[int]Item

func AllItems() *map[int]Item {
	if allItems == nil {
		if err := json.Unmarshal(allItems_json, &allItems); err != nil {
			panic(err)
		}
	}
	return &allItems
}

type Item struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	MaxStackSize int    `json:"maxStackSize"`
}
