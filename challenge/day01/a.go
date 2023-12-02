package day01

import (
	"fmt"
	"unicode"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return findIt(challenge)
}

func findIt(chall *challenge.Input) int {
	// Initialize sum
	sum := 0

	for line := range chall.Lines {

		// Find the first and last digits in the line
		var firstDigit, lastDigit int
		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == 0 {
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0')
			}
		}

		// Combine the digits to form a two-digit number
		calibrationValue := firstDigit*10 + lastDigit

		// Add the calibration value to the sum
		sum += calibrationValue
	}

	// Return the final sum
	return sum
}
