package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randFloatString(min, max float64) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%f", min+rand.Float64()*(max-min))
}

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func randIntString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min+1))
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
