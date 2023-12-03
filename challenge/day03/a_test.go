package day03

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	result := partA(input)

	require.Equal(t, 4361, result)
}
