package day06

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
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type Race struct {
	Time     int
	Distance int
}

func partA(challenge *challenge.Input) int {
	times := make([]int, 0)
	distances := make([]int, 0)

	for line := range challenge.Lines {
		if strings.HasPrefix(line, "Time: ") {
			trimmed := strings.TrimPrefix(line, "Time: ")
			splitted := strings.Split(trimmed, " ")

			for _, timeString := range splitted {
				time, err := strconv.Atoi(timeString)
				if err == nil {
					times = append(times, time)
				}
			}
		}
		if strings.HasPrefix(line, "Distance: ") {
			trimmed := strings.TrimPrefix(line, "Distance: ")
			splitted := strings.Split(trimmed, " ")
			for _, distanceString := range splitted {
				distance, err := strconv.Atoi(distanceString)
				if err == nil {
					distances = append(distances, distance)
				}
			}
		}
	}

	sum := 1

	for i, time := range times {
		race := Race{
			Time:     time,
			Distance: distances[i],
		}

		raceWinningVariants := 0
		for i := 0; i <= race.Time; i++ {
			if i*(race.Time-i) > race.Distance {
				raceWinningVariants++
			}
		}

		sum *= raceWinningVariants
	}

	return sum
}
