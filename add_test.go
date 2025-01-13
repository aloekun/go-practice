package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Arrange
	a := 2
	b := 3
	expected := 5

	// Act
	result := Add(a, b)

	// Assert
	assert.Equal(t, expected, result, "they should be equal")
}
