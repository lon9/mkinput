package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNewGeneratorFromStdioSuccess(t *testing.T) {
	count := 1
	err := filepath.Walk("tests/success/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			t.Logf("Test case: %d", count)
			f, err := ioutil.ReadFile(path)
			if err != nil {
				t.Error(err)
			}

			stdin := bytes.NewBufferString(string(f))
			_, err = NewGeneratorFromStdin(stdin)
			if err != nil {
				t.Error(err)
			}
			count++
		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

}

func TestNewGeneratorFromStdioFailed(t *testing.T) {
	count := 1
	err := filepath.Walk("tests/failed/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			t.Logf("Test case: %d", count)
			f, err := ioutil.ReadFile(path)
			if err != nil {
				t.Error(err)
			}

			stdin := bytes.NewBufferString(string(f))
			_, err = NewGeneratorFromStdin(stdin)
			if err == nil {
				t.Fatal("Error was expected, but not occured.")
			}
			count++
		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerate(t *testing.T) {
	count := 1
	err := filepath.Walk("tests/success/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".json" {
			t.Logf("Test case: %d", count)
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
			count++
		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
