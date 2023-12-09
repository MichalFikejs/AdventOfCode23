package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
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
			numberMatrix[i-1] = append(prev, prev[len(prev)-1]+cur[len(cur)-1])
		}
		sum += numberMatrix[0][len(numberMatrix[0])-1]
	}

	return sum
}
