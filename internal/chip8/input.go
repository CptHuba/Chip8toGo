package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

func ProcessInput(keys []uint8) bool {
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
				keys[0] = 1
			case sdl.K_1:
				keys[1] = 1
			case sdl.K_2:
				keys[2] = 1
			case sdl.K_3:
				keys[3] = 1
			case sdl.K_q:
				keys[4] = 1
			case sdl.K_w:
				keys[5] = 1
			case sdl.K_e:
				keys[6] = 1
			case sdl.K_a:
				keys[7] = 1
			case sdl.K_s:
				keys[8] = 1
			case sdl.K_d:
				keys[9] = 1
			case sdl.K_z:
				keys[0xA] = 1
			case sdl.K_c:
				keys[0xB] = 1
			case sdl.K_4:
				keys[0xC] = 1
			case sdl.K_r:
				keys[0xD] = 1
			case sdl.K_f:
				keys[0xE] = 1
			case sdl.K_v:
				keys[0xF] = 1
			}

		case sdl.KEYUP:
			key := event.(*sdl.KeyboardEvent)
			switch key.Keysym.Sym {
			case sdl.K_x:
				keys[0] = 0
			case sdl.K_1:
				keys[1] = 0
			case sdl.K_2:
				keys[2] = 0
			case sdl.K_3:
				keys[3] = 0
			case sdl.K_q:
				keys[4] = 0
			case sdl.K_w:
				keys[5] = 0
			case sdl.K_e:
				keys[6] = 0
			case sdl.K_a:
				keys[7] = 0
			case sdl.K_s:
				keys[8] = 0
			case sdl.K_d:
				keys[9] = 0
			case sdl.K_z:
				keys[0xA] = 0
			case sdl.K_c:
				keys[0xB] = 0
			case sdl.K_4:
				keys[0xC] = 0
			case sdl.K_r:
				keys[0xD] = 0
			case sdl.K_f:
				keys[0xE] = 0
			case sdl.K_v:
				keys[0xF] = 0
			}
		}
	}

	return quit
}
