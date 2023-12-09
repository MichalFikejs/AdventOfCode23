package day09

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

const testingInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(testingInput)

	result := partA(input)

	require.Equal(t, 114, result)
}
