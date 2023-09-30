package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	display := SDLDisplay{}
	var scalingFactor int32 = 12
	var bg uint32 = 0x00000000 // black
	var fg uint32 = 0xFFFFFFFF // white

	display.init(scalingFactor, bg, fg)
	defer display.destroy()

	display.drawPixel(2, 2)
	display.drawPixel(2, 3)
	//	display.drawPixel(2, 4)
	//	display.drawPixel(2, 5)
	display.update()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("QUIT")
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED {
					keycode := t.Keysym.Sym
					switch string(keycode) {
					case "c":
						fmt.Println("cleaning")
						display.clear()
						display.update()
					}
				}
			}
		}
	}

}
