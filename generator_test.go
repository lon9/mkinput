package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNewGeneratorFromStdioSuccess(t *testing.T) {
	err := filepath.Walk("src/success/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, err := ioutil.ReadFile(path)
			if err != nil {
				t.Error(err)
			}

			stdin := bytes.NewBufferString(string(f))
			_, err = NewGeneratorFromStdin(stdin)
			if err != nil {
				t.Error(err)
			}
		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

}

func TestNewGeneratorFromStdioFailed(t *testing.T) {
	err := filepath.Walk("src/failed/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, err := ioutil.ReadFile(path)
			if err != nil {
				t.Error(err)
			}

			stdin := bytes.NewBufferString(string(f))
			_, err = NewGeneratorFromStdin(stdin)
			if err == nil {
				t.Fatal("Error was expected, but not occured.")
			}

		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerate(t *testing.T) {
	err := filepath.Walk("src/success/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, err := ioutil.ReadFile(path)
			if err != nil {
				t.Error(err)
			}

			stdin := bytes.NewBufferString(string(f))
			g, err := NewGeneratorFromStdin(stdin)
			if err != nil {
				t.Error(err)
			}
			g.Generate()
		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
