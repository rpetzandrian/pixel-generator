package core

import (
	"fmt"
	"pixel-generator/src/entity"
)

func PixelGenerator(data entity.PixelArray) error {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			switch data[i][j] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("░")
			case 2:
				fmt.Print("▒")
			case 3:
				fmt.Print("▓")
			}
		}
		fmt.Println()
	}

	return nil
}

// PixelGeneratorRecursive
func PixelGeneratorRecursive(data entity.PixelArray, i, j int) {
	if i == len(data) {
		return
	}

	if j == len(data[i]) {
		fmt.Println()
		PixelGeneratorRecursive(data, i+1, 0)
		return
	}

	switch data[i][j] {
	case 0:
		fmt.Print(" ")
	case 1:
		fmt.Print("░")
	case 2:
		fmt.Print("▒")
	case 3:
		fmt.Print("▓")
	}

	PixelGeneratorRecursive(data, i, j+1)
}
