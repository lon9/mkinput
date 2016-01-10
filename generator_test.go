package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	count := 1
	err := filepath.Walk("src/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, err := ioutil.ReadFile(path)
			if err != nil {
				t.Error(err)
			}

			var g Generator
			if err = json.Unmarshal(f, &g); err != nil {
				t.Error(err)

			}
			g.Generate()
			count++

		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

}
