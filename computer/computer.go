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

}

func (c *Pep9Computer) loadByte() {
	var result uint16

	switch c.OpCode & 0x07 {
	case 0: // Immediate
		result = uint16(uint8(c.Operand))
		break
	case 1: // Direct
		result = c.ReadByte(c.Operand)
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
