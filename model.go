package main

// Generator is object of generator
type Generator struct {
	Set       int        `json:"set"`
	Sep       string     `json:"sep"`
	Templates []Template `json:"templates"`
}

// Template is template of structure.
type Template struct {
	Min  int64 `json:"min"`
	Max  int64 `json:"max"`
	Rows int   `json:"rows"`
	Cols int   `json:"cols"`
}
