package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"git.com/3Sayed/Go/utils"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("error")
		os.Exit(0)
	}

	// Handle input, Open file, show error if exist, enforce to close the file.
	input, err := os.Open(arg[0])
	if err != nil {
		log.Fatal("Error to open the input file", err)
	}
	defer input.Close()

	// // Handle Output, Open file, show error if exist, enforce to close the file.
	output, err := os.Create(arg[1])
	if err != nil {
		log.Fatal("Error: in generate the input file", err)
	}
	// defer output.Close()
	defer output.Close()
	// To Read and Write the both files

	//Scanner to Read sample.text
	scanner := bufio.NewScanner(input)
	//Writer to write the output
	writer := bufio.NewWriter(output)

	//Read Line by Line
	for scanner.Scan() {
		txt := scanner.Text() // Read the current line
		
		// Process the text
		processed := utils.ProcessText(txt)

		// write line to buffer
		_, err := writer.WriteString(processed + "\n")

		// Handle error
		if err != nil {
			log.Fatal("Error in writing the output", err)
		}
	}

	// flush any buffered data to file
	writer.Flush()

	// Handle scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal("Error in reading the input file.", err)
	}

	// Log to print DateTime and Message
	log.Println("Successfully wrote the wrote to the Result.txt")
}
