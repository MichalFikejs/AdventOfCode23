package day08

import (
	"testing"

	"github.com/michalfikejs/AdventOfCode23/challenge"
	"github.com/stretchr/testify/require"
)

const testingInput = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const testingInput2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(testingInput)

	result := partA(input)

	require.Equal(t, 2, result)
}

func TestASecondInput(t *testing.T) {
	input := challenge.FromLiteral(testingInput2)

	result := partA(input)

	require.Equal(t, 6, result)
}
