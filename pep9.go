package main

import "pep9emulator/computer"

func main() {
	p := computer.Pep9Computer{}
	program := []byte{0x00}

	p.Initialize()
	p.LoadProgram(program)
	p.ExecuteVonNeumann()
}
