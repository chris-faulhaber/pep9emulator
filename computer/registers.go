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
	Ram                                 [65535]uint8
	StandardInput, StandardOutput       [256]uint8
	StandardInputLoc, StandardOutputLoc int
}

func (c *Memory) LoadByte(location uint16) uint16 {
	if location == 0xFC15 { //Standard input (byte stream)
		value := uint16(c.StandardInput[c.StandardInputLoc])
		c.StandardInputLoc = +1
		if c.StandardInputLoc > 255 {
			c.StandardInputLoc = 0
		}
		return value
	}
	return uint16(c.Ram[location])
}

func (c *Memory) LoadWord(location uint16) uint16 {
	word := c.LoadByte(location) << 8
	word |= c.LoadByte(location + 1)
	return word
}

func (c *Memory) StoreByte(value uint16, location uint16) {
	if location == 0xFC16 { //Standard input (byte stream)
		c.StandardOutput[c.StandardOutputLoc] = uint8(value)
		c.StandardOutputLoc = +1
		if c.StandardOutputLoc > 255 {
			c.StandardOutputLoc = 0
		}
	} else {
		c.Ram[location] = uint8(value)
	}
}

func (c *Memory) StoreWord(value uint16, location uint16) {
	c.StoreByte(uint16(uint8(value>>8)), location)
	c.StoreByte(uint16(uint8(value)), location+1)
}
