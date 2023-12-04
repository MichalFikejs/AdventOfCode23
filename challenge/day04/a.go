package day04

import (
	"fmt"
	"slices"
	"strings"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	// Initialize sum
	sum := 0

	for line := range challenge.Lines {
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], "|")

		winning := strings.Fields(parts[0])
		ours := strings.Fields(parts[1])

		var haveWinning int
		for _, n := range ours {
			if slices.Contains(winning, n) {
				haveWinning++
			}
		}

		if haveWinning >= 1 {
			sum += 1 << (haveWinning - 1)
		}
	}

	// Return the final sum
	return sum
}
