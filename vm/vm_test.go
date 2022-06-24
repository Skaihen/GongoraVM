package vm

import (
	"testing"
)

func TestCreateVM(t *testing.T) {
	test_vm := VM{}
	if test_vm.Registers[0] != 0 {
		t.Errorf("Expected 0, got %d", test_vm.Registers[0])
	}

	if test_vm.pc != 0 {
		t.Errorf("Expected 0, got %d", test_vm.pc)
	}
}

func TestGetTestVM(t *testing.T) {
	test_vm := Get_test_vm()
	if test_vm.Registers[0] != 5 {
		t.Errorf("Expected 5, got %d", test_vm.Registers[0])
	}
	if test_vm.Registers[1] != 10 {
		t.Errorf("Expected 10, got %d", test_vm.Registers[1])
	}
}

func TestHLTOpcode(t *testing.T) {
	test_vm := VM{}
	test_vm.Program = []byte{5, 0, 0, 0}
	RunVM_once(&test_vm)
	if test_vm.pc != 1 {
		t.Errorf("Expected 1, got %d", test_vm.pc)
	}
}

func TestIGLOpcode(t *testing.T) {
	test_vm := VM{}
	test_vm.Program = []byte{200, 0, 0, 0}
	RunVM_once(&test_vm)
	if test_vm.pc != 1 {
		t.Errorf("Expected 1, got %d", test_vm.pc)
	}
}

func TestLOADOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Program = []byte{0, 0, 1, 244}
	RunVM(&test_vm)
	if test_vm.Registers[0] != 500 {
		t.Errorf("Expected 500, got %d", test_vm.Registers[0])
	}
}

func TestJMPOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Registers[0] = 4
	test_vm.Program = []byte{6, 0, 0, 0}
	RunVM_once(&test_vm)
	if test_vm.pc != 4 {
		t.Errorf("Expected 4, got %d", test_vm.pc)
	}
}

func TestJMPFOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Registers[0] = 2
	test_vm.Program = []byte{7, 0, 0, 0, 5, 0, 0, 0}
	RunVM(&test_vm)
	if test_vm.pc != 5 {
		t.Errorf("Expected 5, got %d", test_vm.pc)
	}
}

func TestJMPBOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Registers[0] = 2
	test_vm.Registers[3] = 4
	test_vm.Program = []byte{7, 0, 0, 5, 0, 8, 3, 0}
	RunVM(&test_vm)
	if test_vm.pc != 8 {
		t.Errorf("Expected 8, got %d", test_vm.pc)
	}
}

func TestADDOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Program = []byte{1, 0, 1, 2}
	RunVM_once(&test_vm)
	if test_vm.Registers[2] != 15 {
		t.Errorf("Expected 15, got %d", test_vm.Registers[2])
	}
}

func TestEQOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Registers[0] = 10
	test_vm.Program = []byte{9, 0, 1, 2, 9, 0, 1, 3}
	RunVM_once(&test_vm)
	if test_vm.Registers[2] != 1 {
		t.Errorf("Expected 1, got %d", test_vm.Registers[2])
	}
	test_vm.Registers[0] = 20
	RunVM_once(&test_vm)
	if test_vm.Registers[3] != 0 {
		t.Errorf("Expected 0, got %d", test_vm.Registers[3])
	}
}

func TestJCNOpcode(t *testing.T) {
	test_vm := Get_test_vm()
	test_vm.Registers[0] = 10
	test_vm.Registers[3] = 11
	test_vm.Program = []byte{9, 0, 1, 2, 0, 15, 2, 3, 2, 5, 0, 0, 5}
	RunVM(&test_vm)
	if test_vm.pc != 13 {
		t.Errorf("Expected 13, got %d", test_vm.pc)
	}
}

func TestLoop(t *testing.T) {
	test_vm := VM{}
	test_vm.Program = []byte{00, 01, 00, 01, 00, 02, 00, 10, 00, 03, 00, 12, 01, 00, 01, 00, 10, 00, 02, 04, 15, 04, 03, 05}
	RunVM(&test_vm)
	if test_vm.Registers[0] != 10 {
		t.Errorf("Expected 10, got %d", test_vm.Registers[0])
	}
}
