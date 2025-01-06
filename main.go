package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"pixel-generator/src/core"
	"pixel-generator/src/entity"
	"pixel-generator/src/helpers"
	"strconv"
	"time"
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

	mode := os.Args[2]
	if mode != entity.RECURSIVE && mode != entity.LOOP {
		fmt.Println("Error: Invalid mode argument. Please provide a valid mode (recursive/loop).")
		return
	}

	csvFilePath := "result.csv"
	file, err := os.OpenFile(csvFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	arr := helpers.GeneratePixelArray(size)

	start := time.Now()
	if mode == entity.RECURSIVE {
		core.PixelGeneratorRecursive(arr, 0, 0)
	} else {
		core.PixelGenerator(arr)
	}
	elapsed := time.Since(start)

	data := []string{strconv.Itoa(size), mode, elapsed.String()}
	if err := writer.Write(data); err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}

	fmt.Println("New row added successfully.")
}
