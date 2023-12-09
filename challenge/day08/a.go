package day08

import (
	"fmt"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 8, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type Instructions struct {
	Left  string
	Right string
}

func partA(challenge *challenge.Input) int {
	sum := 0

	var instructions string
	allInstructions := make(map[string]Instructions)

	currentPosition := "AAA"
	for line := range challenge.Lines {
		if instructions == "" {
			instructions = line
		} else if line == "" {
			continue
		} else {
			current, left, right := parseLine(line)
			allInstructions[current] = Instructions{left, right}
		}
	}

	for {
		for _, instruction := range instructions {
			sum += 1
			switch string(instruction) {
			case "L":
				currentPosition = allInstructions[currentPosition].Left
			case "R":
				currentPosition = allInstructions[currentPosition].Right
			}
			if currentPosition == "ZZZ" {
				goto Exit
			}
		}
	}

Exit:
	return sum
}

func parseLine(line string) (string, string, string) {
	var current, left, right string
	current = line[:3]
	left = line[7:10]
	right = line[12:15]
	return current, left, right
}
