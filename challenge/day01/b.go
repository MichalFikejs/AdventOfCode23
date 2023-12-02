package day01

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	return findItB(challenge)
}

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findItB(chall *challenge.Input) int {
	// Initialize sum
	sum := 0
	numerOfLines := 0

	for line := range chall.Lines {
		numerOfLines++
		// Get the current line

		// Find the first and last digits in the line
		var firstDigit, lastDigit, firstPosition, lastPosition int
		firstPosition, lastPosition = -1, -1
		var wordString string
		for i, char := range line {
			if unicode.IsDigit(char) {
				if wordString != "" {
					for key := range digitMap {
						in := strings.Index(wordString, key)
						if in != -1 {
							if in < firstPosition || firstPosition == -1 {
								firstPosition = in
								firstDigit = digitMap[key]
							}
						}
						// Your code here
						// Use the key variable to access each key in the digitMap
					}
					wordString = ""
				}
				if firstDigit == 0 {
					firstPosition = i
					firstDigit = int(char - '0')
				}
				lastDigit = int(char - '0')
				lastPosition = i
			} else {
				wordString += string(char)
			}
		}
		if wordString != "" {
			wholeWordStringPosition := strings.Index(line, wordString)
			for key := range digitMap {
				in := strings.Index(wordString, key)
				if in != -1 {
					in += wholeWordStringPosition
					if in < firstPosition || firstPosition == -1 {
						firstPosition = in
						firstDigit = digitMap[key]
					}
					if in > lastPosition {
						lastPosition = in
						lastDigit = digitMap[key]
					}
				}
				// Your code here
				// Use the key variable to access each key in the digitMap
			}
			wordString = ""

		}

		// Combine the digits to form a two-digit number
		calibrationValue := firstDigit*10 + lastDigit

		// Add the calibration value to the sum
		sum += calibrationValue
	}

	// Return the final sum
	return sum
}
