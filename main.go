package main

import (
	"fmt"
	"os"
)

// size of the board, always a square, ex: 4 means 4x4
const SIZE = 4

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		if (arg == "-p") || (arg == "--play") {
			playGame()
			return
		}
	}

	fmt.Println("Starting onnainai autoplay.")
    fmt.Println("If you wanted to play 2048 in the cli, use the -p option.")
    fmt.Println("")
}
