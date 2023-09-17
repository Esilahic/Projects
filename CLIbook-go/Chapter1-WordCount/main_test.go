package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three")
	exp := 3
	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("one\ntwo\nthree")
	exp := 3
	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d", exp, res)
	}
}
