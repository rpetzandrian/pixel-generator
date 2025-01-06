package core

import (
	"fmt"
	"pixel-generator/src/entity"
)

var pixelChars = []string{" ", "░", "▒", "▓"}

func PixelGenerator(data entity.PixelArray) error {
	size := len(data)
	for i := 0; i < size; i++ {
		row := data[i]
		for j := 0; j < size; j++ {
			print(pixelChars[row[j]])
		}
		fmt.Println()
	}

	return nil
}

// PixelGeneratorRecursive
func PixelGeneratorRecursive(data entity.PixelArray, i, j int) {
	size := len(data)
	if i == size {
		return
	}

	if j == size {
		fmt.Println()
		PixelGeneratorRecursive(data, i+1, 0)
		return
	}

	print(pixelChars[data[i][j]])

	PixelGeneratorRecursive(data, i, j+1)
}
