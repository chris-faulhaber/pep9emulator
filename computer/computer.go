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

func (c *Pep9Computer) fetch() {
	c.OpCode = uint8(c.ReadByte(c.PC))
	c.PC += 1

	if c.OpCode >= 0x12 { // if OpCode requires an Operand, fetch it
		c.Operand = c.ReadWord(c.PC)
		c.PC += 2
	}
}

func (c *Pep9Computer) execute() {
	switch c.OpCode {
	case 0x00: // HALT
		break
	case 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1c, 0x1D, 0x1E, 0x1F, 0x20, 0x21, 0x22, 0x23:
		c.branch()
	case 0xC0, 0xC1, 0xC2, 0xC3, 0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xCB, 0xCC, 0xCD, 0xCE, 0xCF,
		0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF:
		c.load()
		break
	case 0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
		0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA, 0xFB, 0xFC, 0xFD, 0xFE, 0xFF:
		c.store()
		break
	case 0x06, 0x07, 0x08, 0x09, 0x10, 0x11,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F,
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F,
		0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F,
		0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F,
		0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F:
		c.arithmetic()
		c.arithmetic()
		break
	default:
		log.Fatal("Unknown opcode")
	}
}

func (c *Pep9Computer) load() {
	var result uint16
	var loadFunc func(location uint16) uint16
	var isWord bool

	switch c.OpCode & 0x10 {
	case 0x00: // Word
		loadFunc = c.ReadWord
		isWord = true
	case 0x10: // Byte
		loadFunc = c.ReadByte
		isWord = false
	}

	switch c.OpCode & 0x07 {
	case 0: // Immediate
		if isWord {
			result = c.Operand
		} else {
			result = uint16(uint8(c.Operand))
		}

		break
	case 1: // Direct
		result = loadFunc(c.Operand)
		break
	case 2: // Indirect
		location := c.ReadWord(c.Operand)
		result = loadFunc(location)
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

	c.N = isNegative(result)
	c.Z = result == 0 // Set 'Z' if the result is zero.
}

func (c *Pep9Computer) store() {
	var value uint16
	var writeFunc func(value uint16, location uint16)

	switch c.OpCode & 0x10 {
	case 0x00: // Word
		writeFunc = c.WriteWord
		break
	case 0x10: // Byte
		writeFunc = c.WriteByte
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
		break
	case 1: // Direct
		writeFunc(value, c.Operand)
		break
	case 2: // Indirect
		location := c.ReadWord(c.Operand)
		writeFunc(value, location)
		break
	default:
		log.Fatal("Not yet implemented")
	}
}

func (c *Pep9Computer) branch() {
	toBranch := false

	switch c.OpCode {
	case 0x12, 0x13: // unconditional
		toBranch = true
		break
	case 0x14, 0x15: // <=
		toBranch = c.N || c.Z
		break
	case 0x16, 0x17: // <
		toBranch = c.N
		break
	case 0x18, 0x19: // ==
		toBranch = c.Z
		break
	case 0x1A, 0x1B: // !=
		toBranch = !c.Z
		break
	case 0x1C, 0x1D: // >=
		toBranch = !c.N
		break
	case 0x1E, 0x1F: // >
		toBranch = !c.N && !c.Z
		break
	case 0x20, 0x21: // V
		toBranch = c.V
		break
	case 0x22, 0x23: // C
		toBranch = c.C
		break
	default:
		log.Fatal("Not yet implemented")
	}

	var location uint16

	switch c.OpCode & 0x1 {
	case 0: // indirect
		location = c.Operand
		break
	case 0x1:
		//TODO branch by index
		log.Fatal("Not yet implemented")
	default:
		log.Fatal("Not yet implemented")
	}

	if toBranch {
		c.PC = location
	}
}

func (c *Pep9Computer) compare() {
	result := c.Operand - c.A
	c.N = isNegative(result)
	c.Z = result == 0
	c.V = isOverflow(c.Operand, c.A)
	c.C = isCarry(c.Operand, c.A)
}

func (c *Pep9Computer) arithmetic() {
	var value *uint16

	if c.OpCode&0x1 == 0 {
		value = &c.A
	} else {
		value = &c.X
	}

	switch c.OpCode {
	case 0x06, 0x07: //Bitwise invert
		*value = *value ^ 0xFFFF
		break
	case 0x08, 0x09: // Negate
		break
	case 0x0A, 0x0B: // Arithmetic shift Left
		*value = *value << 1
		break
	case 0x0C, 0x0D: // Arithmetic shift Right
		*value = *value >> 1
		break
	case 0x0E, 0x0F: // Rotate Left
		break
	case 0x10, 0x11: // Rotate Right
		break

	default:
		log.Fatal("Not yet implemented")
	}

}

func isNegative(a uint16) bool {
	return a&0x8000 != 0 // Set 'N' if the result is negative (the leftmost bit is set).
}

func isCarry(a, b uint16) bool {
	return a > a-b
}

func isOverflow(a, b uint16) bool {
	subResult := a - b
	return isNegative(a) && isNegative(b) && !isNegative(subResult) ||
		!isNegative(a) && !isNegative(b) && isNegative(subResult)
}
