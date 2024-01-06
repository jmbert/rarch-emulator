package instructions

import (
	"errors"
	"fmt"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Load(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatC)

	register := rarch_encoding.DecodeRegister(instr.Register)
	prefix := instr.Prefix

	switch prefix {
	case rarch_encoding.PREF_LITERAL:
		value := instr.Immediate.Value

		s.Registers.WriteReg(register, value)

		break
	case rarch_encoding.PREF_ABSOLUTE:

	case rarch_encoding.PREF_INDREL:
		index := instr.Immediate.Value

		var basereg rarch_encoding.Register
		basereg.Register_size = rarch_encoding.RegisterLen64
		basereg.Register_type = rarch_encoding.RegisterIndex
		basereg.Register = byte(s.Registers.Status.Value>>rarch_encoding.STATUS_INDEX_BIT) & 0xF
		base := s.Registers.ReadReg(basereg)

		size := uint64(1 << ((s.Registers.Status.Value >> rarch_encoding.STATUS_OFFSIZE_BIT) & 0x3))

		addr := base + (size * index)

		var value uint64
		switch register.Register_size {
		case rarch_encoding.RegisterLen8:
			value = uint64(s.Memory.PeekB(addr))
		case rarch_encoding.RegisterLen16:
			value = uint64(s.Memory.PeekS(addr))
		case rarch_encoding.RegisterLen32:
			value = uint64(s.Memory.PeekD(addr))
		case rarch_encoding.RegisterLen64:
			value = uint64(s.Memory.PeekQ(addr))
		}
		s.Registers.WriteReg(register, value)
	default:
		return errors.New(fmt.Sprintf("Prefix 0x%.2X not supported", prefix))
	}

	return nil
}
