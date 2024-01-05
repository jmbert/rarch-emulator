package instructions

import (
	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

var InstructionMap map[byte]func(i rarch_encoding.Instruction, s *state.State) error = map[byte]func(i rarch_encoding.Instruction, s *state.State) error{
	rarch_encoding.OP_NOP: Nop,
	rarch_encoding.OP_HLT: Halt,

	rarch_encoding.OP_LD:  Load,
	rarch_encoding.OP_JMP: Jump,
}
