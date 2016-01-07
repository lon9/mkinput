package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func main() {

	// Reading Json rom stdin.
	dec := json.NewDecoder(os.Stdin)
	var g Generator
	dec.Decode(&g)
	fmt.Printf("%+v\n", g)

	file, err := os.OpenFile("template.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	for i := 0; i < g.Set; i++ {
		for _, t := range g.Templates {
			for j := 0; i < t.Rows; i++ {
				row := ""
				for k := 0; k < t.Cols; k++ {

				}
				row += "\n"
			}

		}
		writer.Write([]byte(g.Sep))
	}

}
