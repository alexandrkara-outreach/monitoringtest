package util

import "math/rand"

func Lucky(rate float32) bool {
	return rate > rand.Float32()
}
