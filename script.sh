#!/bin/bash  
  
# Define the CSV file path  
CSV_FILE="result.csv"  
  
# Check if the CSV file exists, if not create it with a header  
if [ ! -f "$CSV_FILE" ]; then  
    echo "Creating CSV file with header..."  
    echo "size,algoritma,execution-time" > "$CSV_FILE"  
fi  
  
# Define the sizes and algorithms  
sizes=(10 100 1000)  
algorithms=("loop" "recursive")  
  
# Loop through sizes and algorithms  
for size in "${sizes[@]}"; do  
    for algo in "${algorithms[@]}"; do  
        # Run the Go program with the specified size and algorithm  
        go run main.go "$size" "$algo"  
    done  
done  
