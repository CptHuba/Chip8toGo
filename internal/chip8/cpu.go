package chip8

import (
	"math/rand"
)

func (c *Chip8) FetchOpcode() uint16 {
	return uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
}

func (c *Chip8) OP_00E0() { //CLS
	for i := range c.gfx {
		c.gfx[i] = 0
	}
}

func (c *Chip8) OP_00EE() { //RET
	c.pc = c.stack[c.sp]
	c.sp--
}

func (c *Chip8) OP_1NNN(opcode uint16) { //JP addr
	c.pc = opcode & 0x0FFF
}

func (c *Chip8) OP_2NNN(opcode uint16) { //CALL addr
	c.sp++
	c.stack[c.sp] = c.pc
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
}

func (c *Chip8) OP_8XY2(opcode uint16) { //AND Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] &= c.v[y]
}

func (c *Chip8) OP_8XY3(opcode uint16) { //XOR Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	c.v[x] ^= c.v[y]
}

func (c *Chip8) OP_8XY4(opcode uint16) { //ADD Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	sum := uint16(c.v[x]) + uint16(c.v[y])

	if sum > 255 {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}

	c.v[x] = byte(sum & 0xFF)
}

func (c *Chip8) OP_8XY5(opcode uint16) { //SUB Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	if c.v[x] > c.v[y] {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}

	c.v[x] -= c.v[y]
}

func (c *Chip8) OP_8XY6(opcode uint16) { //SHR Vx, {, Vy}
	x := (opcode & 0x0F00) >> 8

	c.v[0xF] = c.v[x] & 0x1

	c.v[x] >>= 1
}

func (c *Chip8) OP_8XY7(opcode uint16) { //SUBN Vx, Vy
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4

	if c.v[y] > c.v[x] {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}

	c.v[x] = c.v[y] - c.v[x]
}

func (c *Chip8) OP_8XYE(opcode uint16) { //SHL Vx, {, Vy}
	x := (opcode & 0x0F00) >> 8

	c.v[0xF] = c.v[x] & 0x8

	c.v[x] <<= 1
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

func (c *Chip8) OP_DXYN(opcode uint16) { //DRW Vx, Vy, nibble
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	height := (opcode & 0x000F)

	c.v[0xF] = 0 //Reset collision flag

	for row := uint16(0); row < height; row++ {
		spriteByte := c.memory[c.index+row]

		for col := uint16(0); col < 8; col++ {
			spritePixel := spriteByte & (0x80 >> byte(col))

			posX := (uint16(x) + col) % VIDEO_WIDTH
			posY := (uint16(y) + row) % VIDEO_HEIGHT

			screenIndex := posY*VIDEO_WIDTH + posX
			screenPixel := c.gfx[screenIndex]

			if spritePixel == 1 && screenPixel == 1 {
				c.v[0xF] = 1 //Collision happened
			}

			c.gfx[screenIndex] ^= spritePixel
		}
	}
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
		c.pc -= 2 //Redo the instruction until key pressed
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
		c.memory[c.index+i] = c.v[i]
	}

}

func (c *Chip8) OP_FX65(opcode uint16) { //LD Vx, [I]
	x := (opcode & 0x0F00) >> 8

	for i := uint16(0); i <= x; i++ {
		c.v[i] = c.memory[c.index+i]
	}
}
