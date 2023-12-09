package day08

import (
	"fmt"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	sum := 1

	var instructions string
	allInstructions := make(map[string]Instructions)

	currentPositions := make([]string, 0)
	for line := range challenge.Lines {
		if instructions == "" {
			instructions = line
		} else if line == "" {
			continue
		} else {
			current, left, right := parseLine(line)
			allInstructions[current] = Instructions{left, right}
			if current[2] == 'A' {
				currentPositions = append(currentPositions, current)
			}
		}
	}

	for i, _ := range currentPositions {
		cnt := 0
	InstructionsFor:
		for {
			for _, instruction := range instructions {
				cnt++

				switch string(instruction) {
				case "L":
					currentPositions[i] = allInstructions[currentPositions[i]].Left
				case "R":
					currentPositions[i] = allInstructions[currentPositions[i]].Right
				}
				if currentPositions[i][2] == 'Z' {
					break InstructionsFor
				}
			}
		}
		sum = lcm(sum, cnt)
	}

	return sum
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
