package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x := 4
	y := 2
	z := Add(x, y)
	result := 6
	assert.EqualValues(t, result, z)

}

func TestNotEqual(t *testing.T) {
	t.Run("Should be return error when input equal error", func(t *testing.T) {
		x := "error"
		y := NotEqual(x)
		e := "incorrect"
		assert.EqualError(t, y, e)
	})
	t.Run("Should be return nil when input not equal error", func(t *testing.T) {
		x := "not error"
		y := NotEqual(x)
		assert.NoError(t, y)
	})
}
