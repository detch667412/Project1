package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz1(t *testing.T) {
	number := 1
	result := Fizzbuzz(number)
	expect := "1"

	assert.EqualValues(t, expect, result)
}
