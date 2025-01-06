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

	if os.Args[2] != entity.RECURSIVE && os.Args[2] != entity.LOOP {
		fmt.Println("Error: Invalid mode argument. Please provide a valid mode (recursive/loop).")
		return
	}

	// Path to the existing CSV file
	csvFilePath := "result.csv"

	// Open the existing CSV file in append mode
	file, err := os.OpenFile(csvFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	arr := helpers.GeneratePixelArray(size)

	var data []string
	if os.Args[2] == entity.RECURSIVE {
		start := time.Now()
		core.PixelGeneratorRecursive(arr, 0, 0)
		elapsed := time.Since(start)

		// Prepare the data to write
		data = append(data, strconv.Itoa(size), "recursive", elapsed.String())

	} else {
		start := time.Now()
		core.PixelGenerator(arr)
		elapsed := time.Since(start)

		// Prepare the data to write
		data = append(data, strconv.Itoa(size), "looping", elapsed.String())
	}

	// Write the new row to the existing CSV file
	if err := writer.Write(data); err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}

	fmt.Println("New row added successfully.")
}
