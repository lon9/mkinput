package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

	file, err := os.OpenFile("input.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(file)
	defer file.Close()
	file.Truncate(0)

	for i := 0; i < g.Set; i++ {
		for _, t := range g.Templates {
			rows := randInt(t.MinRows, t.MaxRows)
			if t.RowSize {
				fmt.Fprint(w, strconv.Itoa(rows)+"\n")
			}
			for j := 0; j < rows; j++ {
				cols := randInt(t.MinCols, t.MaxCols)
				if t.ColSize {
					fmt.Fprint(w, strconv.Itoa(cols)+"\n")
				}
				vals := make([]string, cols)
				for k := 0; k < cols; k++ {
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

func randIntString(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(min + rand.Intn(max-min+1))
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
