package main

import "github.com/veandco/go-sdl2/sdl"

type SDLKeyboard struct {
	scancodeIntoMap map[uint32]uint8
}

func (kb *SDLKeyboard) isKeyPressed(key uint32) bool {
	keyboardState := sdl.GetKeyboardState()
	return keyboardState[kb.scancodeIntoMap[key]] == 1
}

func (kb *SDLKeyboard) init() {
	kb.scancodeIntoMap = map[uint32]uint8{
		uint32(sdl.GetScancodeFromName("1")): uint8(sdl.GetScancodeFromName("1")),
		uint32(sdl.GetScancodeFromName("2")): uint8(sdl.GetScancodeFromName("2")),
		uint32(sdl.GetScancodeFromName("3")): uint8(sdl.GetScancodeFromName("3")),
		uint32(sdl.GetScancodeFromName("C")): uint8(sdl.GetScancodeFromName("4")),
		uint32(sdl.GetScancodeFromName("4")): uint8(sdl.GetScancodeFromName("Q")),
		uint32(sdl.GetScancodeFromName("5")): uint8(sdl.GetScancodeFromName("W")),
		uint32(sdl.GetScancodeFromName("6")): uint8(sdl.GetScancodeFromName("E")),
		uint32(sdl.GetScancodeFromName("D")): uint8(sdl.GetScancodeFromName("R")),
		uint32(sdl.GetScancodeFromName("7")): uint8(sdl.GetScancodeFromName("A")),
		uint32(sdl.GetScancodeFromName("8")): uint8(sdl.GetScancodeFromName("S")),
		uint32(sdl.GetScancodeFromName("9")): uint8(sdl.GetScancodeFromName("D")),
		uint32(sdl.GetScancodeFromName("E")): uint8(sdl.GetScancodeFromName("F")),
		uint32(sdl.GetScancodeFromName("A")): uint8(sdl.GetScancodeFromName("Z")),
		uint32(sdl.GetScancodeFromName("0")): uint8(sdl.GetScancodeFromName("X")),
		uint32(sdl.GetScancodeFromName("B")): uint8(sdl.GetScancodeFromName("C")),
		uint32(sdl.GetScancodeFromName("F")): uint8(sdl.GetScancodeFromName("V")),
	}
}
