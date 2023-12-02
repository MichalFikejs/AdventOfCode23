package day02

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
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return findIt(challenge)
}

const (
	totalRed   = 12
	totalGreen = 13
	totalBlue  = 14
)

func findIt(chall *challenge.Input) int {
	// Initialize sum
	sum := 0

	for line := range chall.Lines {
		id, isOk := processGame(line)

		if isOk {
			sum += id
		}
	}

	// Return the final sum
	return sum
}

func processGame(game string) (int, bool) {
	isValid := true

	// Split of game string by :
	gameSplit := strings.Split(game, ":")
	idString := strings.Split(gameSplit[0], " ")[1]
	idInt, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println("Error converting ID to int:", err)
		return 0, false
	}

	draws := strings.Split(gameSplit[1], ";")
	for _, draw := range draws {
		balls := strings.Split(draw, ",")
		for _, ball := range balls {
			ballSplit := strings.Split(strings.TrimSpace(ball), " ")
			ballCount, err := strconv.Atoi(ballSplit[0])
			ballColor := ballSplit[1]

			if err != nil {
				fmt.Println("Error converting red to int:", err)
				return 0, false
			}

			switch ballColor {
			case "red":
				isValid = isValid && ballCount <= totalRed
			case "green":
				isValid = isValid && ballCount <= totalGreen
			case "blue":
				isValid = isValid && ballCount <= totalBlue
			}
		}
	}

	return idInt, isValid
}
