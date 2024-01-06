package instructions

import (
	"errors"
	"fmt"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func store_reg(addr uint64, register rarch_encoding.Register, s *state.State) {

	var value uint64
	switch register.Register_size {
	case rarch_encoding.RegisterLen8:
		s.Memory.PokeB(addr, byte(value))
	case rarch_encoding.RegisterLen16:
		s.Memory.PokeS(addr, uint16(value))
	case rarch_encoding.RegisterLen32:
		s.Memory.PokeD(addr, uint32(value))
	case rarch_encoding.RegisterLen64:
		s.Memory.PokeQ(addr, uint64(value))
	}
}

func Store(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatC)

	register := rarch_encoding.DecodeRegister(instr.Register)
	prefix := instr.Prefix

	switch prefix {
	case rarch_encoding.PREF_ABSOLUTE:
		addr := instr.Immediate.Value

		store_reg(addr, register, s)
	case rarch_encoding.PREF_PCREL:
		offset := int64(instr.Immediate.Value)

		base := s.Registers.PC.Value

		var addr uint64

		if offset < 0 {
			addr = base - uint64(-offset)
		} else {
			addr = base + uint64(offset)
		}

		store_reg(addr, register, s)

	case rarch_encoding.PREF_INDREL:
		index := instr.Immediate.Value

		var basereg rarch_encoding.Register
		basereg.Register_size = rarch_encoding.RegisterLen64
		basereg.Register_type = rarch_encoding.RegisterIndex
		basereg.Register = byte(s.Registers.Status.Value>>rarch_encoding.STATUS_INDEX_BIT) & 0xF
		base := s.Registers.ReadReg(basereg)

		size := uint64(1 << ((s.Registers.Status.Value >> rarch_encoding.STATUS_OFFSIZE_BIT) & 0x3))

		addr := base + (size * index)

		store_reg(addr, register, s)
	default:
		return errors.New(fmt.Sprintf("Prefix 0x%.2X not supported", prefix))
	}

	return nil
}
