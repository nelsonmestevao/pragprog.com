package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("nao sei o que dizer\n")

	exp := 20

	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("nao sei o que dizer\n")

	exp := 5

	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("hello world\nwhat you do\nokay then")

	exp := 3

	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
