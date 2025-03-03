package chance

import (
	"math/rand"
	"time"
)

// import "fmt"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

var Animals = []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	// Copy default Animals
	x := make([]string, len(Animals))
	copy(x, Animals)
	// Shuffle
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	// fmt.Printf("x:%v\n", x)
	return x
}
