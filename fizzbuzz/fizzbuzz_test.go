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
func TestFizzbuzz2(t *testing.T) {
	number := 2
	result := Fizzbuzz(number)
	expect := "2"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz3(t *testing.T) {
	number := 3
	result := Fizzbuzz(number)
	expect := "fizz"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz4(t *testing.T) {
	number := 4
	result := Fizzbuzz(number)
	expect := "4"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz5(t *testing.T) {
	number := 5
	result := Fizzbuzz(number)
	expect := "buzz"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz12(t *testing.T) {
	number := 12
	result := Fizzbuzz(number)
	expect := "12"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz15(t *testing.T) {
	number := 15
	result := Fizzbuzz(number)
	expect := "fizzbuzz"

	assert.EqualValues(t, expect, result)
}
