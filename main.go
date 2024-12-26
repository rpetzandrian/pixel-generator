package main

import (
	"fmt"
	"os"
	"pixel-generator/src/core"
	"pixel-generator/src/entity"
	"pixel-generator/src/helpers"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <size> <recursive/loop>")
		return
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Invalid size argument. Please provide a valid integer.")
		return
	}

	if os.Args[2] != entity.RECURSIVE && os.Args[2] != entity.LOOP {
		fmt.Println("Error: Invalid mode argument. Please provide a valid mode (recursive/loop).")
		return
	}

	arr := helpers.GeneratePixelArray(size)

	if os.Args[2] == entity.RECURSIVE {
		core.PixelGeneratorRecursive(arr, 0, 0)
	} else {
		core.PixelGenerator(arr)
	}
}
