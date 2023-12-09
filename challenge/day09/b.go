package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	sum := 0

	for line := range challenge.Lines {
		numbers := strings.Fields(line)

		numberMatrix := make([][]int, 1)
		for _, number := range numbers {
			n, err := strconv.Atoi(number)

			if err != nil {
				panic(err)
			}
			numberMatrix[0] = append(numberMatrix[0], n)
		}

		i := 0
	NumbersFor:
		for {
			newLine := make([]int, 0)
			isAllZeros := true
			for n := 0; n < len(numberMatrix[i])-1; n++ {
				num := numberMatrix[i][n+1] - numberMatrix[i][n]
				newLine = append(newLine, num)
				if num != 0 {
					isAllZeros = false
				}
			}
			if isAllZeros {
				break NumbersFor
			}
			numberMatrix = append(numberMatrix, newLine)
			i++
		}

		for i := len(numberMatrix) - 1; i > 0; i-- {
			prev := numberMatrix[i-1]
			cur := numberMatrix[i]

			numberMatrix[i-1] = append([]int{prev[0] - cur[0]}, prev...)
		}
		sum += numberMatrix[0][0]
	}

	return sum
}
