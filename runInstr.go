package main

import (
	"errors"
	"fmt"

	"github.com/jmbert/rarch-emulator/instructions"
	"github.com/jmbert/rarch-emulator/state"
	rencoding "github.com/jmbert/rarch-encoding"
)

func RunInstr(i rencoding.Instruction, s *state.State) error {
	opcode := i.Encode()[1]
	instr, ok := instructions.InstructionMap[opcode]

	if !ok {
		return errors.New(fmt.Sprintf("No such instruction 0x%.2X", opcode))
	}
	return instr(i, s)
}
