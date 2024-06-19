package computer

type StatusBits struct {
	N bool // Negative
	Z bool // Zero
	V bool // oVerflow
	C bool // Carry
}

type Memory struct {
	Ram [65535]uint8
}

type Registers struct {
	A       uint16 // Accumulator
	X       uint16 // Index
	PC      uint16 // Program Counter
	SP      uint16 // Stack Pointer
	OpCode  uint8  // Loaded Operation Code
	Operand uint16 // Applicable Operand a parameter to the OpCode
}
