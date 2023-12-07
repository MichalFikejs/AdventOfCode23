package day07

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func CreateHandB(line string) Hand {
	parts := strings.Fields(line)

	bid, err := strconv.Atoi(parts[1])

	if err != nil {
		panic(err)
	}

	hand := Hand{
		Cards: []rune(parts[0]),
		Bid:   bid,
	}

	counts := make(map[rune]int)
	for _, card := range hand.Cards {
		counts[card]++
	}

	if counts['J'] != 0 {
		maxValue := 0
		var maxCard rune
		for c, v := range counts {
			if c != 'J' && v > maxValue {
				maxCard = c
				maxValue = v
			}
		}
		counts[maxCard] += counts['J']
		counts['J'] = 0
	}

	cardCounts := make([]int, 0)
	for _, count := range counts {
		if count > 0 {
			cardCounts = append(cardCounts, count)
		}
	}
	slices.Sort(cardCounts)

	switch {
	case equal(cardCounts, []int{5}):
		hand.Type = FiveOfAKind
	case equal(cardCounts, []int{1, 4}):
		hand.Type = FourOfAKind
	case equal(cardCounts, []int{2, 3}):
		hand.Type = FullHouse
	case equal(cardCounts, []int{1, 1, 3}):
		hand.Type = ThreeOfAKind
	case equal(cardCounts, []int{1, 2, 2}):
		hand.Type = TwoPair
	case equal(cardCounts, []int{1, 1, 1, 2}):
		hand.Type = OnePair
	case equal(cardCounts, []int{1, 1, 1, 1, 1}):
		hand.Type = HighCard
	default:
		fmt.Println("Unexpected card counts")
	}
	return hand
}

func partB(challenge *challenge.Input) int {
	hands := make(ListOfHands, 0)

	cardValues['J'] = -1

	for line := range challenge.Lines {
		hand := CreateHandB(line)
		hands = append(hands, &hand)
	}

	sort.Sort(hands)

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}
	return sum
}
