package main

// Generator is object of generator
type Generator struct {
	Set       int64      `json:"set"`
	Sep       string     `json:"sep"`
	Templates []Template `json:"templates"`
}

// Template is template of structure.
type Template struct {
	Min  int64 `json:"min"`
	Max  int64 `json:"max"`
	Rows int64 `json:"rows"`
	Cols int64 `json:"cols"`
}
