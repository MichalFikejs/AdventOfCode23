package day06

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`Time:      7  15   30
Distance:  9  40  200`)

	result := partA(input)

	require.Equal(t, 288, result)
}
