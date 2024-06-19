package computer

type StatusBits struct {
	N bool // Negative
	Z bool // Zero
	V bool // oVerflow
	C bool // Carry
}

type Registers struct {
	A       uint16 // Accumulator
	X       uint16 // Index
	PC      uint16 // Program Counter
	SP      uint16 // Stack Pointer
	OpCode  uint8  // Loaded Operation Code
	Operand uint16 // Applicable Operand a parameter to the OpCode
}

type Memory struct {
	Ram [65535]uint8
}

func (c *Memory) ReadByte(location uint16) uint16 {
	return uint16(c.Ram[location])
}

func (c *Memory) ReadWord(location uint16) uint16 {
	word := uint16(c.Ram[location]) << 8
	word |= uint16(c.Ram[location+1])
	return word
}

func (c *Memory) WriteByte(value uint8, location uint16) {
	c.Ram[location] = value
}

func (c *Memory) WriteWord(value uint16, location uint16) {
	c.Ram[location] = uint8(value)
	c.Ram[location+1] = uint8(value >> 8)
}

func (s StatusBits) UpdateStatusBits(negative, zero, carry, overflow *bool) {
	if negative != nil {
		s.N = *negative
	}
	if zero != nil {
		s.Z = *zero
	}
	if carry != nil {
		s.C = *carry
	}
	if overflow != nil {
		s.V = *overflow
	}
}
