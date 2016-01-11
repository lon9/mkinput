package main

import (
	"math/rand"
	"strconv"
	"time"
)

func randFloatString(min, max float64, dig uint) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatFloat(min+rand.Float64()*(max-min), 'f', int(dig), 64)
}

func randIntString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min+1))
}

func randInt(min, max uint) int {
	rand.Seed(time.Now().UnixNano())
	return int(min) + rand.Intn(int(max-min+1))
}
