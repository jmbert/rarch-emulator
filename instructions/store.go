package instructions

import (
	"errors"
	"fmt"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Store(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatC)

	//register := rarch_encoding.DecodeRegister(instr.Register)
	prefix := instr.Prefix

	switch prefix {

	default:
		return errors.New(fmt.Sprintf("Prefix 0x%.2X not supported", prefix))
	}
}
