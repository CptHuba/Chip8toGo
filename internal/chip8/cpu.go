package chip8

import (
	"math/rand"
)

func (c *Chip8) FetchOpcode() uint16 {
	return uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
}

func (c *Chip8) DecodeAndExecute(opcode uint16) {
	switch opcode & 0xF000 {
	case 0x0000:
		c.table0[opcode&0x000F](opcode)
	case 0x1000:
		c.table[opcode>>12](opcode)
	case 0x2000:
		c.table[opcode>>12](opcode)
	case 0x3000:
		c.table[opcode>>12](opcode)
	case 0x4000:
		c.table[opcode>>12](opcode)
	case 0x5000:
		c.table[opcode>>12](opcode)
	case 0x6000:
		c.table[opcode>>12](opcode)
	case 0x7000:
		c.table[opcode>>12](opcode)
	case 0x8000:
		c.table8[opcode&0x000F](opcode)
	case 0x9000:
		c.table[opcode>>12](opcode)
	case 0xA000:
		c.table[opcode>>12](opcode)
	case 0xB000:
		c.table[opcode>>12](opcode)
	case 0xC000:
		c.table[opcode>>12](opcode)
	case 0xD000:
		c.table[opcode>>12](opcode)
	case 0xE000:
		c.tableE[opcode&0x000F](opcode)
	case 0xF000:
		c.tableF[opcode&0x00FF](opcode)
	default:
		c.OP_NULL(opcode)
	}
}

func (c *Chip8) OP_00E0(opcode uint16) { //CLS
	for i := 0; i < len(c.gfx); i++ {
		for j := 0; j < len(c.gfx[i]); j++ {
			c.gfx[i][j] = 0
		}
	}
	c.draw = true
}

func (c *Chip8) OP_00EE(opcode uint16) { //RET
	c.sp--
	c.pc = c.stack[c.sp]
}

func (c *Chip8) OP_1NNN(opcode uint16) { //JP addr
	c.pc = opcode & 0x0FFF
}

func (c *Chip8) OP_2NNN(opcode uint16) { //CALL addr
	c.stack[c.sp] = c.pc
	c.sp++
	c.pc = opcode & 0x0FFF
}

func (c *Chip8) OP_3XKK(opcode uint16) { //SE Vx, byte
	x := (opcode & 0x0F00) >> 8
	kk := byte(opcode & 0x00FF)

	if kk == c.v[x] {
		c.pc += 2
	}
}

func (c *Chip8) OP_4XKK(opcode uint16) { //SNE Vx, byte
	x := (opcode & 0x0F00) >> 8
	kk := byte(opcode & 0x00FF)

	if kk != c.v[x] {
		c.pc += 2
	}
}

func (c *Chip8) OP_5XY0(opcode uint16) { //SE Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	if c.v[x] == c.v[y] {
		c.pc += 2
	}
}

func (c *Chip8) OP_6XKK(opcode uint16) { //LD Vx, byte
	x := (opcode & 0x0F00) >> 8
	kk := byte(opcode & 0x00FF)

	c.v[x] = kk
}

func (c *Chip8) OP_7XKK(opcode uint16) { //ADD Vx, byte
	x := (opcode & 0x0F00) >> 8
	kk := byte(opcode & 0x00FF)

	c.v[x] += kk
}

func (c *Chip8) OP_8XY0(opcode uint16) { //LD Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] = c.v[y]
}

func (c *Chip8) OP_8XY1(opcode uint16) { //OR Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] |= c.v[y]
	c.v[0xF] = 0
}

func (c *Chip8) OP_8XY2(opcode uint16) { //AND Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] &= c.v[y]
	c.v[0xF] = 0
}

func (c *Chip8) OP_8XY3(opcode uint16) { //XOR Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] ^= c.v[y]
	c.v[0xF] = 0
}

func (c *Chip8) OP_8XY4(opcode uint16) { //ADD Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	sum := uint16(c.v[x]) + uint16(c.v[y])

	c.v[x] = byte(sum & 0xFF)

	if sum > 255 {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}
}

func (c *Chip8) OP_8XY5(opcode uint16) { //SUB Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	temp := c.v[x] >= c.v[y]

	c.v[x] -= c.v[y]

	if temp {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}
}

func (c *Chip8) OP_8XY6(opcode uint16) { //SHR Vx, {, Vy}
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	carryBit := c.v[x] & 0x1
	c.v[x] = c.v[y]
	c.v[x] >>= 1

	c.v[0xF] = carryBit
}

func (c *Chip8) OP_8XY7(opcode uint16) { //SUBN Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] = c.v[y] - c.v[x]

	if c.v[y] > c.v[x] {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}
}

func (c *Chip8) OP_8XYE(opcode uint16) { //SHL Vx, {, Vy}
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	carryBit := (c.v[x] & 0x80) >> 7

	c.v[x] = c.v[y]
	c.v[x] <<= 1

	c.v[0xF] = carryBit
}

func (c *Chip8) OP_9XY0(opcode uint16) { //SNE Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	if c.v[x] != c.v[y] {
		c.pc += 2
	}
}

func (c *Chip8) OP_ANNN(opcode uint16) { //LD I, addr
	imm := opcode & 0x0FFF

	c.index = imm
}

func (c *Chip8) OP_BNNN(opcode uint16) { //JP V0, addr
	imm := opcode & 0x0FFF

	c.pc = imm + uint16(c.v[0])
}

func (c *Chip8) OP_CXKK(opcode uint16) { //RND Vx, byte
	x := (opcode & 0x0F00) >> 8
	kk := byte(opcode & 0x00FF)

	c.v[x] = byte(rand.Intn(256)) & kk
}

func (c *Chip8) OP_DXYN(opcode uint16) { // DRW Vx, Vy, nibble
	x := c.v[(opcode&0x0F00)>>8] % byte(VIDEO_WIDTH)
	y := c.v[(opcode&0x00F0)>>4] % byte(VIDEO_HEIGHT)
	height := (opcode & 0x000F)

	c.v[0xF] = 0 // reset collision flag

	for j := uint16(0); j < height; j++ {
		spriteByte := c.memory[c.index+j]

		for i := uint16(0); i < 8; i++ {
			spritePixel := spriteByte & (0x80 >> i)

			if y+uint8(j) < 32 && x+uint8(i) < 64 {
				if spritePixel != 0 {
					if c.gfx[y+uint8(j)][x+uint8(i)] == 1 {
						c.v[0xF] = 1 // collision happened
					}
					c.gfx[y+uint8(j)][x+uint8(i)] ^= 1
				}
			}
		}
	}
	c.draw = true
}

func (c *Chip8) OP_EX9E(opcode uint16) { //SKP Vx
	x := (opcode & 0x0F00) >> 8

	if c.key[c.v[x]] > 0 {
		c.pc += 2
	}
}

func (c *Chip8) OP_EXA1(opcode uint16) { //SKNP Vx
	x := (opcode & 0x0F00) >> 8

	if c.key[c.v[x]] == 0 {
		c.pc += 2
	}
}

func (c *Chip8) OP_FX07(opcode uint16) { //LD Vx, DT
	x := (opcode & 0x0F00) >> 8

	c.v[x] = c.delayTimer
}

func (c *Chip8) OP_FX0A(opcode uint16) { //LD Vx, K
	x := (opcode & 0x0F00) >> 8

	keyPressed := false

	for i := 0; i < 16; i++ {
		if c.key[i] != 0 {
			c.v[x] = byte(i)
			keyPressed = true
			break
		}
	}

	if !keyPressed {
		c.pc -= 2 //redo the instruction until key pressed
	}
}

func (c *Chip8) OP_FX15(opcode uint16) { //LD DT, Vx
	x := (opcode & 0x0F00) >> 8

	c.delayTimer = c.v[x]
}

func (c *Chip8) OP_FX18(opcode uint16) { //LD ST, Vx
	x := (opcode & 0x0F00) >> 8

	c.soundTimer = c.v[x]
}

func (c *Chip8) OP_FX1E(opcode uint16) { //ADD I, Vx
	x := (opcode & 0x0F00) >> 8

	c.index += uint16(c.v[x])
}

func (c *Chip8) OP_FX29(opcode uint16) { //LD F, Vx
	x := (opcode & 0x0F00) >> 8

	c.index = FONTSET_START_ADDRESS + (5 * uint16(c.v[x]))
}

func (c *Chip8) OP_FX33(opcode uint16) { //LD B, Vx
	x := (opcode & 0x0F00) >> 8

	val := c.v[x]

	c.memory[c.index+2] = val % 10
	val /= 10

	c.memory[c.index+1] = val % 10
	val /= 10

	c.memory[c.index] = val % 10
}

func (c *Chip8) OP_FX55(opcode uint16) { //LD [I], Vx
	x := (opcode & 0x0F00) >> 8

	for i := uint16(0); i <= x; i++ {
		c.memory[c.index] = c.v[i]
		c.index++
	}

}

func (c *Chip8) OP_FX65(opcode uint16) { //LD Vx, [I]
	x := (opcode & 0x0F00) >> 8

	for i := uint16(0); i <= x; i++ {
		c.v[i] = c.memory[c.index]
		c.index++
	}
}

func (c *Chip8) OP_NULL(opcode uint16) {
	//no operation
}
