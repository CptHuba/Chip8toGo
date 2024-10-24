package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Display struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	videoScale int32
}

func NewDisplay(title string, width, height, videoScale int32) (*Display, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width*videoScale, height*videoScale, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		window.Destroy()
		return nil, err
	}

	return &Display{
		window:     window,
		renderer:   renderer,
		videoScale: videoScale,
	}, nil
}

func (d *Display) Update(cpu *Chip8) {
	if cpu.Draw() {
		vector := cpu.Display()
		d.renderer.Clear()
		for j := 0; j < 32; j++ {
			for i := 0; i < 64; i++ {
				if vector[j][i] != 0 {
					d.renderer.SetDrawColor(255, 255, 255, 255)
				} else {
					d.renderer.SetDrawColor(0, 0, 0, 255)
				}
				d.renderer.FillRect(&sdl.Rect{
					X: int32(i) * d.videoScale,
					Y: int32(j) * d.videoScale,
					W: d.videoScale,
					H: d.videoScale,
				})
			}
		}
		d.renderer.Present()
	}
}

func (d *Display) CleanUp() {
	d.renderer.Destroy()
	d.window.Destroy()
	sdl.Quit()
}
