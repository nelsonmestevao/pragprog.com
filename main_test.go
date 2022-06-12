package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("nao sei o que dizer\n")

	exp := 5

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("hello world\nwhat you do\nokay then")

	exp := 3

	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
