package instructions

import (
	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Jump(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatB)

	register := rarch_encoding.DecodeRegister(instr.Register)

	s.Registers.PC.Value = s.Registers.ReadReg(register)

	return nil
}
