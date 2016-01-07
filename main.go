package main

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"fmt"
	"time"
)

func main() {

	// Reading Json rom stdin.
	dec := json.NewDecoder(os.Stdin)
	var g Generator
	dec.Decode(&g)

	file, err := os.OpenFile("input.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(file)
	defer file.Close()
	file.Truncate(0)

	for i := 0; i < g.Set; i++ {
		for _, t := range g.Templates {
			if t.Size{
				fmt.Fprint(w, strconv.Itoa(t.Cols) + "\n")
			}
			for j := 0; j < t.Rows; j++ {
				vals := make([]string, t.Cols)
				for k := 0; k < t.Cols; k++ {
					v := randIntString(t.Min, t.Max)
					vals[k] = v
				}
				row := strings.Join(vals, t.Sep)
				fmt.Fprint(w, row+"\n")
			}

		}
		fmt.Fprint(w, g.Sep+"\n")
	}
	w.Flush()

}

func randIntString(min, max int)string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min))
}
