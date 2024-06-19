package computer

import "log"

type Pep9Computer struct {
	Processor
	Memory
}

func (c *Pep9Computer) Initialize() {
	// Default vectors for running a program starting at 0x0000
	c.PC = 0x0000
	c.A = 0x0000
	c.X = 0x0000
	c.SP = 0xFB8F
}

func (c *Pep9Computer) LoadProgram(program []byte) {
	for loc, value := range program {
		c.Ram[loc] = value
	}
}

func (c *Pep9Computer) ExecuteVonNeumann() {
	var HALT = uint8(0x00)
	for c.OpCode != HALT || c.PC == 0 {
		c.fetch()
		c.execute()
	}
}

func (c *Pep9Computer) load() {
	var result uint16
	var loadFunc func(location uint16) uint16

	switch c.OpCode & 0x10 {
	case 0x00: // Word
		loadFunc = c.ReadWord
	case 0x10: // Byte
		loadFunc = c.ReadByte
	}

	switch c.OpCode & 0x07 {
	case 0: // Immediate
		result = uint16(uint8(c.Operand))
		break
	case 1: // Direct
		result = loadFunc(c.Operand)
		break
	default:
		log.Fatal("Not yet implemented")
	}

	switch c.OpCode & 0x08 {
	case 0:
		c.A = result
		break
	case 0x08:
		c.X = result
	}
}

func (c *Pep9Computer) store() {
	var value uint16
	var writeFunc func(value uint16, location uint16)

	switch c.OpCode & 0x10 {
	case 0x00: // Word
		writeFunc = c.WriteByte
	case 0x10: // Byte
		writeFunc = c.WriteWord
	}

	switch c.OpCode & 0x08 {
	case 0:
		value = c.A
		break
	case 0x08:
		value = c.X
	}

	switch c.OpCode & 0x07 {
	case 0:
		// Can't store to immediate value
		log.Fatal("Can't store to immediate value")
	case 1: // Direct
		writeFunc(value, c.Operand)
		break
	default:
		log.Fatal("Not yet implemented")
	}
}

func (c *Pep9Computer) fetch() {
	c.OpCode = uint8(c.ReadByte(c.PC))
	c.PC += 1

	if c.OpCode >= 0x13 { // if OpCode requires an Operand, fetch it
		c.Operand = c.ReadWord(c.PC)
		c.PC += 2
	}
}

func (c *Pep9Computer) execute() {
	switch c.OpCode {
	case 0x00: // HALT
		break
	case 0xC0, 0xC1, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xCB, 0xCC, 0xCD, 0xCE, 0xCF,
		0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF:
		c.load()
		break
	case 0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF:
		c.store()
		break
	default:
		log.Fatal("Unknown opcode")
	}
}
