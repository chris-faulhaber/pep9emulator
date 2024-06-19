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
