package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("nao sei o que dizer\n")

	exp := 5

	res := count(b)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
