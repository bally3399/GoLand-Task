package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	expected := 15
	num := 5
	num1 := 10
	got := add(num, num1)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}

}
