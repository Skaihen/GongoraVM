package vm

import (
	"testing"
)

func TestCreateHLT(t *testing.T) {
	opcode := HLT
	if opcode != HLT {
		t.Errorf("Expected 5, got %d", opcode)
	}
}
