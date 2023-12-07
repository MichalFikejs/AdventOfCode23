package day07

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(testingInput)

	result := partB(input)

	require.Equal(t, 5905, result)
}
