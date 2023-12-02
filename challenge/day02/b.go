package day02

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
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	// Initialize sum
	sum := 0

	for line := range challenge.Lines {
		sum += processGameB(line)
	}

	// Return the final sum
	return sum
}

func processGameB(game string) int {
	var maxBlue, maxGreen, maxRed int

	// Split of game string by :
	gameSplit := strings.Split(game, ":")

	draws := strings.Split(gameSplit[1], ";")
	for _, draw := range draws {
		balls := strings.Split(draw, ",")
		for _, ball := range balls {
			ballSplit := strings.Split(strings.TrimSpace(ball), " ")
			ballCount, err := strconv.Atoi(ballSplit[0])
			ballColor := ballSplit[1]

			if err != nil {
				fmt.Println("Error converting red to int:", err)
				return 0
			}

			switch ballColor {
			case "red":
				if ballCount > maxRed {
					maxRed = ballCount
				}
			case "green":
				if ballCount > maxGreen {
					maxGreen = ballCount
				}
			case "blue":
				if ballCount > maxBlue {
					maxBlue = ballCount
				}
			}
		}
	}

	return maxBlue * maxGreen * maxRed
}
