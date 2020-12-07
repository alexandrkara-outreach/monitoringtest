package util

import "math/rand"

// Lucky generates a random boolean with true having a probability of rate (0 <= rate <= 1).
func Lucky(rate float32) bool {
	return rate > rand.Float32()
}
