package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAnswer(t *testing.T) {
	sut := NewHitAndBrow(4)

	err := sut.SetAnswer([]int{1, 2, 3, 4})

	assert.Nil(t, err)
	assert.ElementsMatch(t, sut.Answer, []int{1, 2, 3, 4})
}

func TestSetAnswerOtherSize(t *testing.T) {
	sut := NewHitAndBrow(4)

	err := sut.SetAnswer([]int{1, 2, 3})

	assert.NotNil(t, err)
}

func TestRandomAnswer(t *testing.T) {
	sut := NewHitAndBrow(4)

	sut.SetRandomAnswer()

	assert.Len(t, sut.Answer, 4)
	assert.GreaterOrEqual(t, sut.Answer[0], 1)
	assert.LessOrEqual(t, sut.Answer[0], 9)
	assert.GreaterOrEqual(t, sut.Answer[1], 1)
	assert.LessOrEqual(t, sut.Answer[1], 9)
	assert.GreaterOrEqual(t, sut.Answer[2], 1)
	assert.LessOrEqual(t, sut.Answer[2], 9)
	assert.GreaterOrEqual(t, sut.Answer[3], 1)
	assert.LessOrEqual(t, sut.Answer[3], 9)
}

func TestHitAll(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Hit([]int{1, 2, 3, 4})

	assert.Equal(t, 4, result)
}

func TestHitPart(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Hit([]int{1, 2, 5, 6})

	assert.Equal(t, 2, result)
}

func TestHitZero(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Hit([]int{4, 3, 2, 1})

	assert.Equal(t, 0, result)
}

func TestHitZero2(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Hit([]int{5, 6, 7, 8})

	assert.Equal(t, 0, result)
}

func TestBlowAll(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Blow([]int{4, 3, 2, 1})

	assert.Equal(t, 4, result)
}

func TestBlowPart(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Blow([]int{5, 6, 7, 1})

	assert.Equal(t, 1, result)
}

func TestBlowZero(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Blow([]int{5, 6, 7, 8})

	assert.Equal(t, 0, result)
}

func TestBlowZero2(t *testing.T) {
	sut := NewHitAndBrow(4)
	sut.SetAnswer([]int{1, 2, 3, 4})

	result := sut.Blow([]int{1, 2, 3, 4})

	assert.Equal(t, 0, result)
}
