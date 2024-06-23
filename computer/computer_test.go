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
	p.Operand = 0x7FFF
	p.A = 0x7FFF

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

	if p.A != expected {
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

func TestROR(t *testing.T) {
	expected := uint16(0x7FFF)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x10
	p.A = 0xFFFF
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestROL(t *testing.T) {
	expected := uint16(0xFFFE)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.OpCode = 0x0E
	p.A = 0xFFFF
	p.unaryArithmetic()

	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestIsCarry(t *testing.T) {
	testsValues := []struct {
		n1, n2   uint16
		expected bool
	}{
		// Existing positive test cases
		{0xFFFF, 0x0001, true},
		{0xFFFE, 0x0001, false},
		{0xFFFE, 0xFFFF, true},
		{0x0001, 0x0001, false},
		{0x0001, 0xFFFF, true},
		{0x0001, 0xFFFE, false},

		// Existing negative test cases
		{0x8000, 0x7000, false},
		{0x7000, 0x8000, false},
		{0x8000, 0x8000, true},
		{0xFFFF, 0xFFFF, true},
		{0xFFFF, 0x8000, true},
		{0x00FF, 0x0001, false},
		{0x8000, 0x7FFF, false},

		// Additional test cases
		// Minimum possible values - no carry
		{0x0000, 0x0000, false},

		// Manipulation middle bits - no carry
		{0x0F00, 0x00F0, false},

		// Manipulation middle bits - expecting carry
		{0x0F00, 0xF0F0, false},

		// Cross-sections,  no carry
		{0x1234, 0x4567, false},

		// Cross-sections, expecting carry
		{0xCDEF, 0xABCD, true},

		// Different order of input values - no carry
		{0x0001, 0xFFFE, false},
		{0xFFFE, 0x0001, false},

		// Different order of input values - expecting carry
		{0x0001, 0xFFFF, true},
		{0xFFFF, 0x0001, true},
	}

	for _, v := range testsValues {
		if v.expected != isCarry(v.n1, v.n2) {
			t.Errorf("left: %b, right: %b expected %t got %t", v.n1, v.n2, v.expected, isCarry(v.n1, v.n2))
			t.FailNow()
		}
	}
}

func TestIsOverflow(t *testing.T) {
	testsValues := []struct {
		n1, n2   uint16
		expected bool
	}{
		{0x7FFF, 0x7FFF, true},
		{0x7FFE, 0x0002, true},
		{0x7FFD, 0x0003, true},
		{0x0002, 0x7FFE, true},
		{0x0003, 0x7FFD, true},

		//Just at the edge of overflow
		{0x7FFF, 0x0001, true},
		{0x0001, 0x7FFF, true},

		//Just avoiding the overflow
		{0x7FFE, 0x0001, false},
		{0x0001, 0x7FFE, false},

		//Combination of negative values that cause overflow
		{0x8000, 0x8000, true},
		{0x8001, 0x8FFF, true},
		{0x8FFF, 0x8001, true},

		//Just at the edge of overflow
		{0x8000, 0x8000, true},

		//Just avoiding the overflow
		{0x8000, 0x7FFF, false},
		{0x7FFF, 0x8000, false},

		//Opposing signs never overflow
		{0x7FFF, 0x8000, false},
		{0x8000, 0x7FFF, false},
		{0x0000, 0xFFFF, false},
		{0xFFFF, 0x0000, false},

		//Opposing signs never overflow
		{0xFFFE, 0x0001, false},
	}

	for _, v := range testsValues {
		if v.expected != isOverflow(v.n1, v.n2) {
			t.Errorf("left: %b, right: %b expected %t got %t", v.n1, v.n2, v.expected, isOverflow(v.n1, v.n2))
			t.FailNow()
		}
	}
}

func TestIsNegative(t *testing.T) {
	testsValues := []struct {
		n        uint16
		expected bool
	}{
		{0x0000, false},
		{0x7FFF, false},
		{0x8000, true},
		{0xFFFF, true},
	}
	for _, v := range testsValues {
		if v.expected != isNegative(v.n) {
			t.Errorf("value: %b expected %t got %t", v.n, v.expected, isNegative(v.n))
			t.FailNow()
		}
	}
}

func TestMoveSPToA(t *testing.T) {
	expected := uint16(0xBEEF)
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}
	p.SP = expected
	p.OpCode = 0x03
	p.unaryArithmetic()
	if p.A != expected {
		t.Errorf("Expected %b got %b", expected, p.A)
		t.FailNow()
	}
}

func TestReturnFromCall(t *testing.T) {
	expected := uint16(0xBEEF)
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	// Set up stack with return addressC
	p.SP = 0x1000
	p.Ram[0x1000] = 0xBE
	p.Ram[0x1000+1] = 0xEF // Push 0xBEEF to the stack
	p.OpCode = 0x01        // RET instruction
	p.callAndReturn()

	if p.PC != expected {
		t.Errorf("Expected PC to be %b but got %b", expected, p.PC)
		t.FailNow()
	}
}

func TestCall(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	// Set up stack with return addressC
	p.SP = 0x1000
	p.PC = 0xBEEF
	p.Operand = 0xF00D
	p.OpCode = 0x24 // RET instruction
	p.callAndReturn()

	if p.PC != 0xF00D {
		t.Errorf("Expected PC to be %b but got %b", 0xF00D, p.PC)
		t.FailNow()
	}

	if p.ReadWord(p.SP) != 0xBEEF {
		t.Errorf("Expected PC to be %b but got %b", 0xBEEF, p.ReadWord(p.SP))
		t.FailNow()
	}
}

func TestMVSPA(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	// Set up stack with return addressC
	p.SP = 0x1000
	p.PC = 0xBEEF
	p.Operand = 0xF00D
	p.OpCode = 0x24 // RET instruction
	p.callAndReturn()

	if p.PC != 0xF00D {
		t.Errorf("Expected PC to be %b but got %b", 0xF00D, p.PC)
		t.FailNow()
	}

	if p.ReadWord(p.SP) != 0xBEEF {
		t.Errorf("Expected PC to be %b but got %b", 0xBEEF, p.ReadWord(p.SP))
		t.FailNow()
	}
}

func TestLoadFlags(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.N = true
	p.Z = true
	p.C = true
	p.V = true

	p.OpCode = 0x04
	p.unaryArithmetic()

	if p.A != 0x000F {
		t.Errorf("Expected %b but got %b", 0x000F, p.A)
		t.FailNow()
	}
}

func TestStoreFlags(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.A = 0x000F
	p.OpCode = 0x05
	p.unaryArithmetic()

	if !p.N {
		t.Errorf("Negitive failed")
		t.FailNow()
	}
	if !p.Z {
		t.Errorf("Z failed")
		t.FailNow()
	}
	if !p.V {
		t.Errorf("V failed")
		t.FailNow()
	}
	if !p.C {
		t.Errorf("C failed")
		t.FailNow()
	}

}

func TestAddSubToSP(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	// Set up stack with return addressC
	p.SP = 0x1000
	p.OpCode = 0x50 // Add to SP
	p.Operand = 0xAEEF
	p.nonUnaryArithmetic()

	if p.SP != 0xBEEF {
		t.Errorf("Expected PC to be %b but got %b", 0xBEEF, p.SP)
		t.FailNow()
	}
	// Test Subtract from SP
	p.OpCode = 0x58 // Subtract from SP
	p.Operand = 0x2000
	p.nonUnaryArithmetic()

	if p.SP != 0x9EEF {
		t.Errorf("Expected SP to be %b but got %b", 0x9EEF, p.SP)
		t.FailNow()
	}
}

func TestAddSubToA(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.A = 0x1000
	p.OpCode = 0x60 // Add to SP
	p.Operand = 0xAEEF
	p.nonUnaryArithmetic()

	if p.A != 0xBEEF {
		t.Errorf("Expected PC to be %b but got %b", 0xBEEF, p.A)
		t.FailNow()
	}
	// Test Subtract from SP
	p.OpCode = 0x70 // Subtract from SP
	p.Operand = 0x2000
	p.nonUnaryArithmetic()

	if p.A != 0x9EEF {
		t.Errorf("Expected SP to be %b but got %b", 0x9EEF, p.A)
		t.FailNow()
	}
}

func TestAndToA(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.A = 0xBEEF
	p.OpCode = 0x80 // Add to SP
	p.Operand = 0xFF00
	p.nonUnaryArithmetic()

	if p.A != 0xBE00 {
		t.Errorf("Expected PC to be %b but got %b", 0xBE00, p.A)
		t.FailNow()
	}
}

func TestOrToA(t *testing.T) {
	//Initialize a new Pep9Computer
	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.A = 0x00EF
	p.OpCode = 0x90 // Add to SP
	p.Operand = 0xBE00
	p.nonUnaryArithmetic()

	if p.A != 0xBEEF {
		t.Errorf("Expected PC to be %b but got %b", 0xBEEF, p.A)
		t.FailNow()
	}
}
