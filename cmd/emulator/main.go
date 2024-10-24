package main

import (
	"fmt"
	"os"
	"strconv"

	"Chip8toGo/internal/chip8"
	"Chip8toGo/internal/utils"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <Scale> <ROM>\n", os.Args[0])
		os.Exit(1)
	}

	videoScale, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid scale:", os.Args[1])
		os.Exit(1)
	}

	romFilename := os.Args[2]

	cpu := chip8.NewChip8()

	if err := utils.LoadROMFile(cpu, romFilename); err != nil {
		panic(err)
	}

	display, err := chip8.NewDisplay("Chip8 - Emulator", chip8.VIDEO_WIDTH, chip8.VIDEO_HEIGHT, int32(videoScale))
	if err != nil {
		panic(err)
	}
	defer display.CleanUp()

	quit := false

	for !quit {
		quit = chip8.ProcessInput(cpu)
		cpu.Cycle()

		display.Update(cpu)
	}

}
