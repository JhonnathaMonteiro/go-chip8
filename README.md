# CHIP-8 Golang Emulator (interpreter)

Chip-8 is a simple, interpreted, programming language which was first used on some do-it-yourself computer system in the late 1970s and early 1980s.
The COSMAC VIP, DREAM 6800, and ETI 660 computers are a few examples. Theses computers typically were designed to use a television as a display, had
between 1 and 4k of RAM, and used a 16-key hexadecimal keypad for input. The interpreter took up only 512 bytes of memory, and programs, which were
entered into the computer in hexadecimal, where even smaller.

## Dependecies

- [SDL2](https://github.com/veandco/go-sdl2#examples)

## Specifications

CHIP-8 has the following components:

- Memory: CHIP-8 is capable of accessing up to 4KB (4,096 bytes) of RAM, from location 0x000 (0) to 0xFFF (4095). The first 512 bytes, from 0x000 to
  0x1FF, are where the original interpreter was located, and should not be used by programs. Most CHIP-8 programs start at location 0x200 (512), but
  some begin at 0x600 (1536). Programs beginning at 0x600 are intended for the ETI 660 computer.

Memory Map:
+---------------+= 0xFFF (4095) End of Chip-8 RAM
|               |
|               |
|               |
|               |
|               |
| 0x200 to 0xFFF|
|     Chip-8    |
| Program / Data|
|     Space     |
|               |
|               |
|               |
+- - - - - - - -+= 0x600 (1536) Start of ETI 660 Chip-8 programs
|               |
|               |
|               |
+---------------+= 0x200 (512) Start of most Chip-8 programs
| 0x000 to 0x1FF|
| Reserved for  |
|  interpreter  |
+---------------+= 0x000 (0) Start of Chip-8 RAM
## CHIP-8 instructions

## Display

The original implementation of the Chip-8 language used a 64x32-pixel monochrome display.

## Keyboard layout mapping

|1|2|3|C|
|4|5|6|D|
|7|8|9|E|
|A|0|B|F|

into

|1|2|3|4|
|Q|W|E|R|
|A|S|D|F|
|Z|X|C|V|

## References

- [Cowgod's](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM)
- [Venkos](https://www.youtube.com/watch?v=MBWyVwyBMhk)

## TODO

[x] read rom
[x] learn sdl
    [x] keyboard
    [x] screen
[] Room loop
    [] how to set frequency
[] opcodes

