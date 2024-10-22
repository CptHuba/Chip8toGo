package utils

import (
	"Chip8toGo/internal/chip8"
	"os"
)

func loadROMFile(chip8 *chip8.Chip8, filename string) error {
	rom, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	chip8.LoadROM(rom)
	return nil
}
