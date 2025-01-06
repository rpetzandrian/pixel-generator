package helpers

import (
	"math/rand"
	"pixel-generator/src/entity"
	"time"
)

func GeneratePixelArray(size int) entity.PixelArray {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pixels := make(entity.PixelArray, size)

	for i := 0; i < size; i++ {
		row := make([]int, size)
		for j := 0; j < size; j++ {
			row[j] = r.Intn(4)
		}
		pixels[i] = row
	}

	return pixels
}
