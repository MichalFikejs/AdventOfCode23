package day06

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
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	time := 0
	var distance int

	for line := range challenge.Lines {
		if strings.HasPrefix(line, "Time: ") {
			trimmed := strings.TrimPrefix(line, "Time: ")
			replaced := strings.ReplaceAll(trimmed, " ", "")

			time, _ = strconv.Atoi(replaced)

		}
		if strings.HasPrefix(line, "Distance: ") {
			trimmed := strings.TrimPrefix(line, "Distance: ")
			replaced := strings.ReplaceAll(trimmed, " ", "")

			distance, _ = strconv.Atoi(replaced)
		}
	}

	raceWinningVariants := 0
	for i := 0; i <= time; i++ {
		if i*(time-i) > distance {
			raceWinningVariants++
		}
	}

	return raceWinningVariants
}
