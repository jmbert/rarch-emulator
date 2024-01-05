package instructions

import (
	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Halt(i rarch_encoding.Instruction, s *state.State) error {
	s.Registers.Status.Value |= rarch_encoding.STATUS_HALTED
	return nil
}
