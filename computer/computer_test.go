package computer

import "testing"

func TestBranchUnconditionally(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x12
	p.Operand = expected
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchLessEqual(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x14
	p.Operand = expected
	p.N = true
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchLess(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x16
	p.Operand = expected
	p.N = true
	p.Z = false
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchEqual(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x18
	p.Operand = expected
	p.Z = true
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchNotEqual(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x1A
	p.Operand = expected
	p.Z = false
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchGreaterEqual(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x1C
	p.Operand = expected
	p.Z = false
	p.N = false
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchGreater(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x1E
	p.Operand = expected
	p.Z = false
	p.N = false
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchOverflow(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x20
	p.Operand = expected
	p.V = true
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestBranchCarry(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x1A
	p.Operand = expected
	p.C = true
	p.PC = 0x0000

	p.branch()

	if p.PC != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadByteImmediate(t *testing.T) {
	expected := uint8(0xEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0xD0
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != uint16(expected) {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadByteDirect(t *testing.T) {
	expected := uint8(0xED)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = expected
	p.OpCode = 0xD1
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != uint16(expected) {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadByteIndirect(t *testing.T) {
	expected := uint8(0x88)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0xF0
	p.Ram[0xBEF0] = 0x0D
	p.Ram[0xF00D] = expected
	p.OpCode = 0xD2
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != uint16(expected) {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestStoreByteDirect(t *testing.T) {
	expected := uint8(0x0D)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0x00
	p.OpCode = 0xF1
	p.Operand = 0xBEEF
	p.A = 0xF00D

	p.store()

	if p.Ram[0xBEEF] != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestStoreByteIndirect(t *testing.T) {
	expected := uint8(0x88)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0xF0
	p.Ram[0xBEF0] = 0x0D
	p.Ram[0xF00D] = 0x00
	p.OpCode = 0xF2
	p.Operand = 0xBEEF
	p.A = 0x7788

	p.store()

	if p.Ram[0xF00D] != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadWordDirect(t *testing.T) {
	expected := uint16(0xFEED)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0xFE
	p.Ram[0xBEF0] = 0xED
	p.OpCode = 0xC1
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != 0xFEED {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadWordImmediate(t *testing.T) {
	expected := uint16(0xBEEF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0xC0
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != uint16(expected) {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestLoadWordIndirect(t *testing.T) {
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0xF0
	p.Ram[0xBEF0] = 0x0D
	p.Ram[0xF00D] = 0x12
	p.Ram[0xF00E] = 0x34
	p.OpCode = 0xC2
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != 0x1234 {
		t.Errorf("Expected %b got %b", 0x1234, p.A)
		t.FailNow()
	}
}

func TestStoreWordDirect(t *testing.T) {
	expected1 := uint8(0xF0)
	expected2 := uint8(0x0D)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0x00
	p.Ram[0xBEF0] = 0x00
	p.OpCode = 0xC1
	p.Operand = 0xBEEF
	p.A = 0xF00D

	p.store()

	if p.Ram[0xBEEF] != expected1 {
		t.Errorf("Expected %b got %b", expected1, p.Ram[0xBEEF])
		t.FailNow()
	}

	if p.Ram[0xBEF0] != expected2 {
		t.Errorf("Expected %b got %b", expected2, p.Ram[0xBEF0])
		t.FailNow()
	}

}

func TestStoreWordIndirect(t *testing.T) {
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = 0xF0
	p.Ram[0xBEF0] = 0x0D
	p.Ram[0xF00D] = 0x00
	p.Ram[0xF00E] = 0x00
	p.OpCode = 0xC2
	p.Operand = 0xBEEF
	p.A = 0x7788

	p.store()

	if p.Ram[0xF00D] != 0x77 {
		t.Errorf("Expected %b got %b", 0x77, p.Ram[0xF00D])
		t.FailNow()
	}

	if p.Ram[0xF00E] != 0x88 {
		t.Errorf("Expected %b got %b", 0x88, p.Ram[0xF00E])
		t.FailNow()
	}
}
