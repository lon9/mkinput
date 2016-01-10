package main

import (
	"encoding/json"
	"os"
)

func main() {

	// Reading Json rom stdin.
	dec := json.NewDecoder(os.Stdin)
	var g Generator
	err := dec.Decode(&g)
	if err != nil {
		panic(err)
	}
	g.Generate()
}
