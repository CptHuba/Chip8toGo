
# Chip8toGo

Chip8toGo is a simple emulator for the CHIP-8 programming language, implemented in Go. This project is just for fun and serves as an educational practice for me.

## Features

- **CHIP-8 Instruction Set**: Supports the complete instruction set of CHIP-8.
- **ROM Support**: Load and run CHIP-8 ROMs.
- **Basic Graphics**: Simple display to visualize the output of executed ROMs.
- **Sound Support**: NOT SUPPORTED.

## Getting Started

### Prerequisites

- Go (version 1.23.2 or later)
- SDL2 (make sure to install the Go SDL2 bindings)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/CptHuba/Chip8toGo.git
cd Chip8toGo
```

2. Install SDL2 development libraries:
```bash
sudo apt-get install libsdl2-dev
```

3. Install Go dependencies:
```bash
go mod tidy
```

### Usage

To run the emulator, use the following command:

```bash
go run . <Scale> <ROM>
```

- `<Scale>`: An integer to define the scaling of the display.
- `<ROM>`: The path to the CHIP-8 ROM file you want to run.

**Example**:

```bash
go run . 15 tetris.ch8
```