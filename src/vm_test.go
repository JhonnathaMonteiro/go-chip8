package main

import "testing"

func TestLoadRomIntoMemory(t *testing.T) {
	vm, err := newVM("../roms/testrom/BC_test.ch8")
	if err != nil {
		t.Error(err)
	}

	if vm.memory[200] != 0 {
		t.Errorf("expected vm.memory[200] to be %#.2x, found %#.2x", 0, vm.memory[200])
	}
}
