package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func isStringValidNumericFormat(s string) bool {
	var a, b, c int
	n, _ := fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
	return n == 3
}

type mappingStruct struct {
	Min    int
	Max    int
	NewMin int
}

func partA(challenge *challenge.Input) int {
	seeds := make([]int, 0)

	mappings := make([][]string, 0)
	currentMapping := make([]string, 0)

	for line := range challenge.Lines {
		if strings.HasPrefix(line, "seeds: ") {
			trimmed := strings.TrimPrefix(line, "seeds: ")
			splitted := strings.Split(trimmed, " ")
			for _, seedString := range splitted {
				seed, err := strconv.Atoi(seedString)
				if err == nil {
					seeds = append(seeds, seed)
				}
			}
		}
		if isStringValidNumericFormat(line) {
			currentMapping = append(currentMapping, line)
		}
		if line == "" {
			if len(currentMapping) > 0 {
				mappings = append(mappings, currentMapping)
			}
			currentMapping = make([]string, 0)
		}
	}

	for _, mainMapping := range mappings {
		mappings := make([]mappingStruct, 0)

		for _, mapping := range mainMapping {
			numbers := make([]int, 0)

			splitted := strings.Split(mapping, " ")
			for _, number := range splitted {
				seed, err := strconv.Atoi(number)
				if err == nil {
					numbers = append(numbers, seed)
				}
			}

			mappings = append(mappings, mappingStruct{
				Min:    numbers[1],
				Max:    numbers[1] + numbers[2] - 1,
				NewMin: numbers[0],
			})

		}

		for i, seed := range seeds {
			for _, mapping := range mappings {
				if seed >= mapping.Min && seed <= mapping.Max {
					diff := seed - mapping.Min
					seeds[i] = mapping.NewMin + diff
				}
			}
		}
	}

	slices.Sort(seeds)
	return seeds[0]
}
