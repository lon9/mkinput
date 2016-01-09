package main

import (
	"math/rand"
	"strconv"
	"time"
)

func randIntString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min+1))
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
