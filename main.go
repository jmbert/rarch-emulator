package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmbert/rarch-emulator/state"
	rarch_encoding "github.com/jmbert/rarch-encoding"
)

func Emulate(file string) error {

	fcontents, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var s state.State
	s.Init(fcontents)

	for i := 0; ; i++ {
		pc := s.Registers.PC.Value
		var instrBuffer []byte = make([]byte, 11)
		for i := 0; i < len(instrBuffer); i++ {
			instrBuffer[i] = s.Memory.PeekB(pc + uint64(i))
		}
		instr, size := rarch_encoding.DecodeInstruction(instrBuffer)
		s.Registers.PC.Value += size
		err = RunInstr(instr, &s)
		if err != nil {
			return err
		}

		if (s.Registers.Status.Value & rarch_encoding.STATUS_HALTED) == rarch_encoding.STATUS_HALTED {
			break
		}
	}

	fmt.Printf("%v\n", s)

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
