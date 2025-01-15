package main

import (
	"errors"
	"math/rand/v2"
)

type HitAndBlow struct {
	Size   int
	Answer []int
}

func NewHitAndBrow(size int) HitAndBlow {
	return HitAndBlow{
		Size:   size,
		Answer: make([]int, size),
	}
}

func (h *HitAndBlow) SetAnswer(answer []int) error {
	if len(answer) != h.Size {
		return errors.New("The size of the answer is invalid.")
	}
	h.Answer = answer
	return nil
}

func (h *HitAndBlow) SetRandomAnswer() {
	for i := 0; i < len(h.Answer); i++ {
		randNumber := rand.IntN(9) + 1
		h.Answer[i] = randNumber
	}
}

func (h *HitAndBlow) Hit(guess []int) int {
	hit := 0
	for i := 0; i < len(h.Answer); i++ {
		if h.Answer[i] == guess[i] {
			hit++
		}
	}
	return hit
}

func (h *HitAndBlow) Blow(guess []int) int {
	blow := 0
	for i := 0; i < len(h.Answer); i++ {
		for j := 0; j < len(guess); j++ {
			if i != j && h.Answer[i] == guess[j] {
				blow++
			}
		}
	}
	return blow
}
