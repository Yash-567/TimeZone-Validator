package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFileName  = "input.txt"
	outputFileName = "output.txt"
)

func main() {
	// Open input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Open output file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Initialize map to store unique date-time values
	uniqueDateTimes := make(map[string]bool)

	// Read input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		dateTime := strings.TrimSpace(scanner.Text())

		// Validate date-time format using custom validation function
		validFormat := validateDateTime(dateTime)
		if validFormat {
			// Normalize valid date-time values
			normalizedDateTime, err := normalizeDateTime(dateTime)
			if err == nil {
				// Store unique normalized date-time values
				uniqueDateTimes[normalizedDateTime] = true
			} else {
				fmt.Printf("Error normalizing date-time '%s': %v\n", dateTime, err)
			}
		} else {
			fmt.Printf("Invalid date-time format: %s\n", dateTime)
		}
	}

	// Write unique normalized date-time values to output file
	for dateTime := range uniqueDateTimes {
		fmt.Fprintln(outputFile, dateTime)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input file:", err)
		return
	}

	fmt.Println("Saved all Unique, Validated and Normalized DateTimes to output file")
}
