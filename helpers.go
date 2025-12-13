package main

import "math/rand"

func getRandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func getRandFloat(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
