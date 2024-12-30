package core

import (
	"fmt"
	"pixel-generator/src/entity"
)

func PixelGenerator(data entity.PixelArray) error {
	colorCodes := map[string]string{
		"GRAY":  "▒",
		"RED":   "\033[31m█\033[0m",
		"GREEN": "\033[32m█\033[0m",
		"BLUE":  "\033[34m█\033[0m",
		"WHITE": "\033[37m█\033[0m",
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if code, exists := colorCodes[data[i][j].Color]; exists {
				fmt.Print(code)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return nil
}

func DrawEllipse(data *entity.PixelArray, centerX, centerY, width, height int, color string) {
	a := width / 2
	b := height / 2

	for y := -b; y <= b; y++ {
		for x := -a; x <= a; x++ {
			if (x*x*b*b + y*y*a*a) <= (a * a * b * b) {
				if withinBounds(data, centerX+x, centerY+y) {
					(*data)[centerX+x][centerY+y] = entity.Pixel{Value: "", Color: color}
				}
			}
		}
	}
}

// Function to draw a rectangle
func DrawRectangle(data *entity.PixelArray, startX, startY, width, height int, color string) {
	// Define the four corners of the rectangle
	x1, y1 := startX, startY
	x2, y2 := startX+width, startY
	x3, y3 := startX, startY+height
	x4, y4 := startX+width, startY+height

	// Check bounds for all corners
	if withinBounds(data, x1, y1) && withinBounds(data, x2, y2) && withinBounds(data, x3, y3) && withinBounds(data, x4, y4) {
		// Draw the four edges of the rectangle
		drawLine(data, x1, y1, x2, y2, color) // Top edge
		drawLine(data, x1, y1, x3, y3, color) // Left edge
		drawLine(data, x2, y2, x4, y4, color) // Right edge
		drawLine(data, x3, y3, x4, y4, color) // Bottom edge
	} else {
		fmt.Println("Rectangle coordinates are out of bounds")
	}
}

// Function to draw a line
func drawLine(data *entity.PixelArray, x1, y1, x2, y2 int, color string) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := -1
	sy := -1

	if x1 < x2 {
		sx = 1
	}
	if y1 < y2 {
		sy = 1
	}

	err := dx - dy

	for {
		if withinBounds(data, x1, y1) {
			(*data)[x1][y1] = entity.Pixel{Value: "", Color: color}
		}
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func withinBounds(data *entity.PixelArray, x, y int) bool {
	if x < 0 || y < 0 || x >= len(*data) || y >= len((*data)[0]) {
		return false
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
