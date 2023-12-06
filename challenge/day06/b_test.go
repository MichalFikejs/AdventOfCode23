package day06

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(`Time:      7  15   30
Distance:  9  40  200`)

	result := partB(input)

	require.Equal(t, 71503, result)
}
