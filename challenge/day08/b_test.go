package day08

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

const bTestingInput = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestB(t *testing.T) {
	input := challenge.FromLiteral(bTestingInput)

	result := partB(input)

	require.Equal(t, 6, result)
}
