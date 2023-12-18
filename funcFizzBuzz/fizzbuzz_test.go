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
	expect := "fizz"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz15(t *testing.T) {
	number := 15
	result := Fizzbuzz(number)
	expect := "fizzbuzz"

	assert.EqualValues(t, expect, result)
}

func TestFizzbuzz20(t *testing.T) {
	number := 20
	result := Fizzbuzz(number)
	expect := "buzz"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz26(t *testing.T) {
	number := 26
	result := Fizzbuzz(number)
	expect := "26"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz199(t *testing.T) {
	number := 199
	result := Fizzbuzz(number)
	expect := "199"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz6(t *testing.T) {
	number := 6
	result := Fizzbuzz(number)
	expect := "fizz"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz30(t *testing.T) {
	number := 30
	result := Fizzbuzz(number)
	expect := "fizzbuzz"

	assert.EqualValues(t, expect, result)
}
func TestFizzbuzz50(t *testing.T) {
	number := 50
	result := Fizzbuzz(number)
	expect := "buzz"

	assert.EqualValues(t, expect, result)
}
