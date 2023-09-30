package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type VM struct {
	pc         uint16
	sp         uint16
	I          uint16
	opcode     uint16
	clockspeed uint16
	drawflag   bool
	stack      [16]uint16
	screen     [32][8]uint8
	memory     [4096]uint8
	V          [16]uint8
}

type Display interface {
	clear()
	update()
	drawPixel(int32, int32)
	destroy()
}

type Keyboard interface {
	// TODO
	isKeyPressed(key uint8) bool
}

func newVM(romPath string) (*VM, error) {
	vm := &VM{
		pc: 0x200,
	}

	romData, err := readRom(romPath)
	if err != nil {
		return nil, err
	}
	copy(vm.memory[0x200:], romData)
	vm.initialiseFont()
	return vm, nil
}

func (vm *VM) initialiseFont() {
	//0x000-0x1FF - Chip 8 interpreter (contains font set in emu)
	//0x050-0x0A0 - Used for the built in 4x5 pixel font set (0-F)
	//0 0x050
	vm.memory[0x050] = 0xF0
	vm.memory[0x051] = 0x90
	vm.memory[0x052] = 0x90
	vm.memory[0x053] = 0x90
	vm.memory[0x054] = 0xF0
	// 1
	vm.memory[0x055] = 0x20
	vm.memory[0x056] = 0x60
	vm.memory[0x057] = 0x20
	vm.memory[0x058] = 0x20
	vm.memory[0x059] = 0x70
	// 2
	vm.memory[0x05A] = 0xF0
	vm.memory[0x05B] = 0x10
	vm.memory[0x05C] = 0xF0
	vm.memory[0x05D] = 0x80
	vm.memory[0x05E] = 0xF0
	// 3
	vm.memory[0x05F] = 0xF0
	vm.memory[0x060] = 0x10
	vm.memory[0x061] = 0xF0
	vm.memory[0x062] = 0x10
	vm.memory[0x063] = 0xF0
	// 4
	vm.memory[0x064] = 0x90
	vm.memory[0x065] = 0x90
	vm.memory[0x066] = 0xF0
	vm.memory[0x067] = 0x10
	vm.memory[0x068] = 0x10
	// 5
	vm.memory[0x069] = 0xF0
	vm.memory[0x06A] = 0x80
	vm.memory[0x06B] = 0xF0
	vm.memory[0x06C] = 0x10
	vm.memory[0x06D] = 0xF0
	// 6
	vm.memory[0x06E] = 0xF0
	vm.memory[0x06F] = 0x80
	vm.memory[0x070] = 0xF0
	vm.memory[0x071] = 0x90
	vm.memory[0x072] = 0xF0
	// 7
	vm.memory[0x073] = 0xF0
	vm.memory[0x074] = 0x10
	vm.memory[0x075] = 0x20
	vm.memory[0x076] = 0x40
	vm.memory[0x077] = 0x40
	// 8
	vm.memory[0x078] = 0xF0
	vm.memory[0x079] = 0x90
	vm.memory[0x07A] = 0xF0
	vm.memory[0x07B] = 0x90
	vm.memory[0x07C] = 0xF0
	// 9
	vm.memory[0x07D] = 0xF0
	vm.memory[0x07E] = 0x90
	vm.memory[0x07F] = 0xF0
	vm.memory[0x080] = 0x10
	vm.memory[0x081] = 0xF0
	// A
	vm.memory[0x082] = 0xF0
	vm.memory[0x083] = 0x90
	vm.memory[0x084] = 0xF0
	vm.memory[0x085] = 0x90
	vm.memory[0x086] = 0x90
	// B
	vm.memory[0x087] = 0xE0
	vm.memory[0x088] = 0x90
	vm.memory[0x089] = 0xE0
	vm.memory[0x08A] = 0x90
	vm.memory[0x08B] = 0xE0
	// C
	vm.memory[0x08C] = 0xF0
	vm.memory[0x08D] = 0x80
	vm.memory[0x08E] = 0x80
	vm.memory[0x08F] = 0x80
	vm.memory[0x090] = 0xF0
	// D
	vm.memory[0x091] = 0xE0
	vm.memory[0x092] = 0x90
	vm.memory[0x093] = 0x90
	vm.memory[0x094] = 0x90
	vm.memory[0x095] = 0xE0
	// E
	vm.memory[0x096] = 0xF0
	vm.memory[0x097] = 0x80
	vm.memory[0x098] = 0xF0
	vm.memory[0x099] = 0x80
	vm.memory[0x09A] = 0xF0
	// F
	vm.memory[0x09B] = 0xF0
	vm.memory[0x09C] = 0x80
	vm.memory[0x09D] = 0xF0
	vm.memory[0x09E] = 0x80
	vm.memory[0x09F] = 0x80
}

func (vm *VM) parseOpcode() bool {
	// TODO
	vm.drawflag = false
	vm.opcode = uint16(vm.memory[vm.pc])<<8 | uint16(vm.memory[vm.pc+1])
	switch vm.opcode & 0xF000 {
	case 0x0000:
		switch vm.opcode {
		// 00E0 - CLS
		// Clear the display
		case 0x00E0:
			vm.drawflag = true
			for x := 0; x < 32; x++ {
				for y := 0; y < 8; y++ {
					vm.screen[x][y] = 0
				}
			}
		case 0x00EE:
			// 00EE - RET
			// Return from a subroutine.
			// The interpreter sets the program counter to the address at the top of the stack, then subtracts 1 from the stack pointer.
			vm.pc = vm.stack[vm.sp]
			vm.sp--

		default:
			log.Fatal(fmt.Errorf("unkown %#.4X opcode at pc %#.4X", vm.opcode, vm.pc))
		}
	case 0x1000:
		// 1nnn - JP addr
		// Jump to location nnn.
		// The interpreter sets the program counter to nnn.
		vm.pc = 0x0FFF & vm.opcode
	case 0x2000:
		// 2nnn - CALL addr
		// Call subroutine at nnn.
		// The interpreter increments the stack pointer, then puts the current PC on the top of the stack. The PC is then set to nnn.
		vm.sp++
		vm.stack[vm.sp] = vm.pc
		vm.pc = 0x0FFF & vm.opcode
	case 0x3000:
		// 3xkk - SE Vx, byte
		// Skip next instruction if Vx = kk.
		// The interpreter compares register Vx to kk, and if they are equal, increments the program counter by 2.
		if vm.V[0x0F00&vm.opcode>>8] == uint8(0x00FF&vm.opcode) {
			vm.pc += 2
		}
	case 0x4000:
		// 4xkk - SNE Vx, byte
		// Skip next instruction if Vx != kk.
		// The interpreter compares register Vx to kk, and if they are not equal, increments the program counter by 2.
		if vm.V[0x0F00&vm.opcode>>8] != uint8(0x00FF&vm.opcode) {
			vm.pc += 2
		}
	case 0x5000:
		// 5xy0 - SE Vx, Vy
		// Skip next instruction if Vx = Vy.
		// The interpreter compares register Vx to register Vy, and if they are equal, increments the program counter by 2.
		if vm.V[0x0F00&vm.opcode>>8] == vm.V[0x00F0&vm.opcode>>4] {
			vm.pc += 2
		}
	case 0x6000:
		// 6xkk - LD Vx, byte
		// Set Vx = kk.
		// The interpreter puts the value kk into register Vx.
		vm.V[0x0F00&vm.opcode>>8] = uint8(0x00FF & vm.opcode)
	case 0x7000:
		// 7xkk - ADD Vx, byte
		// Set Vx = Vx + kk.
		// Adds the value kk to the value of register Vx, then stores the result in Vx.
		vm.V[0x0F00&vm.opcode>>8] += uint8(0x00FF & vm.opcode)
	case 0x8000:
		switch 0x000F & vm.opcode {
		case 0x0000:
			// 8xy0 - LD Vx, Vy
			// Set Vx = Vy.
			// Stores the value of register Vy in register Vx.
			vm.V[0x0F00&vm.opcode>>8] = vm.V[0x00F0&vm.opcode>>4]
		case 0x0001:
			// 8xy1 - OR Vx, Vy
			// Set Vx = Vx OR Vy.
			// Performs a bitwise OR on the values of Vx and Vy, then stores the result in Vx.
			vm.V[0x0F00&vm.opcode>>8] = vm.V[0x00F0&vm.opcode>>4] | vm.V[0x0F00&vm.opcode>>8]
		case 0x0002:
			// 8xy2 - AND Vx, Vy
			// Set Vx = Vx AND Vy.
			// Performs a bitwise AND on the values of Vx and Vy, then stores the result in Vx.
			vm.V[0x0F00&vm.opcode>>8] = vm.V[0x00F0&vm.opcode>>4] & vm.V[0x0F00&vm.opcode>>8]
		case 0x0003:
			// 8xy3 - XOR Vx, Vy
			// Set Vx = Vx XOR Vy.
			// Performs a bitwise exclusive OR on the values of Vx and Vy, then stores the result in Vx.
			vm.V[0x0F00&vm.opcode>>8] = vm.V[0x00F0&vm.opcode>>4] ^ vm.V[0x0F00&vm.opcode>>8]
		case 0x0004:
			// 8xy4 - ADD Vx, Vy
			// Set Vx = Vx + Vy, set VF = carry.
			// The values of Vx and Vy are added together. If the result is greater than 8 bits (i.e., > 255,) VF is set to 1, otherwise 0. Only the lowest 8 bits of the result are kept, and stored in Vx.
			sum := vm.V[0x00F0&vm.opcode>>4] + vm.V[0x0F00&vm.opcode>>8]
			if sum > 0x00FF {
				vm.V[0xF] = 1
			} else {
				vm.V[0xF] = 0
			}
			vm.V[0x0F00&vm.opcode>>8] = sum & 0x00FF
		case 0x0005:
			// 8xy5 - SUB Vx, Vy
			// Set Vx = Vx - Vy, set VF = NOT borrow.
			// If Vx > Vy, then VF is set to 1, otherwise 0. Then Vy is subtracted from Vx, and the results stored in Vx.
			if vm.V[0x0F00&vm.opcode>>8] > vm.V[0x00F0&vm.opcode>>4] {
				vm.V[0xF] = 1
			} else {
				vm.V[0xF] = 0
			}
			vm.V[0x0F00&vm.opcode>>8] -= vm.V[0x00F0&vm.opcode>>4]
		case 0x0006:
			// 8xy6 - SHR Vx {, Vy}
			// Set Vx = Vx SHR 1.
			// If the least-significant bit of Vx is 1, then VF is set to 1, otherwise 0. Then Vx is divided by 2.
			if vm.V[vm.opcode&0x0F00>>8]&1 == 1 {
				vm.V[0xF] = 1
			} else {
				vm.V[0xF] = 0
			}
			vm.V[vm.opcode&0x0F00>>8] >>= 1
		case 0x0007:
			// 8xy7 - SUBN Vx, Vy
			// Set Vx = Vy - Vx, set VF = NOT borrow.
			// If Vy > Vx, then VF is set to 1, otherwise 0. Then Vx is subtracted from Vy, and the results stored in Vx.
			if vm.V[vm.opcode&0x00F0>>4] > vm.V[vm.opcode&0x0F00>>8] {
				vm.V[0xF] = 1
			} else {
				vm.V[0xF] = 0
			}
			vm.V[vm.opcode&0x0F00>>8] = vm.V[vm.opcode&0x00F0>>4] - vm.V[vm.opcode&0x0F00>>8]
		case 0x000E:
			// 8xyE - SHL Vx {, Vy}
			// Set Vx = Vx SHL 1.
			// If the most-significant bit of Vx is 1, then VF is set to 1, otherwise to 0. Then Vx is multiplied by 2.
			if vm.V[vm.opcode&0x0F00>>8]&0x80 == 0x80 {
				vm.V[0xF] = 1
			} else {
				vm.V[0xF] = 0
			}
			vm.V[vm.opcode&0x0F00>>8] <<= 1
		default:
			log.Fatal(fmt.Errorf("unkown %#.4X opcode at pc %#.4X", vm.opcode, vm.pc))
		}
	case 0x9000:
		// 9xy0 - SNE Vx, Vy
		// Skip next instruction if Vx != Vy.
		// The values of Vx and Vy are compared, and if they are not equal, the program counter is increased by 2.
		if vm.V[vm.opcode&0x0F00>>8] != vm.V[vm.opcode&0x00F0>>4] {
			vm.pc += 2
		}
	case 0xA000:
		// Annn - LD I, addr
		// Set I = nnn.
		// The value of register I is set to nnn.
		vm.I = vm.opcode & 0x0FFF
	case 0xB000:
		// Bnnn - JP V0, addr
		// Jump to location nnn + V0.
		// The program counter is set to nnn plus the value of V0.
		vm.pc = vm.opcode&0x0FFF + uint16(vm.V[0])
	case 0xC000:
		// Cxkk - RND Vx, byte
		// Set Vx = random byte AND kk.
		// The interpreter generates a random number from 0 to 255, which is then ANDed with the value kk. The results are stored in Vx. See instruction 8xy2 for more information on AND.
		vm.V[0x0F00&vm.opcode>>8] = uint8(rand.Intn(256)) & uint8(0x00FF&vm.opcode)
	case 0xD000:
		// Dxyn - DRW Vx, Vy, nibble
		// Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
		// The interpreter reads n bytes from memory, starting at the address stored in I. These bytes are then displayed as sprites on screen at coordinates (Vx, Vy).
		// Sprites are XORed onto the existing screen. If this causes any pixels to be erased, VF is set to 1, otherwise it is set to 0.
		// If the sprite is positioned so part of it is outside the coordinates of the display, it wraps around to the opposite side of the screen.
		// See instruction 8xy3 for more information on XOR, and section 2.4, Display, for more information on the Chip-8 screen and sprites.
		vm.drawflag = true
		vx := vm.V[vm.opcode&0x0F00] % 64
		vy := vm.V[vm.opcode&0x00F0] % 32
		n := vm.opcode & 0x000F
		vm.V[0xF] = 0
		for i := uint16(0); i < n; i++ {
			currentByte := vm.memory[vm.I+i]
			// TOOD: i need to understand better the concept of sprites and how to draw them
		}
	case 0xE000:
		switch 0x00FF & vm.opcode {
		case 0x009E:
			// Ex9E - SKP Vx
			// Skip next instruction if key with the value of Vx is pressed.
			// Checks the keyboard, and if the key corresponding to the value of Vx is currently in the down position, PC is increased by 2.
		case 0x00A1:
			// ExA1 - SKNP Vx
			// Skip next instruction if key with the value of Vx is not pressed.
			// Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
		default:
			log.Fatal(fmt.Errorf("unkown %#.4X opcode at pc %#.4X", vm.opcode, vm.pc))
		}
	case 0xF000:
		switch 0x00FF {
		case 0x00A1:
			// ExA1 - SKNP Vx
			// Skip next instruction if key with the value of Vx is not pressed.
			// Checks the keyboard, and if the key corresponding to the value of Vx is currently in the up position, PC is increased by 2.
		case 0x0007:
			// Fx07 - LD Vx, DT
			// Set Vx = delay timer value.
			// The value of DT is placed into Vx.
		case 0x000A:
			// Fx0A - LD Vx, K
			// Wait for a key press, store the value of the key in Vx.
			// All execution stops until a key is pressed, then the value of that key is stored in Vx.
		case 0x0015:
			// Fx15 - LD DT, Vx
			// Set delay timer = Vx.
			// DT is set equal to the value of Vx.
		case 0x0018:
			// Fx18 - LD ST, Vx
			// Set sound timer = Vx.
			// ST is set equal to the value of Vx.
		case 0x001E:
			// Fx1E - ADD I, Vx
			// Set I = I + Vx.
			// The values of I and Vx are added, and the results are stored in I.
		case 0x0029:
			// Fx29 - LD F, Vx
			// Set I = location of sprite for digit Vx.
			// The value of I is set to the location for the hexadecimal sprite corresponding to the value of Vx. See section 2.4, Display, for more information on the Chip-8 hexadecimal font.
		case 0x0033:
			// Fx33 - LD B, Vx
			// Store BCD representation of Vx in memory locations I, I+1, and I+2.
			// The interpreter takes the decimal value of Vx, and places the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.
		case 0x0055:
			// Fx55 - LD [I], Vx
			// Store registers V0 through Vx in memory starting at location I.
			// The interpreter copies the values of registers V0 through Vx into memory, starting at the address in I.
		case 0x0065:
			// Fx65 - LD Vx, [I]
			// Read registers V0 through Vx from memory starting at location I.
			// The interpreter reads values from memory starting at location I into registers V0 through Vx.
		default:
			log.Fatal(fmt.Errorf("unkown %#.4X opcode at pc %#.4X", vm.opcode, vm.pc))
		}
	default:
		log.Fatal(fmt.Errorf("unkown %#.4X opcode at pc %X", vm.opcode, vm.pc))
	}

	vm.pc += 2

	return true
}

func (vm *VM) loop(display Display, keyboard Keyboard) {
	// ROOM LOOP
	running := true
	for running {
		running = vm.parseOpcode()
		time.Sleep(time.Duration(1e6/uint32(vm.clockspeed)) * time.Microsecond)
		// TODO: the rest of the loop
		if vm.drawflag {
			display.clear()
		}
	}
}
