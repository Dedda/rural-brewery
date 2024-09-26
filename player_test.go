package main

import (
	"github.com/Dedda/rural-brewery/items"
	"github.com/go-test/deep"
	"testing"
)

var testItem = items.Item{
	Id:           123,
	Name:         "Test item",
	MaxStackSize: 12,
}

func TestAddInventoryItems(t *testing.T) {
	first := &InventoryItem{
		item:   &testItem,
		amount: 5,
	}
	second := &InventoryItem{
		item:   &testItem,
		amount: 9,
	}
	expectedFirst := &InventoryItem{
		item:   &testItem,
		amount: 12,
	}
	expectedRemainder := &InventoryItem{
		item:   &testItem,
		amount: 2,
	}
	actualRemainder := first.AddInventoryItem(second)
	if diff := deep.Equal(expectedFirst, first); diff != nil {
		t.Fatalf("Filled up inventory slot is wrong: %s", diff)
	}
	if diff := deep.Equal(expectedRemainder, actualRemainder); diff != nil {
		t.Fatalf("Remainder is wrong: %s", diff)
	}
}
