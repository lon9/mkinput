package main

// Generator is object of generator
type Generator struct {
	Set       int        `json:"set"`
	Sep       string     `json:"sep"`
	Templates []Template `json:"templates"`
}

// Template is template of structure.
type Template struct {
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	MinRows int `json:"minRows"`
	MaxRows int `json:"maxRows"`
	MinCols int `json:"minCols"`
	MaxCols int `json:"maxCols"`
	Sep  string `json:"sep"`
	Size bool   `json:"size"`
}
