package instructions

import (
	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Nop(i rarch_encoding.Instruction, s *state.State) error {
	return nil
}
