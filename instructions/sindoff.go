package instructions

import (
	"errors"
	"fmt"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Sind(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatB)

	reg := rarch_encoding.DecodeRegister(instr.Register)

	if reg.Register_type != rarch_encoding.RegisterIndex || reg.Register_size != rarch_encoding.RegisterLen64 {
		return errors.New(fmt.Sprintf("Bad register: 0x%.2X", instr.Register))
	}

	s.Registers.Status.Value |= uint64(reg.Register) << rarch_encoding.STATUS_INDEX_BIT

	return nil
}
func Soff(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatB)

	reg := rarch_encoding.DecodeRegister(instr.Register)

	if reg.Register_size != rarch_encoding.RegisterLen8 {
		return errors.New(fmt.Sprintf("Bad register size: 0x%.2X", instr.Register))
	}

	s.Registers.Status.Value |= (s.Registers.ReadReg(reg) & 0x3) << rarch_encoding.STATUS_OFFSIZE_BIT

	return nil
}
