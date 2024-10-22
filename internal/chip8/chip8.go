package chip8

type Chip8 struct {
	memory     [4096]byte    //4K memory
	v          [16]byte      //16 registers (V0-VF)
	index      uint16        //index register
	pc         uint16        //program counter
	stack      [16]uint16    //16-level stack for pc
	sp         byte          //stack pointer
	delayTimer byte          //delay timer
	soundTimer byte          //sound timer
	key        [16]byte      //input key state
	gfx        [64 * 32]byte //display (64x32)
}
