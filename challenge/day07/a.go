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

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards []rune
	Type  HandType
	Value int
	Bid   int
}

type ListOfHands []*Hand

var cardValues = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func (l ListOfHands) Len() int {
	return len(l)
}
func (l ListOfHands) Swap(i, n int) {
	l[i], l[n] = l[n], l[i]
}

func (l ListOfHands) Less(i, j int) bool {
	a := l[i]
	b := l[j]

	if a.Type != b.Type {
		return a.Type < b.Type
	}

	for k := 0; k < 5; k++ {
		if cardValues[a.Cards[k]] != cardValues[b.Cards[k]] {
			return cardValues[a.Cards[k]] < cardValues[b.Cards[k]]
		}
	}

	return false
}

func partA(challenge *challenge.Input) int {
	hands := make(ListOfHands, 0)

	for line := range challenge.Lines {
		hand := CreateHand(line)
		hands = append(hands, &hand)
	}

	sort.Sort(hands)

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}
	return sum
}

func CreateHand(line string) Hand {
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

	cardCounts := make([]int, 0)
	for _, count := range counts {
		cardCounts = append(cardCounts, count)
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
		panic(fmt.Sprintf("Unexpected card counts: %v", cardCounts))
	}
	return hand
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
