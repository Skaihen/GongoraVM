package assembler

import "Skaihen/Plep/vm"

type Op struct {
	Code vm.Opcode
}

type Token struct {
	Op Op
}
