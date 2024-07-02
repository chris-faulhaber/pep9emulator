package computer

import "log"

type Pep9Computer struct {
	Processor
	Memory
	HALT bool
}

func (c *Pep9Computer) Initialize() {
	// Default vectors for running a program starting at 0x0000
	c.PC = 0x0000
	c.A = 0x0000
	c.X = 0x0000
	c.SP = 0xFB8F
	c.HALT = false
}

func (c *Pep9Computer) LoadProgram(program []byte) {
	for loc, value := range program {
		c.Ram[loc] = value
	}
}

func (c *Pep9Computer) ExecuteVonNeumann() {
	for !c.HALT && c.OpCode != 0x00 || c.PC == 0 {
		c.fetch()
		c.execute()
	}
}

func (c *Pep9Computer) fetch() {
	c.OpCode = uint8(c.LoadByte(c.PC))
	c.PC += 1

	if c.OpCode >= 0x12 { // if OpCode requires an Operand, fetch it
		c.Operand = c.LoadWord(c.PC)
		c.PC += 2
	}
}

func (c *Pep9Computer) execute() {
	switch c.OpCode {
	case 0x00: // HALT
		break
	case 0x01, 0x02, 0x24, 0x25:
		c.callAndReturn()
		break
	case 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11:
		c.unaryArithmetic()
		break
	case 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1c, 0x1D, 0x1E, 0x1F,
		0x20, 0x21, 0x22, 0x23:
		c.branch()
	case 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F,
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F,
		0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F,
		0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F,
		0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F:
		c.nonUnaryArithmetic()
		break
	case 0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF,
		0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8, 0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF:
		c.compare()
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
		log.Printf("Unknown opcode")
		c.HALT = true
	}
}

func (c *Pep9Computer) load() {
	result := c.loadWithMode()
	destination := c.getRegisterBit3()
	*destination = result

	c.N = isNegative(result)
	c.Z = result == 0 // Set 'Z' if the result is zero.
}

func (c *Pep9Computer) store() {
	source := c.getRegisterBit3()
	c.storeWithMode(source)
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
		log.Printf("Branch opcode %s not implemented", string(c.OpCode))
		c.HALT = true
	}

	var location uint16

	if c.OpCode&0x1 == 0 { // immediate
		location = c.Operand
	} else {
		location = c.Operand + c.X
	}

	if toBranch {
		c.PC = location
	}
}

func (c *Pep9Computer) compare() {
	var left, right uint16

	right = *c.getRegisterBit3()
	left = c.loadWithMode()

	result := left - right
	c.N = isNegative(result)
	c.Z = result == 0
	c.V = isOverflow(left, right)
	c.C = isCarry(left, right)
}

func (c *Pep9Computer) callAndReturn() {
	var location uint16

	if c.OpCode == 0x01 { // Return
		c.PC = c.LoadWord(c.SP)
		c.SP += 2
	} else { // Call
		if c.OpCode&0x1 == 0 { // immediate
			location = c.Operand
		} else {
			location = c.Operand + c.X
		}
		c.SP -= 2
		c.StoreWord(c.PC, c.SP)
		c.PC = location
	}
}

func (c *Pep9Computer) unaryArithmetic() {
	var value *uint16
	value = c.getRegisterBit0()

	switch c.OpCode {
	case 0x03:
		c.A = c.SP
	case 0x04: // NZVC Flags to A<12..15> 15 is LSB
		c.A = 0
		if c.C {
			c.A |= 1
		}
		if c.V {
			c.A |= 1 << 1
		}
		if c.Z {
			c.A |= 1 << 2
		}
		if c.N {
			c.A |= 1 << 3
		}
	case 0x05:
		c.C = c.A&0x01 == 0x01
		c.Z = c.A&0x02 == 0x02
		c.V = c.A&0x03 == 0x03
		c.N = c.A&0x04 == 0x04
	case 0x06, 0x07: //Bitwise invert
		*value = ^*value
		c.N = isNegative(*value)
		c.Z = *value == 0
		break
	case 0x08, 0x09: // Negate in 2's complement
		prev := *value
		*value = ^*value + 1
		c.N = isNegative(*value)
		c.Z = *value == 0
		c.V = isNegative(prev) && !isNegative(*value) ||
			!isNegative(prev) && isNegative(*value)
		break
	case 0x0A, 0x0B: // Arithmetic shift Left
		prev := *value
		*value = *value << 1
		c.N = isNegative(*value)
		c.Z = *value == 0
		c.V = isNegative(prev) && !isNegative(*value) ||
			!isNegative(prev) && isNegative(*value)
		c.C = prev&0x8000 != 0 // The most significant bit is put into the carry flag.
		break
	case 0x0C, 0x0D: // Arithmetic shift Right
		prev := *value
		var msb uint16
		if isNegative(*value) {
			msb = 0x8000
		}
		*value = (*value >> 1) | msb
		c.N = isNegative(*value)
		c.Z = *value == 0
		c.V = isNegative(prev) && !isNegative(*value) ||
			!isNegative(prev) && isNegative(*value)
		c.C = prev&0x1 != 0 // The least significant bit is put into the carry flag.
		break
	case 0x0E, 0x0F: // Rotate Left with Carry (RLC)
		prev := *value
		*value = *value << 1

		c.C = prev&0x8000 != 0 // The most significant bit is put into the carry flag.
		if c.C {
			*value = *value | 0x0001
		}

		c.N = isNegative(*value)
		c.Z = *value == 0
		c.V = isNegative(prev) && !isNegative(*value) ||
			!isNegative(prev) && isNegative(*value)
		break
	case 0x10, 0x11: // Rotate Right with Carry (RRC)
		prev := *value
		*value = *value >> 1

		c.C = prev&0x1 != 0 // The least significant bit is put into the carry flag.
		if c.C {
			*value = *value | 0x8000 // If the original carry was set, it is put into bit 15.
		}
		c.N = isNegative(*value)
		c.Z = *value == 0
		c.V = isNegative(prev) && !isNegative(*value) ||
			!isNegative(prev) && isNegative(*value)
		break
	default:
		log.Printf("Opcode %s yet implemented", string(c.OpCode))
		c.HALT = true
	}

}

func (c *Pep9Computer) nonUnaryArithmetic() {
	value := c.loadWithMode()
	prev := value
	dest := c.getRegisterBit3()

	switch c.OpCode {
	case 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57:
		c.SP += value
		c.V = isOverflow(prev, value)
		c.C = isCarry(prev, value) // The least significant bit is put into the carry flag.
		break
	case 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F:
		c.SP -= value
		c.V = isOverflow(prev, value)
		c.C = isCarry(prev, value) // The least significant bit is put into the carry flag.
		break
	case 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F:
		*dest += value
		c.V = isOverflow(prev, value)
		c.C = isCarry(prev, value) // The least significant bit is put into the carry flag.
		break
	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F:
		*dest -= value
		c.V = isOverflow(prev, value)
		c.C = isCarry(prev, value) // The least significant bit is put into the carry flag.
		break
	case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C, 0x8D, 0x8E, 0x8F:
		*dest &= value
		break
	case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F:
		*dest |= value
		break
	default:
		log.Printf("Opcdoe %s yet implemented", string(c.OpCode))
		c.HALT = true
	}

	c.N = isNegative(value)
	c.Z = value == 0
}

func (c *Pep9Computer) loadWithMode() uint16 {
	var result uint16
	var loadFunc func(location uint16) uint16
	var isWord bool

	if c.OpCode&0x10 == 0x10 && c.OpCode >= 0xA0 { // Byte if we have too
		loadFunc = c.LoadByte
	} else { // Word is the default
		loadFunc = c.LoadWord
		isWord = true
	}

	switch c.OpCode & 0x07 {
	case 0: // Immediate
		if isWord {
			result = c.Operand
		} else {
			result = uint16(uint8(c.Operand))
		}
		break
	case 1: // Direct Mem[op]
		result = loadFunc(c.Operand)
		break
	case 2: // Indirect Mem[Mem[op]]
		location := c.LoadWord(c.Operand)
		result = loadFunc(location)
		break
	case 3: // Stack relative Mem[SP+op]
		result = loadFunc(c.SP + c.Operand)
		break
	case 4: // Stack-relative deferred Mem[Mem[SP + Op]]
		location := c.LoadWord(c.SP + c.Operand)
		result = loadFunc(location)
		break
	case 5: // Indexed Mem[Op + X]
		result = loadFunc(c.Operand + c.X)
		break
	case 6: // Stack Indexed Mem[SP + Op + X]
		result = loadFunc(c.SP + c.Operand + c.X)
		break
	case 7: // Stack-deferred Indexed Mem[Mem[SP + Op + X]]
		location := loadFunc(c.SP + c.Operand + c.X)
		result = loadFunc(location)
	}

	c.N = isNegative(result)
	c.Z = result == 0
	return result
}

func (c *Pep9Computer) getRegisterBit3() *uint16 {
	return c.getRegister(0x08)
}

func (c *Pep9Computer) getRegisterBit0() *uint16 {
	return c.getRegister(0x01)
}

func (c *Pep9Computer) getRegister(mask uint8) *uint16 {
	if c.OpCode&mask == 0 {
		return &c.A
	} else {
		return &c.X
	}
}
func (c *Pep9Computer) storeWithMode(value *uint16) {
	var writeFunc func(value uint16, location uint16)

	if c.OpCode&0x10 == 0 { // Word
		writeFunc = c.StoreWord
	} else { // Byte
		writeFunc = c.StoreByte
	}

	switch c.OpCode & 0x07 {
	case 0:
		// Can't store to immediate value
		log.Printf("Opcdoe %s is invalid :facepalm: Can't store to an immediate value", string(c.OpCode))
		c.HALT = true
	case 1: // Direct
		writeFunc(*value, c.Operand)
		break
	case 2: // Indirect
		location := c.LoadWord(c.Operand)
		writeFunc(*value, location)
		break
	case 3: // Stack relative Mem[SP+op]
		writeFunc(*value, c.SP+c.Operand)
		break
	case 4: // Stack-relative deferred Mem[Mem[SP + Op]]
		location := c.LoadWord(c.SP + c.Operand)
		writeFunc(*value, location)
		break
	case 5: // Indexed Mem[Op + X]
		writeFunc(*value, c.Operand+c.X)
		break
	case 6: // Stack Indexed Mem[SP + Op + X]
		writeFunc(*value, c.SP+c.Operand+c.X)
		break
	case 7: // Stack-deferred Indexed Mem[Mem[SP + Op + X]]
		location := c.LoadWord(c.SP + c.Operand + c.X)
		writeFunc(*value, location)
	}
}

func isNegative(a uint16) bool {
	return a&0x8000 != 0 // if the result is negative (the leftmost bit is set).
}

func isCarry(a, b uint16) bool {
	return uint32(a)+uint32(b) > 0xFFFF
}

func isOverflow(a, b uint16) bool {
	result := a + b
	return isNegative(a) && isNegative(b) && !isNegative(result) ||
		!isNegative(a) && !isNegative(b) && isNegative(result)
}
