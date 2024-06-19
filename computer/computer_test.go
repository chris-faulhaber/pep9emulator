package computer

import "testing"

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
	expected := uint8(0xED)

	p := Pep9Computer{
		Processor: Processor{},
		Memory:    Memory{},
	}

	p.Ram[0xBEEF] = expected
	p.OpCode = 0xC1
	p.Operand = 0xBEEF
	p.A = 0x0000

	p.load()

	if p.A != uint16(expected) {
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
	p.OpCode = 0xC1
	p.Operand = 0xBEEF
	p.A = 0xF00D

	p.store()

	if p.Ram[0xBEEF] != expected1 {
		t.Errorf("Expected %b got %b", expected1, p.Ram[0xBEF0])
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

	if p.Ram[0xF00D] != 0x12 {
		t.Errorf("Expected 0x00010010 got %b", p.Ram[0xF00D])
		t.FailNow()
	}

	if p.Ram[0xF00E] != 0x34 {
		t.Errorf("Expected 00111000 got %b", p.Ram[0xF00E])
		t.FailNow()
	}
}
