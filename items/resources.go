package items

import (
	_ "embed"
)

var (
	//go:embed items.json
	allItems_json []byte
	//go:embed recipes.json
	allRecipes_json []byte
)
