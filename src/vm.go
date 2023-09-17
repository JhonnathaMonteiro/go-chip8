package main

type VM struct {
	pc     uint16
	memory [4096]uint8
	V      [16]uint8 // General purpose 8-bit registers
}

func newVM(roomPath string) (*VM, error) {
	vm := &VM{
		pc: 0x200,
	}

	roomData, err := readRoom(roomPath)
	if err != nil {
		return nil, err
	}
	copy(vm.memory[200:], roomData)
	return vm, nil
}
