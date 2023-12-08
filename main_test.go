package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	x := 4
	y := 2
	z := Add(x, y)
	result := 6
	if z != result {
		t.Errorf("incorrect,the number %d + %d result should be %d", x, y, result)
	}
}
