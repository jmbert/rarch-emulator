package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Emulate(file string) error {

	var s state.State
	s.Init()

	fcontents, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	for {
		instr, size := rarch_encoding.DecodeInstruction(fcontents[s.Registers.PC.Value:])
		s.Registers.PC.Value += size
		err = RunInstr(instr, &s)
		if err != nil {
			return err
		}

		if (s.Registers.Status.Value & rarch_encoding.STATUS_HALTED) == rarch_encoding.STATUS_HALTED {
			break
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide input file")
		os.Exit(1)
	}
	err := Emulate(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
