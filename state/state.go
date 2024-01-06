package state

import (
	"fmt"

	rarch_encoding "github.com/jmbert/rarch-encoding"
)

const rom_end = 0xF000

type Register struct {
	Value uint64
}

type Registers struct {
	PC Register

	Status Register

	R0  Register
	R1  Register
	R2  Register
	R3  Register
	R4  Register
	R5  Register
	R6  Register
	R7  Register
	R8  Register
	R9  Register
	R10 Register
	R11 Register
	R12 Register
	R13 Register

	SP Register
	BP Register

	I0  Register
	I1  Register
	I2  Register
	I3  Register
	I4  Register
	I5  Register
	I6  Register
	I7  Register
	I8  Register
	I9  Register
	I10 Register
	I11 Register
	I12 Register
	I13 Register
	I14 Register
	I15 Register

	S0 Register
	S1 Register
	S2 Register
	S3 Register
	S4 Register
	S5 Register
	S6 Register
	S7 Register

	D0 Register
	D1 Register
	D2 Register
	D3 Register
	D4 Register
	D5 Register
	D6 Register
	D7 Register

	C0  Register
	C1  Register
	C2  Register
	C3  Register
	C4  Register
	C5  Register
	C6  Register
	C7  Register
	C8  Register
	C9  Register
	C10 Register
	C11 Register
	C12 Register
	C13 Register
	C14 Register
	C15 Register
}

func (regs *Registers) WriteReg(r rarch_encoding.Register, value uint64) {
	switch r.Register_type {
	case rarch_encoding.RegisterGP:
		switch r.Register {
		case rarch_encoding.REG_R0:

		case rarch_encoding.REG_R1:
			regs.R1.Value = value
		case rarch_encoding.REG_R2:
			regs.R2.Value = value
		case rarch_encoding.REG_R3:
			regs.R3.Value = value
		case rarch_encoding.REG_R4:
			regs.R4.Value = value
		case rarch_encoding.REG_R5:
			regs.R5.Value = value
		case rarch_encoding.REG_R6:
			regs.R6.Value = value
		case rarch_encoding.REG_R7:
			regs.R7.Value = value
		case rarch_encoding.REG_R8:
			regs.R8.Value = value
		case rarch_encoding.REG_R9:
			regs.R9.Value = value
		case rarch_encoding.REG_R10:
			regs.R10.Value = value
		case rarch_encoding.REG_R11:
			regs.R11.Value = value
		case rarch_encoding.REG_R12:
			regs.R12.Value = value
		case rarch_encoding.REG_R13:
			regs.R13.Value = value
		case rarch_encoding.REG_SP:
			regs.SP.Value = value
		case rarch_encoding.REG_BP:
			regs.BP.Value = value
		}
	case rarch_encoding.RegisterIndex:
		switch r.Register {
		case rarch_encoding.REG_I0:
			regs.I0.Value = value
		case rarch_encoding.REG_I1:
			regs.I1.Value = value
		case rarch_encoding.REG_I2:
			regs.I2.Value = value
		case rarch_encoding.REG_I3:
			regs.I3.Value = value
		case rarch_encoding.REG_I4:
			regs.I4.Value = value
		case rarch_encoding.REG_I5:
			regs.I5.Value = value
		case rarch_encoding.REG_I6:
			regs.I6.Value = value
		case rarch_encoding.REG_I7:
			regs.I7.Value = value
		case rarch_encoding.REG_I8:
			regs.I8.Value = value
		case rarch_encoding.REG_I9:
			regs.I9.Value = value
		case rarch_encoding.REG_I10:
			regs.I10.Value = value
		case rarch_encoding.REG_I11:
			regs.I11.Value = value
		case rarch_encoding.REG_I12:
			regs.I12.Value = value
		case rarch_encoding.REG_I13:
			regs.I13.Value = value
		case rarch_encoding.REG_I14:
			regs.I14.Value = value
		case rarch_encoding.REG_I15:
			regs.I15.Value = value
		}
	case rarch_encoding.RegisterMMove:
		switch r.Register {
		case rarch_encoding.REG_S0:
			regs.S0.Value = value
		case rarch_encoding.REG_S1:
			regs.S1.Value = value
		case rarch_encoding.REG_S2:
			regs.S2.Value = value
		case rarch_encoding.REG_S3:
			regs.S3.Value = value
		case rarch_encoding.REG_S4:
			regs.S4.Value = value
		case rarch_encoding.REG_S5:
			regs.S5.Value = value
		case rarch_encoding.REG_S6:
			regs.S6.Value = value
		case rarch_encoding.REG_S7:
			regs.S7.Value = value
		case rarch_encoding.REG_D0:
			regs.D0.Value = value
		case rarch_encoding.REG_D1:
			regs.D1.Value = value
		case rarch_encoding.REG_D2:
			regs.D2.Value = value
		case rarch_encoding.REG_D3:
			regs.D3.Value = value
		case rarch_encoding.REG_D4:
			regs.D4.Value = value
		case rarch_encoding.REG_D5:
			regs.D5.Value = value
		case rarch_encoding.REG_D6:
			regs.D6.Value = value
		case rarch_encoding.REG_D7:
			regs.D7.Value = value
		}
	case rarch_encoding.RegisterControl:
		switch r.Register {
		case rarch_encoding.REG_C0:
			regs.C0.Value = value
		case rarch_encoding.REG_C1:
			regs.C1.Value = value
		case rarch_encoding.REG_C2:
			regs.C2.Value = value
		case rarch_encoding.REG_C3:
			regs.C3.Value = value
		case rarch_encoding.REG_C4:
			regs.C4.Value = value
		case rarch_encoding.REG_C5:
			regs.C5.Value = value
		case rarch_encoding.REG_C6:
			regs.C6.Value = value
		case rarch_encoding.REG_C7:
			regs.C7.Value = value
		case rarch_encoding.REG_C8:
			regs.C8.Value = value
		case rarch_encoding.REG_C9:
			regs.C9.Value = value
		case rarch_encoding.REG_C10:
			regs.C10.Value = value
		case rarch_encoding.REG_C11:
			regs.C11.Value = value
		case rarch_encoding.REG_C12:
			regs.C12.Value = value
		case rarch_encoding.REG_C13:
			regs.C13.Value = value
		case rarch_encoding.REG_C14:
			regs.C14.Value = value
		case rarch_encoding.REG_C15:
			regs.C15.Value = value
		}
	}
}

func (regs *Registers) ReadReg(r rarch_encoding.Register) uint64 {
	switch r.Register_type {
	case rarch_encoding.RegisterGP:
		switch r.Register {
		case rarch_encoding.REG_R0:
			return 0
		case rarch_encoding.REG_R1:
			return regs.R1.Value
		case rarch_encoding.REG_R2:
			return regs.R2.Value
		case rarch_encoding.REG_R3:
			return regs.R3.Value
		case rarch_encoding.REG_R4:
			return regs.R4.Value
		case rarch_encoding.REG_R5:
			return regs.R5.Value
		case rarch_encoding.REG_R6:
			return regs.R6.Value
		case rarch_encoding.REG_R7:
			return regs.R7.Value
		case rarch_encoding.REG_R8:
			return regs.R8.Value
		case rarch_encoding.REG_R9:
			return regs.R9.Value
		case rarch_encoding.REG_R10:
			return regs.R10.Value
		case rarch_encoding.REG_R11:
			return regs.R11.Value
		case rarch_encoding.REG_R12:
			return regs.R12.Value
		case rarch_encoding.REG_R13:
			return regs.R13.Value
		case rarch_encoding.REG_SP:
			return regs.SP.Value
		case rarch_encoding.REG_BP:
			return regs.BP.Value
		}
	case rarch_encoding.RegisterIndex:
		switch r.Register {
		case rarch_encoding.REG_I0:
			return regs.I0.Value
		case rarch_encoding.REG_I1:
			return regs.I1.Value
		case rarch_encoding.REG_I2:
			return regs.I2.Value
		case rarch_encoding.REG_I3:
			return regs.I3.Value
		case rarch_encoding.REG_I4:
			return regs.I4.Value
		case rarch_encoding.REG_I5:
			return regs.I5.Value
		case rarch_encoding.REG_I6:
			return regs.I6.Value
		case rarch_encoding.REG_I7:
			return regs.I7.Value
		case rarch_encoding.REG_I8:
			return regs.I8.Value
		case rarch_encoding.REG_I9:
			return regs.I9.Value
		case rarch_encoding.REG_I10:
			return regs.I10.Value
		case rarch_encoding.REG_I11:
			return regs.I11.Value
		case rarch_encoding.REG_I12:
			return regs.I12.Value
		case rarch_encoding.REG_I13:
			return regs.I13.Value
		case rarch_encoding.REG_I14:
			return regs.I14.Value
		case rarch_encoding.REG_I15:
			return regs.I15.Value
		}
	case rarch_encoding.RegisterMMove:
		switch r.Register {
		case rarch_encoding.REG_S0:
			return regs.S0.Value
		case rarch_encoding.REG_S1:
			return regs.S1.Value
		case rarch_encoding.REG_S2:
			return regs.S2.Value
		case rarch_encoding.REG_S3:
			return regs.S3.Value
		case rarch_encoding.REG_S4:
			return regs.S4.Value
		case rarch_encoding.REG_S5:
			return regs.S5.Value
		case rarch_encoding.REG_S6:
			return regs.S6.Value
		case rarch_encoding.REG_S7:
			return regs.S7.Value
		case rarch_encoding.REG_D0:
			return regs.D0.Value
		case rarch_encoding.REG_D1:
			return regs.D1.Value
		case rarch_encoding.REG_D2:
			return regs.D2.Value
		case rarch_encoding.REG_D3:
			return regs.D3.Value
		case rarch_encoding.REG_D4:
			return regs.D4.Value
		case rarch_encoding.REG_D5:
			return regs.D5.Value
		case rarch_encoding.REG_D6:
			return regs.D6.Value
		case rarch_encoding.REG_D7:
			return regs.D7.Value
		}
	case rarch_encoding.RegisterControl:
		switch r.Register {
		case rarch_encoding.REG_C0:
			return regs.C0.Value
		case rarch_encoding.REG_C1:
			return regs.C1.Value
		case rarch_encoding.REG_C2:
			return regs.C2.Value
		case rarch_encoding.REG_C3:
			return regs.C3.Value
		case rarch_encoding.REG_C4:
			return regs.C4.Value
		case rarch_encoding.REG_C5:
			return regs.C5.Value
		case rarch_encoding.REG_C6:
			return regs.C6.Value
		case rarch_encoding.REG_C7:
			return regs.C7.Value
		case rarch_encoding.REG_C8:
			return regs.C8.Value
		case rarch_encoding.REG_C9:
			return regs.C9.Value
		case rarch_encoding.REG_C10:
			return regs.C10.Value
		case rarch_encoding.REG_C11:
			return regs.C11.Value
		case rarch_encoding.REG_C12:
			return regs.C12.Value
		case rarch_encoding.REG_C13:
			return regs.C13.Value
		case rarch_encoding.REG_C14:
			return regs.C14.Value
		case rarch_encoding.REG_C15:
			return regs.C15.Value
		}
	}

	return 0
}

type Memory struct {
	mem []byte

	rom []byte
}

func (m *Memory) PokeB(addr uint64, value byte) {
	if addr < rom_end {
		return
	}
	if addr >= mem_size {
		return
	}
	m.mem[addr] = value
}
func (m *Memory) PeekB(addr uint64) byte {
	if addr < rom_end {
		return m.rom[addr]
	}
	if addr >= mem_size {
		return 0
	}
	return m.mem[addr]
}

func (m *Memory) PokeS(addr uint64, value uint16) {
	m.PokeB(addr, byte(value>>8))
	m.PokeB(addr+1, byte(value))
}
func (m *Memory) PeekS(addr uint64) uint16 {
	b1 := uint16(m.PeekB(addr))
	b2 := uint16(m.PeekB(addr + 1))

	return (b1 << 8) | b2
}

func (m *Memory) PokeD(addr uint64, value uint32) {
	m.PokeS(addr, uint16(value>>16))
	m.PokeS(addr+2, uint16(value))
}
func (m *Memory) PeekD(addr uint64) uint32 {
	value := uint32(m.PeekS(addr)) << 16
	value |= uint32(m.PeekS(addr + 2))

	return value
}

func (m *Memory) PokeQ(addr uint64, value uint64) {
	m.PokeD(addr, uint32(value>>32))
	m.PokeD(addr+4, uint32(value))
}
func (m *Memory) PeekQ(addr uint64) uint64 {
	value := uint64(m.PeekD(addr)) << 32
	value |= uint64(m.PeekD(addr + 4))

	return value
}

type State struct {
	Registers Registers

	Memory Memory
}

const mem_size = 0x100000

func (s *State) Init(rom []byte) {
	s.Memory.mem = make([]byte, mem_size)
	s.Memory.rom = make([]byte, rom_end)
	copy(s.Memory.rom, rom)
}

func (s State) String() string {
	return fmt.Sprintf(`
PC: %.16X Status: %.16X
R1: %.16X R2: %.16X R3: %.16X R4: %.16X R5: %.16X R6: %.16X R7: %.16X R8: %.16X R9: %.16X R10: %.16X R11: %.16X R12: %.16X R13: %.16X
SP: %.16X BP: %.16X
I0: I1: %.16X I2: %.16X I3: %.16X I4: %.16X I5: %.16X I6: %.16X I7: %.16X I8: %.16X I9: %.16X I10: %.16X I11: %.16X I12: %.16X I13: %.16X I14: %.16X I15: %.16X
S0: S1: %.16X S2: %.16X S3: %.16X S4: %.16X S5: %.16X S6: %.16X S7: %.16X
D0: D1: %.16X D2: %.16X D3: %.16X D4: %.16X D5: %.16X D6: %.16X D7: %.16X`,
		s.Registers.PC.Value, s.Registers.Status.Value,
		s.Registers.R1.Value, s.Registers.R2.Value, s.Registers.R3.Value, s.Registers.R4.Value, s.Registers.R5.Value, s.Registers.R6.Value, s.Registers.R7.Value, s.Registers.R8.Value, s.Registers.R9.Value, s.Registers.R10.Value, s.Registers.R11.Value, s.Registers.R12.Value, s.Registers.R13.Value,
		s.Registers.SP.Value, s.Registers.BP.Value,
		s.Registers.I0.Value, s.Registers.I1.Value, s.Registers.I2.Value, s.Registers.I3.Value, s.Registers.I4.Value, s.Registers.I5.Value, s.Registers.I6.Value, s.Registers.I7.Value, s.Registers.I8.Value, s.Registers.I9.Value, s.Registers.I10.Value, s.Registers.I11.Value, s.Registers.I12.Value, s.Registers.I13.Value, s.Registers.I14.Value, s.Registers.I15.Value,
		s.Registers.S0.Value, s.Registers.S1.Value, s.Registers.S2.Value, s.Registers.S3.Value, s.Registers.S4.Value, s.Registers.S5.Value, s.Registers.S6.Value, s.Registers.S7.Value,
		s.Registers.D0.Value, s.Registers.D1.Value, s.Registers.D2.Value, s.Registers.D3.Value, s.Registers.D4.Value, s.Registers.D5.Value, s.Registers.D6.Value, s.Registers.D7.Value)
}
