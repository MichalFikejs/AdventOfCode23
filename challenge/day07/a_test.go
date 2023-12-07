package day07

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

const testingInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(testingInput)

	result := partA(input)

	require.Equal(t, 6440, result)
}
