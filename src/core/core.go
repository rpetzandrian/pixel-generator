package core

import (
	"bytes"
	"fmt"
	"pixel-generator/src/entity"
)

var pixelChars = []string{" ", "░", "▒", "▓"}

func PixelGenerator(data entity.PixelArray) error {
	var buffer bytes.Buffer
	size := len(data)
	for i := 0; i < size; i++ {
		row := data[i]
		for j := 0; j < size; j++ {
			buffer.WriteString(pixelChars[row[j]])
		}
		buffer.WriteString("\n")
	}
	fmt.Print(buffer.String())
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

	fmt.Print(pixelChars[data[i][j]])
	PixelGeneratorRecursive(data, i, j+1)
}
