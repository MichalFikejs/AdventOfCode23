package day09

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(`10  13  16  21  30  45`)

	result := partB(input)

	require.Equal(t, 5, result)
}
