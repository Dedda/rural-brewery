package main

import (
	"github.com/go-test/deep"
	"testing"
)

type TestItem struct{}

func (i *TestItem) Id() int {
	return 1
}

func (i *TestItem) Name() string {
	return "TestItem"
}

func (i *TestItem) MaxStackSize() int {
	return 12
}

func TestAddInventoryItems(t *testing.T) {
	first := &InventoryItem{
		item:   &TestItem{},
		amount: 5,
	}
	second := &InventoryItem{
		item:   &TestItem{},
		amount: 9,
	}
	expectedFirst := &InventoryItem{
		item:   &TestItem{},
		amount: 12,
	}
	expectedRemainder := &InventoryItem{
		item:   &TestItem{},
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
