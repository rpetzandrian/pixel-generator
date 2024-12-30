package main

import (
	"fmt"
	"os"
	"pixel-generator/src/core"
	"pixel-generator/src/entity"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 5 {
		fmt.Println("[+] COLORS [+]")
		fmt.Println("[1]. GRAY")
		fmt.Println("[2]. RED")
		fmt.Println("[3]. GREEN")
		fmt.Println("[4]. BLUE")
		fmt.Println("[5]. WHITE")
		fmt.Println("Usage: go run . <width> <height> <color> <ellipse/rectangle>")
		return
	}

	widthStr := args[1]
	heightStr := args[2]

	// Convert width and height to integers
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		fmt.Printf("Invalid width: %s", widthStr)
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		fmt.Printf("Invalid height: %s", heightStr)
		return
	}

	if os.Args[3] != "GRAY" && os.Args[3] != "RED" && os.Args[3] != "GREEN" && os.Args[3] != "BLUE" && os.Args[3] != "WHITE" {
		fmt.Println("Error: Invalid color")
		return
	}

	if os.Args[4] != entity.ELLIPSE && os.Args[4] != entity.RECTANGLE {
		fmt.Println("Error: Invalid mode argument. Please provide a valid mode (<ellipse/rectangle>).")
		return
	}

	if os.Args[4] == entity.ELLIPSE {
		data := make(entity.PixelArray, width+10)
		for i := range data {
			data[i] = make([]entity.Pixel, height)
		}
		x := 5
		y := 5
		if width > 10 {
			x = width / 2
		}
		if height > 10 {
			y = height / 2
		}
		core.DrawEllipse(&data, x, y, width, height, os.Args[3])
		core.PixelGenerator(data)
	} else if os.Args[4] == entity.RECTANGLE {
		data := make(entity.PixelArray, width+10)
		for i := range data {
			data[i] = make([]entity.Pixel, height+10)
		}
		core.DrawRectangle(&data, 0, 0, width, height, os.Args[3])
		core.PixelGenerator(data)
	}
}
