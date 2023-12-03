package day03

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func isSymbol(runeToCheck rune) bool {
	return !unicode.IsNumber(runeToCheck) && string(runeToCheck) != "."
}

func partA(challenge *challenge.Input) int {
	// Initialize sum
	sum := 0

	matrix := make([][]rune, 0)

	for line := range challenge.Lines {
		matrix = append(matrix, []rune(line))
	}

	for line := range matrix {
		for i := 0; i < len(matrix[line]); i++ {
			// loop body
			if unicode.IsDigit(matrix[line][i]) {
				minPosition, maxPosition := i, i
				number := string(matrix[line][i])
				for i+1 < len(matrix[line]) && unicode.IsDigit(matrix[line][i+1]) {
					number += string(matrix[line][i+1])
					i++
					maxPosition = i
				}

				minNumberTocheck := minPosition
				if minPosition > 0 {
					minNumberTocheck--
				}

				maxNumberTocheck := maxPosition
				if maxPosition < len(matrix[line])-1 {
					maxNumberTocheck++
				}

				isValid := false
				for i := -1; i <= 1; i++ {
					lineToCheck := line + i
					if lineToCheck >= 0 && lineToCheck < len(matrix) {
						for symbolToCheck := minNumberTocheck; symbolToCheck <= maxNumberTocheck; symbolToCheck++ {
							if isSymbol(matrix[lineToCheck][symbolToCheck]) {
								isValid = true
							}
						}
					}
				}

				if isValid {
					numberInt, err := strconv.Atoi(number)

					if err != nil {
						fmt.Println("Error converting to int:", err)
						break
					}

					sum += int(numberInt)
				}
			}
		}
	}

	// Return the final sum
	return sum
}
