package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	// Reading Json rom stdin.
	dec := json.NewDecoder(os.Stdin)
	var g Generator
	err := dec.Decode(&g)
	if err != nil {
		panic(err)
	}
	g.Generate("input.txt")
}

func randIntString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min+1))
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
