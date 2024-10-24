package chip8

import (
	"log"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
}

func NewDisplay(title string, windowWidth, windowHeight, textureWidth, textureHeight int32) *Display {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("Unable to initialize SDL: %v", err)
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Could not create window: %v", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Could not create renderer: %v", err)
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, textureWidth, textureHeight)
	if err != nil {
		log.Fatalf("Could not create texture: %v", err)
	}

	return &Display{
		window:   window,
		renderer: renderer,
		texture:  texture,
	}
}

func (d *Display) Update(buffer unsafe.Pointer, pitch int) {
	d.renderer.Clear()
	d.texture.Update(nil, buffer, pitch)
	d.renderer.Copy(d.texture, nil, nil)
	d.renderer.Present()
}

func (d *Display) Cleanup() {
	d.texture.Destroy()
	d.renderer.Destroy()
	d.window.Destroy()
	sdl.Quit()
}
