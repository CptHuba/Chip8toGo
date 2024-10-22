package chip8

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
	x := (opcode >> 8) & 0x000F
	kk := byte(opcode & 0x000F)

	if kk == c.v[x] {
		c.pc += 2
	}
}

func (c *Chip8) OP_4XKK(opcode uint16) { //SNE Vx, byte
	x := (opcode >> 8) & 0x000F
	kk := byte(opcode & 0x000F)

	if kk != c.v[x] {
		c.pc += 2
	}
}

func (c *Chip8) OP_5XY0(opcode uint16) { //SE Vx, Vy
	x := (opcode >> 8) & 0x000F
	y := (opcode >> 4) & 0x000F

	if c.v[x] == c.v[y] {
		c.pc += 2
	}
}
