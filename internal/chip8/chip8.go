package chip8

const START_ADDRESS = 0x200
const FONTSET_START_ADDRESS = 0x50
const VIDEO_WIDTH = 64
const VIDEO_HEIGHT = 32

var chip8FontSet = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

type Chip8 struct {
	memory     [4096]byte                       //4K memory
	v          [16]byte                         //16 registers (V0-VF)
	index      uint16                           //index register
	pc         uint16                           //program counter
	stack      [16]uint16                       //16-level stack for pc
	sp         byte                             //stack pointer
	delayTimer byte                             //delay timer
	soundTimer byte                             //sound timer
	key        [16]byte                         //input key state
	gfx        [VIDEO_WIDTH * VIDEO_HEIGHT]byte //display (64x32)
}

func NewChip8() *Chip8 {
	chip := &Chip8{
		pc: 0x200,
	}

	return chip
}

func (c *Chip8) LoadROM(rom []byte) {
	for i, b := range rom {
		c.memory[START_ADDRESS+i] = b
	}
}

func (c *Chip8) LoadFontSet() {
	for i, b := range chip8FontSet {
		c.memory[FONTSET_START_ADDRESS+i] = b
	}
}
