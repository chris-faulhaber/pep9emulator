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

func (c *Pep9Computer) loadByte() {
	var result uint16

	switch c.OpCode & 0x07 {
	case 0: // Immediate
		result = uint16(uint8(c.Operand))
		break
	case 1: // Direct
		result = uint16(c.ReadByte(c.Operand))
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

func (c *Pep9Computer) storeByte() {
	var value uint8

	switch c.OpCode & 0x08 {
	case 0:
		value = uint8(c.A)
		break
	case 0x08:
		value = uint8(c.X)
	}

	switch c.OpCode & 0x07 {
	case 0:
		// Can't store to immediate value
		log.Fatal("Can't store to immediate value")
	case 1: // Direct
		c.WriteByte(value, c.Operand)
		break
	default:
		log.Fatal("Not yet implemented")
	}
}

func (c *Pep9Computer) fetch() {
	c.OpCode = c.ReadByte(c.PC)
	c.PC += 1

	if c.OpCode >= 0x13 { // if OpCode requires an Operand, fetch it
		c.Operand = c.ReadWord(uint16(c.PC))
		c.PC += 2
	}
}

func (c *Pep9Computer) execute() {
	switch c.OpCode {
	case 0x00: // HALT
		break
	case 0xD0, 0xD1, 0xD2:
		c.loadByte()
		break
	case 0xF0, 0xF1, 0xF2:
		c.storeByte()
		break
	// Implement more cases here for different opcodes as needed
	default:
		log.Fatal("Unknown opcode")
	}
}
