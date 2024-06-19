package main

import (
	"pep9emulator/computer"
	"testing"
)

func TestInitialize(t *testing.T) {
	expected := 0x0000
	p := computer.Pep9Computer{}
	p.PC = 0xFFFF
	p.Initialize()

	if p.PC != 0x0000 {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadProgram(t *testing.T) {
	p := computer.Pep9Computer{}
	expected := []byte{0x12, 0x34, 0x56}
	p.LoadProgram(expected)

	for i, val := range expected {
		if val != p.Ram[i] {
			t.Errorf("Expected %b got %b", val, p.Ram[i])
			t.FailNow()
		}
	}
}

func TestExecuteVonNeumann(t *testing.T) {

}
