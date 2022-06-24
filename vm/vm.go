package vm

type VM struct {
	Registers [32]int32
	Program   []byte
	pc        byte
	remainder uint32
}

func RunVM(vm *VM) {
	is_done := false
	for !is_done {
		is_done = Execute_instruction(vm)
	}
}

func RunVM_once(vm *VM) {
	Execute_instruction(vm)
}

func Execute_instruction(vm *VM) bool {
	if vm.pc >= byte(len(vm.Program)) {
		return true
	}
	op := Decode_opcode(vm)
	switch op {
	case LOAD:
		register := byte(Next_8_bits(vm))
		value := uint32(Next_16_bits(vm))
		vm.Registers[register] = int32(value)

	case ADD:
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		vm.Registers[byte(Next_8_bits(vm))] = int32(register1 + register2)

	case SBT:
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		vm.Registers[byte(Next_8_bits(vm))] = int32(register1 - register2)

	case MUL:
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		vm.Registers[byte(Next_8_bits(vm))] = int32(register1 * register2)

	case DIV:
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		vm.Registers[byte(Next_8_bits(vm))] = int32(register1 / register2)
		vm.remainder = uint32(register1 % register2)

	case HLT:
		println("HLT encountered")
		return true

	case JMP:
		target := vm.Registers[byte(Next_8_bits(vm))]
		vm.pc = byte(target)

	case JMPF:
		value := vm.Registers[byte(Next_8_bits(vm))]
		vm.pc += byte(value)

	case JMPB:
		value := vm.Registers[byte(Next_8_bits(vm))]
		vm.pc -= byte(value)

	case EQ:
		var value int32 = 0
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		if register1 == register2 {
			value = 1
		}
		vm.Registers[byte(Next_8_bits(vm))] = value

	case NEQ:
		var value int32 = 0
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		if register1 != register2 {
			value = 1
		}
		vm.Registers[byte(Next_8_bits(vm))] = value

	case GT:
		var value int32 = 0
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		if register1 > register2 {
			value = 1
		}
		vm.Registers[byte(Next_8_bits(vm))] = value

	case GTE:
		var value int32 = 0
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		if register1 >= register2 {
			value = 1
		}
		vm.Registers[byte(Next_8_bits(vm))] = value

	case LT:
		var value int32 = 0
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		if register1 < register2 {
			value = 1
		}
		vm.Registers[byte(Next_8_bits(vm))] = value

	case LTE:
		var value int32 = 0
		register1 := vm.Registers[byte(Next_8_bits(vm))]
		register2 := vm.Registers[byte(Next_8_bits(vm))]
		if register1 <= register2 {
			value = 1
		}
		vm.Registers[byte(Next_8_bits(vm))] = value

	case JCN:
		cond := vm.Registers[byte(Next_8_bits(vm))]
		pointer := byte(vm.Registers[byte(Next_8_bits(vm))])

		if cond != 0 {
			vm.pc = pointer
		}

	default:
		println("Unknown opcode encountered. Terminating.")
		return true
	}
	return false
}

func Decode_opcode(vm *VM) Opcode {
	opcode := U8_to_op(vm.Program[vm.pc])
	vm.pc++
	return opcode
}

func Next_8_bits(vm *VM) byte {
	result := vm.Program[vm.pc]
	vm.pc++
	return byte(result)
}

func Next_16_bits(vm *VM) uint16 {
	result := uint16(vm.Program[vm.pc])<<8 | uint16(vm.Program[vm.pc+1])
	vm.pc += 2
	return result
}

func Add_byte(vm *VM, b byte) {
	vm.Program = append(vm.Program, b)
}

func Add_bytes(vm *VM, b []byte) {
	vm.Program = append(vm.Program, b...)
}

func Get_test_vm() VM {
	test_vm := VM{}
	test_vm.Registers[0] = 5
	test_vm.Registers[1] = 10
	return test_vm
}
