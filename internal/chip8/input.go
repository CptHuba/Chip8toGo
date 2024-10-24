package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

func ProcessInput(c *Chip8) bool {
	quit := false
	var event sdl.Event

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			quit = true

		case sdl.KEYDOWN:
			key := event.(*sdl.KeyboardEvent)
			switch key.Keysym.Sym {
			case sdl.K_ESCAPE:
				quit = true
			case sdl.K_x:
				c.key[0] = 1
			case sdl.K_1:
				c.key[1] = 1
			case sdl.K_2:
				c.key[2] = 1
			case sdl.K_3:
				c.key[3] = 1
			case sdl.K_q:
				c.key[4] = 1
			case sdl.K_w:
				c.key[5] = 1
			case sdl.K_e:
				c.key[6] = 1
			case sdl.K_a:
				c.key[7] = 1
			case sdl.K_s:
				c.key[8] = 1
			case sdl.K_d:
				c.key[9] = 1
			case sdl.K_z:
				c.key[0xA] = 1
			case sdl.K_c:
				c.key[0xB] = 1
			case sdl.K_4:
				c.key[0xC] = 1
			case sdl.K_r:
				c.key[0xD] = 1
			case sdl.K_f:
				c.key[0xE] = 1
			case sdl.K_v:
				c.key[0xF] = 1
			}

		case sdl.KEYUP:
			key := event.(*sdl.KeyboardEvent)
			switch key.Keysym.Sym {
			case sdl.K_x:
				c.key[0] = 0
			case sdl.K_1:
				c.key[1] = 0
			case sdl.K_2:
				c.key[2] = 0
			case sdl.K_3:
				c.key[3] = 0
			case sdl.K_q:
				c.key[4] = 0
			case sdl.K_w:
				c.key[5] = 0
			case sdl.K_e:
				c.key[6] = 0
			case sdl.K_a:
				c.key[7] = 0
			case sdl.K_s:
				c.key[8] = 0
			case sdl.K_d:
				c.key[9] = 0
			case sdl.K_z:
				c.key[0xA] = 0
			case sdl.K_c:
				c.key[0xB] = 0
			case sdl.K_4:
				c.key[0xC] = 0
			case sdl.K_r:
				c.key[0xD] = 0
			case sdl.K_f:
				c.key[0xE] = 0
			case sdl.K_v:
				c.key[0xF] = 0
			}
		}
	}

	return quit
}
