package main

import (
	"fmt"
	"os"
)

func main() {

	// Reading Json rom stdin.
	g, err := NewGeneratorFromStdin(os.Stdin)
	if err != nil {
		switch err {
		case ErrJSON:
			fmt.Println(ErrJSON)
			os.Exit(1)
		case ErrInvalidRange:
			fmt.Println(ErrInvalidRange)
			os.Exit(1)
		default:
			fmt.Println(err)
			os.Exit(1)
		}
	}
	g.Generate()
}
