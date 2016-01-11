package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
	Dig     uint    `json:"dig"`
	MinRows uint    `json:"minRows"`
	MaxRows uint    `json:"maxRows"`
	MinCols uint    `json:"minCols"`
	MaxCols uint    `json:"maxCols"`
	Sep     string  `json:"sep"`
	RowSize bool    `json:"rowSize"`
	ColSize bool    `json: "colSize"`
}

var (

	// ErrJSON is raised when failed to decode JSON.
	ErrJSON = errors.New("Failed to decode JSON.")

	// ErrInvalidRange is raised when maximum value was greater than maximum value.
	ErrInvalidRange = errors.New("Maximum value must be equal or greater than minimum value.")

	// ErrInvalidMin is raised when minRows or minCols is 0.
	ErrInvalidMin = errors.New("minRows and minCols must be greater than 1")
)

// NewGeneratorFromStdin constructs Generator from stdin.
func NewGeneratorFromStdin(stdin io.Reader) (g *Generator, err error) {
	dec := json.NewDecoder(stdin)
	if err = dec.Decode(&g); err != nil {
		err = ErrJSON
		return
	}
	if err = g.checkMinMax(); err != nil {
		return
	}
	if err = g.checkMinRowsCols(); err != nil {
		return
	}
	return
}

func (g *Generator) checkMinMax() error {
	for _, t := range g.Templates {
		if t.Max < t.Min || t.MaxRows < t.MinRows || t.MaxCols < t.MinCols {
			return ErrInvalidRange
		}
	}
	return nil
}

func (g *Generator) checkMinRowsCols() error {
	for _, t := range g.Templates {
		if t.MinRows == 0 || t.MinCols == 0 {
			return ErrInvalidMin
		}
	}
	return nil
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
					vals[k] = randFloatString(t.Min, t.Max, t.Dig)
				}
				fmt.Println(strings.Join(vals, t.Sep))
			}

		}
		fmt.Println(g.Sep)
	}
}
