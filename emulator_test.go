package main

import (
	"testing"
)

func TestHlt(t *testing.T) {
	err := Emulate("tests/test1.bin")
	if err != nil {
		t.Error(err)
	}
}

func TestJumpAbs(t *testing.T) {
	err := Emulate("tests/test2.bin")
	if err != nil {
		t.Error(err)
	}
}
