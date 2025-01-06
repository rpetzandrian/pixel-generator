package core

import (
	"fmt"
	"pixel-generator/src/entity"
)

var pixelChars = []string{" ", "░", "▒", "▓"}

func PixelGenerator(data entity.PixelArray) error {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			print(pixelChars[data[i][j]])
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

	print(pixelChars[data[i][j]])

	PixelGeneratorRecursive(data, i, j+1)
}
