package main

import (
	"encoding/json"
	"flag"
	"os"
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
