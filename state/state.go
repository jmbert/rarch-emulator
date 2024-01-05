package state

type Register struct {
	Value uint64
}

type Registers struct {
	PC Register

	Status Register

	R0  Register
	R1  Register
	R2  Register
	R3  Register
	R4  Register
	R5  Register
	R6  Register
	R7  Register
	R8  Register
	R9  Register
	R10 Register
	R11 Register
	R12 Register
	R13 Register

	SP Register
	BP Register

	I0  Register
	I1  Register
	I2  Register
	I3  Register
	I4  Register
	I5  Register
	I6  Register
	I7  Register
	I8  Register
	I9  Register
	I10 Register
	I11 Register
	I12 Register
	I13 Register
	I14 Register
	I15 Register

	S0 Register
	S1 Register
	S2 Register
	S3 Register
	S4 Register
	S5 Register
	S6 Register
	S7 Register

	D0 Register
	D1 Register
	D2 Register
	D3 Register
	D4 Register
	D5 Register
	D6 Register
	D7 Register

	C0  Register
	C1  Register
	C2  Register
	C3  Register
	C4  Register
	C5  Register
	C6  Register
	C7  Register
	C8  Register
	C9  Register
	C10 Register
	C11 Register
	C12 Register
	C13 Register
	C14 Register
	C15 Register
}

type Memory struct {
	mem []byte
}

func (m *Memory) PokeB(addr uint64, value byte) {
	m.mem[addr] = value
}
func (m *Memory) PeekB(addr uint64) byte {
	return m.mem[addr]
}

type State struct {
	Registers Registers

	Memory Memory
}

const mem_size = 0x100000

func (s *State) Init() {
	s.Memory.mem = make([]byte, mem_size)
}
