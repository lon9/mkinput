package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

func main() {
	d := make(map[string]string)
	d["a"] = "muda"
	d["b"] = "dora"
	d["c"] = "bora"

	const t = "{{.a}} = {{.b}} + {{.c}}"
	tmpl, err := template.New("test").Parse(t)
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	if err := tmpl.Execute(&doc, d); err != nil {
		panic(err)
	}
	s := doc.String()
	fmt.Printf("Result: %s\n", s)

	// Reading Json rom stdin.
	dec := json.NewDecoder(os.Stdin)
	var g Generator
	dec.Decode(&g)
	fmt.Printf("%+v\n", g)
}
