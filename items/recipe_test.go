package items

import "testing"

func TestValidateAllRecipes(t *testing.T) {
	if err := validateRecipes(); err != nil {
		t.Fatal(err)
	}
}
