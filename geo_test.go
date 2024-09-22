package main

import (
	"github.com/go-test/deep"
	"testing"
)

func TestCenterHorizontally(t *testing.T) {
	expected := 320.0
	actual := CenterWidthHorizontally(40, 680)
	if diff := deep.Equal(expected, actual); diff != nil {
		t.Fatal(diff)
	}
}
