package day03

import (
	"fmt"
	"strconv"
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
	// Initialize sum

	matrix := make([][]rune, 0)

	for line := range challenge.Lines {
		matrix = append(matrix, []rune(line))
	}

	gearsWithNumbers := make(map[string][]int)

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

				for i := -1; i <= 1; i++ {
					lineToCheck := line + i
					if lineToCheck >= 0 && lineToCheck < len(matrix) {
						for symbolToCheck := minNumberTocheck; symbolToCheck <= maxNumberTocheck; symbolToCheck++ {
							if string(matrix[lineToCheck][symbolToCheck]) == "*" {
								numberInt, err := strconv.Atoi(number)

								if err != nil {
									fmt.Println("Error converting to int:", err)
									break
								}
								gearsWithNumbers[fmt.Sprintf("%dx%d", lineToCheck, symbolToCheck)] = append(gearsWithNumbers[fmt.Sprintf("%dx%d", lineToCheck, symbolToCheck)], numberInt)

							}

						}
					}
				}
			}
		}
	}

	sum := 0

	for _, value := range gearsWithNumbers {
		if len(value) >= 2 {
			addOfSymbolNumbers := 1
			for _, number := range value {
				addOfSymbolNumbers *= number
			}
			sum += addOfSymbolNumbers
		}
	}

	// Return the final sum
	return sum
}
