package main

import (
	"fmt"
	"strings"
)

// Generator is object of generator
type Generator struct {
	Set       int        `json:"set"`
	Sep       string     `json:"sep"`
	Templates []Template `json:"templates"`
}

// Template is template of structure.
type Template struct {
	Min     int    `json:"min"`
	Max     int    `json:"max"`
	MinRows int    `json:"minRows"`
	MaxRows int    `json:"maxRows"`
	MinCols int    `json:"minCols"`
	MaxCols int    `json:"maxCols"`
	Sep     string `json:"sep"`
	RowSize bool   `json:"rowSize"`
	ColSize bool   `json: "colSize"`
}

// Generate generates input.txt.
func (g *Generator) Generate() {

	for i := 0; i < g.Set; i++ {
		for _, t := range g.Templates {
			rows := randInt(t.MinRows, t.MaxRows)
			if t.RowSize {
				fmt.Println(rows)
			}
			for j := 0; j < rows; j++ {
				cols := randInt(t.MinCols, t.MaxCols)
				if t.ColSize {
					fmt.Println(cols)
				}
				vals := make([]string, cols)
				for k := 0; k < cols; k++ {
					v := randIntString(t.Min, t.Max)
					vals[k] = v
				}
				row := strings.Join(vals, t.Sep)
				fmt.Println(row)
			}

		}
		fmt.Println(g.Sep)
	}

}
