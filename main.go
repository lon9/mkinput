package main

import (
	"encoding/json"
	"flag"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var filepath = flag.String("f", "input.txt", "filepath for generated file")
	flag.Parse()
	if filepath == nil {
		*filepath = "input.txt"
	}

	// Reading Json rom stdin.
	dec := json.NewDecoder(os.Stdin)
	var g Generator
	err := dec.Decode(&g)
	if err != nil {
		panic(err)
	}
	err = g.Generate(*filepath)
	if err != nil {
		panic(err)
	}
}

func randIntString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min+1))
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
