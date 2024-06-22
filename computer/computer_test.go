package computer

import "testing"

func TestInitialize(t *testing.T) {
	expected := 0x0000
	p := Pep9Computer{}
	p.PC = 0xFFFF
	p.Initialize()

	if p.PC != 0x0000 {
		t.Errorf("Expected %b got %b", expected, p.PC)
		t.FailNow()
	}
}

func TestLoadProgram(t *testing.T) {
	p := Pep9Computer{}
	expected := []byte{0x12, 0x34, 0x56}
	p.LoadProgram(expected)

	for i, val := range expected {
		if val != p.Ram[i] {
			t.Errorf("Expected %b got %b", val, p.Ram[i])
			t.FailNow()
		}
	}
}

func TestExecuteVonNeumann(t *testing.T) {
	expected := uint8(0x42)
	p := Pep9Computer{}
	p.Initialize()

	//A simple program that loads 0x48 into the A register
	p.LoadProgram([]byte{0xD1, 0x00, 0x04, 0x00, expected})
	p.ExecuteVonNeumann()

	if p.A != uint16(expected) {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestCompare(t *testing.T) {
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0xA0
	p.Operand = 0x1234
	p.A = 0x1234

	p.compare()

	if !p.Z || p.N {
		t.Errorf("Expected true,false got %t, %t", p.Z, p.N)
		t.FailNow()
	}
}

func TestCompareNegative(t *testing.T) {
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0xA0
	p.Operand = 0x0001
	p.A = 0x1234

	p.compare()

	if !p.N {
		t.Errorf("Expected true got %t", p.N)
		t.FailNow()
	}
}

func TestCompareOverflow(t *testing.T) {
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0xA0
	p.Operand = 0x1
	p.A = 0xF

	p.compare()

	if !p.V {
		t.Errorf("Expected true got %t", p.V)
		t.FailNow()
	}
}

func TestCompareCarry(t *testing.T) {
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0xA0
	p.Operand = 0xFFFF
	p.A = 0xFFFF

	p.compare()

	if !p.C {
		t.Errorf("Expected true got %t", p.C)
		t.FailNow()
	}
}

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

func TestBranchLessEqualNegative(t *testing.T) {
	expected := uint16(0x0000)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x14
	p.Operand = 0xBEEF
	p.N = false
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

	p.OpCode = 0x22
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

func TestBitwiseInvert(t *testing.T) {
	expected := uint16(0xF0F0)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x06
	p.A = 0x0F0F
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestNegate(t *testing.T) {
	expected := uint16(0xFFFF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x08
	p.A = 0x1
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestASL(t *testing.T) {
	expected := uint16(0xAAAA)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x0A
	p.A = 0x5555
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestASRNegative(t *testing.T) {
	expected := uint16(0xFFFE)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x0C
	p.A = 0xFFFC
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestASRPositive(t *testing.T) {
	expected := uint16(0x0004)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x0C
	p.A = 0x0008
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}
