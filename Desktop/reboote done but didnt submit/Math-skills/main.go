package main

import (
	"bufio"
	"fmt"
	"math-skills/formulas"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	numbers, err := ReadNumbers(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if len(numbers) == 0 {	
		fmt.Println("No valid numbers found in the file.")
		os.Exit(1)
	}

	avg := formulas.Average(numbers)
	median := formulas.Median(numbers)
	variance := formulas.Variance(numbers)
	standardDeviation := formulas.StndardDeviation(numbers)

	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Variance: %.2f\n", variance)
	fmt.Printf("Standard Deviation: %.2f\n", standardDeviation)


}

func ReadNumbers(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err !=nil {
		return nil, err
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers, scanner.Err()
}