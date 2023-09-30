package main

import "testing"

func mockVM(romBytes []uint8) VM {
	vm := VM{
		pc: 0x200,
	}
	copy(vm.memory[0x200:], romBytes)
	return vm
}

func TestLoadRomIntoMemory(t *testing.T) {
	vm, err := newVM("../roms/testrom/BC_test.ch8")
	if err != nil {
		t.Error(err)
	}

	if vm.memory[200] != 0 {
		t.Errorf("expected vm.memory[200] to be %#.2x, found %#.2x", 0, vm.memory[200])
	}
}

func Test00E0(t *testing.T) {
	cleanScreen := [32][8]uint8{}
	initialScreen := [32][8]uint8{}
	initialScreen[1][1] = 1
	vm := mockVM([]uint8{
		0x00,
		0xE0,
	})

	vm.screen = initialScreen
	parsed := vm.parseOpcode()
	if !parsed {
		t.Error("unable to parse opcode")
	}

	if vm.screen != cleanScreen {
		t.Error("vm.screen != cleanScreen")
	}
}

func Test00EE(t *testing.T) {
}
