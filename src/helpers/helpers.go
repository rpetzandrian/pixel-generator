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
		for j := 0; j < size; j += 4 {
			val := r.Intn(256)
			row[j] = val & 3
			if j+1 < size {
				row[j+1] = (val >> 2) & 3
			}
			if j+2 < size {
				row[j+2] = (val >> 4) & 3
			}
			if j+3 < size {
				row[j+3] = (val >> 6) & 3
			}
		}
		pixels[i] = row
	}

	return pixels
}
