package vm

type Opcode byte

const (
	LOAD Opcode = iota
	ADD
	SBT
	MUL
	DIV
	HLT
	JMP
	JMPF
	JMPB
	EQ
	NEQ
	GT
	GTE
	LT
	LTE
	JCN
	IGL
)

var StringToOpcode = map[string]Opcode{
	"LOAD": LOAD,
	"ADD":  ADD,
	"SBT":  SBT,
	"MUL":  MUL,
	"DIV":  DIV,
	"HLT":  HLT,
	"JMP":  JMP,
	"JMPF": JMPF,
	"JMPB": JMPB,
	"EQ":   EQ,
	"NEQ":  NEQ,
	"GT":   GT,
	"GTE":  GTE,
	"LT":   LT,
	"LTE":  LTE,
	"JCN":  JCN,
	"IGL":  IGL,
}

func U8_to_op(program byte) Opcode {
	switch program {
	case 0:
		return LOAD

	case 1:
		return ADD

	case 2:
		return SBT

	case 3:
		return MUL

	case 4:
		return DIV

	case 5:
		return HLT

	case 6:
		return JMP

	case 7:
		return JMPF

	case 8:
		return JMPB

	case 9:
		return EQ

	case 10:
		return NEQ

	case 11:
		return GT

	case 12:
		return GTE

	case 13:
		return LT

	case 14:
		return LTE

	case 15:
		return JCN

	default:
		return IGL
	}
}
