package chip8

const START_ADDRESS = 0x200
const FONTSET_START_ADDRESS = 0x50
const VIDEO_WIDTH int32 = 64
const VIDEO_HEIGHT int32 = 32

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

type OpcodeFunc func(opcode uint16)

type Chip8 struct {
	memory     [4096]byte   //4K memory
	v          [16]byte     //16 registers (V0-VF)
	index      uint16       //index register
	pc         uint16       //program counter
	stack      [16]uint16   //16-level stack for pc
	sp         byte         //stack pointer
	delayTimer byte         //delay timer
	soundTimer byte         //sound timer
	key        [16]byte     //input key state
	gfx        [32][64]byte //display (64x32)
	draw       bool

	table  [16]OpcodeFunc  //main opcode table
	table0 [16]OpcodeFunc  //table for opcodes starting with 0
	table8 [16]OpcodeFunc  //table for opcodes starting with 8
	tableE [16]OpcodeFunc  //table for opcodes starting with E
	tableF [256]OpcodeFunc //table for opcodes starting with F
}

func NewChip8() *Chip8 {
	c := &Chip8{
		pc: 0x200,
	}

	c.loadFontSet()

	//initialize the main function table
	c.table[0x1] = c.OP_1NNN
	c.table[0x2] = c.OP_2NNN
	c.table[0x3] = c.OP_3XKK
	c.table[0x4] = c.OP_4XKK
	c.table[0x5] = c.OP_5XY0
	c.table[0x6] = c.OP_6XKK
	c.table[0x7] = c.OP_7XKK
	c.table[0x9] = c.OP_9XY0
	c.table[0xA] = c.OP_ANNN
	c.table[0xB] = c.OP_BNNN
	c.table[0xC] = c.OP_CXKK
	c.table[0xD] = c.OP_DXYN

	//initialize table0
	for i := range c.table0 {
		c.table0[i] = c.OP_NULL
	}

	c.table0[0x0] = c.OP_00E0
	c.table0[0xE] = c.OP_00EE

	//initialize table8
	for i := range c.table8 {
		c.table8[i] = c.OP_NULL
	}

	c.table8[0x0] = c.OP_8XY0
	c.table8[0x1] = c.OP_8XY1
	c.table8[0x2] = c.OP_8XY2
	c.table8[0x3] = c.OP_8XY3
	c.table8[0x4] = c.OP_8XY4
	c.table8[0x5] = c.OP_8XY5
	c.table8[0x6] = c.OP_8XY6
	c.table8[0x7] = c.OP_8XY7
	c.table8[0xE] = c.OP_8XYE

	//initialize tableE
	for i := range c.tableE {
		c.tableE[i] = c.OP_NULL
	}

	c.tableE[0x1] = c.OP_EXA1
	c.tableE[0xE] = c.OP_EX9E

	//initialize tableF
	for i := range c.tableF {
		c.tableF[i] = c.OP_NULL
	}

	c.tableF[0x07] = c.OP_FX07
	c.tableF[0x0A] = c.OP_FX0A
	c.tableF[0x15] = c.OP_FX15
	c.tableF[0x18] = c.OP_FX18
	c.tableF[0x1E] = c.OP_FX1E
	c.tableF[0x29] = c.OP_FX29
	c.tableF[0x33] = c.OP_FX33
	c.tableF[0x55] = c.OP_FX55
	c.tableF[0x65] = c.OP_FX65

	return c
}

func (c *Chip8) LoadROM(rom []byte) {
	for i, b := range rom {
		c.memory[START_ADDRESS+i] = b
	}
}

func (c *Chip8) loadFontSet() {
	for i, b := range chip8FontSet {
		c.memory[FONTSET_START_ADDRESS+i] = b
	}
}

func (c *Chip8) Cycle() {
	opcode := c.FetchOpcode()

	c.pc += 2

	c.DecodeAndExecute(opcode)

}

func (c *Chip8) UpdateTimers() {
	if c.delayTimer > 0 {
		c.delayTimer--
	}

	if c.soundTimer > 0 {
		c.soundTimer--
	}
}

func (c *Chip8) Display() [32][64]byte {
	return c.gfx
}

func (c *Chip8) Draw() bool {
	temp := c.draw
	c.draw = false
	return temp
}
