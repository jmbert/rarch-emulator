package instructions

import (
	"errors"
	"fmt"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func GetRegister(r rarch_encoding.Register, s *state.Registers) *state.Register {
	switch r.Register_type {
	case rarch_encoding.RegisterGP:
		switch r.Register {
		case rarch_encoding.REG_R0:
			return &(s.R0)
		case rarch_encoding.REG_R1:
			return &(s.R1)
		case rarch_encoding.REG_R2:
			return &(s.R2)
		case rarch_encoding.REG_R3:
			return &(s.R3)
		case rarch_encoding.REG_R4:
			return &(s.R4)
		case rarch_encoding.REG_R5:
			return &(s.R5)
		case rarch_encoding.REG_R6:
			return &(s.R6)
		case rarch_encoding.REG_R7:
			return &(s.R7)
		case rarch_encoding.REG_R8:
			return &(s.R8)
		case rarch_encoding.REG_R9:
			return &(s.R9)
		case rarch_encoding.REG_R10:
			return &(s.R10)
		case rarch_encoding.REG_R11:
			return &(s.R11)
		case rarch_encoding.REG_R12:
			return &(s.R12)
		case rarch_encoding.REG_R13:
			return &(s.R13)
		case rarch_encoding.REG_SP:
			return &(s.SP)
		case rarch_encoding.REG_BP:
			return &(s.BP)
		}
	case rarch_encoding.RegisterIndex:
		switch r.Register {
		case rarch_encoding.REG_I0:
			return &(s.I0)
		case rarch_encoding.REG_I1:
			return &(s.I1)
		case rarch_encoding.REG_I2:
			return &(s.I2)
		case rarch_encoding.REG_I3:
			return &(s.I3)
		case rarch_encoding.REG_I4:
			return &(s.I4)
		case rarch_encoding.REG_I5:
			return &(s.I5)
		case rarch_encoding.REG_I6:
			return &(s.I6)
		case rarch_encoding.REG_I7:
			return &(s.I7)
		case rarch_encoding.REG_I8:
			return &(s.I8)
		case rarch_encoding.REG_I9:
			return &(s.I9)
		case rarch_encoding.REG_I10:
			return &(s.I10)
		case rarch_encoding.REG_I11:
			return &(s.I11)
		case rarch_encoding.REG_I12:
			return &(s.I12)
		case rarch_encoding.REG_I13:
			return &(s.I13)
		case rarch_encoding.REG_I14:
			return &(s.I14)
		case rarch_encoding.REG_I15:
			return &(s.I15)
		}
	case rarch_encoding.RegisterMMove:
		switch r.Register {
		case rarch_encoding.REG_S0:
			return &(s.S0)
		case rarch_encoding.REG_S1:
			return &(s.S1)
		case rarch_encoding.REG_S2:
			return &(s.S2)
		case rarch_encoding.REG_S3:
			return &(s.S3)
		case rarch_encoding.REG_S4:
			return &(s.S4)
		case rarch_encoding.REG_S5:
			return &(s.S5)
		case rarch_encoding.REG_S6:
			return &(s.S6)
		case rarch_encoding.REG_S7:
			return &(s.S7)
		case rarch_encoding.REG_D0:
			return &(s.D0)
		case rarch_encoding.REG_D1:
			return &(s.D1)
		case rarch_encoding.REG_D2:
			return &(s.D2)
		case rarch_encoding.REG_D3:
			return &(s.D3)
		case rarch_encoding.REG_D4:
			return &(s.D4)
		case rarch_encoding.REG_D5:
			return &(s.D5)
		case rarch_encoding.REG_D6:
			return &(s.D6)
		case rarch_encoding.REG_D7:
			return &(s.D7)
		}
	case rarch_encoding.RegisterControl:
		switch r.Register {
		case rarch_encoding.REG_C0:
			return &(s.C0)
		case rarch_encoding.REG_C1:
			return &(s.C1)
		case rarch_encoding.REG_C2:
			return &(s.C2)
		case rarch_encoding.REG_C3:
			return &(s.C3)
		case rarch_encoding.REG_C4:
			return &(s.C4)
		case rarch_encoding.REG_C5:
			return &(s.C5)
		case rarch_encoding.REG_C6:
			return &(s.C6)
		case rarch_encoding.REG_C7:
			return &(s.C7)
		case rarch_encoding.REG_C8:
			return &(s.C8)
		case rarch_encoding.REG_C9:
			return &(s.C9)
		case rarch_encoding.REG_C10:
			return &(s.C10)
		case rarch_encoding.REG_C11:
			return &(s.C11)
		case rarch_encoding.REG_C12:
			return &(s.C12)
		case rarch_encoding.REG_C13:
			return &(s.C13)
		case rarch_encoding.REG_C14:
			return &(s.C14)
		case rarch_encoding.REG_C15:
			return &(s.C15)
		}
	}

	return nil
}

func Load(i rarch_encoding.Instruction, s *state.State) error {
	instr := i.(rarch_encoding.FormatC)

	register := rarch_encoding.DecodeRegister(instr.Register)
	prefix := instr.Prefix

	switch prefix {
	case rarch_encoding.PREF_LITERAL:
		value := instr.Immediate.Value

		registerState := GetRegister(register, &s.Registers)
		registerState.Value = value

		break
	default:
		return errors.New(fmt.Sprintf("Prefix 0x%.2X not supported", prefix))
	}

	return nil
}
