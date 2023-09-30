package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SDLDisplay struct {
	window        *sdl.Window
	surface       *sdl.Surface
	scalingFactor int32
	bg            uint32
	fg            uint32
}

func (display *SDLDisplay) init(scalingFactor int32, bg, fg uint32) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	display.window, err = sdl.CreateWindow("GO-CHIP-8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		scalingFactor*64, scalingFactor*32, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	display.surface, err = display.window.GetSurface()
	if err != nil {
		panic(err)
	}

	display.bg = bg
	display.fg = fg
	display.scalingFactor = scalingFactor
}

func (display *SDLDisplay) clear() {
	err := display.surface.FillRect(nil, display.bg)
	if err != nil {
		panic(err)
	}
}

func (display *SDLDisplay) update() {
	// TODO
	err := display.window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func (display *SDLDisplay) drawPixel(x, y int32) {
	// TODO
	rect := sdl.Rect{
		X: x * display.scalingFactor,
		Y: y * display.scalingFactor,
		W: display.scalingFactor,
		H: display.scalingFactor,
	}
	err := display.surface.FillRect(&rect, display.fg)
	if err != nil {
		panic(err)
	}
}

func (display *SDLDisplay) destroy() {
	display.window.Destroy()
	sdl.Quit() // this probably should go into another place
}
