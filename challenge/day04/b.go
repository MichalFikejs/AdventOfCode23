package day04

import (
	"fmt"
	"slices"
	"strings"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	// Initialize sum
	sum := 0

	cards := make([]string, 0)
	for card := range challenge.Lines {
		cards = append(cards, card)
	}

	countsOfCards := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		countsOfCards[i] = 1
	}

	for i, card := range cards {
		parts := strings.Split(card, ":")
		parts = strings.Split(parts[1], "|")

		winning := strings.Fields(parts[0])
		ours := strings.Fields(parts[1])

		var haveWinning int
		for _, n := range ours {
			if slices.Contains(winning, n) {
				haveWinning++
			}
		}

		for io := 1; io <= haveWinning; io++ {
			countsOfCards[i+io] += countsOfCards[i]
		}

		sum += countsOfCards[i]
	}

	// Return the final sum
	return sum
}
