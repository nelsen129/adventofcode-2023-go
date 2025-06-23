package main_test

import (
	"bytes"
	day07 "github.com/nelsen129/adventofcode-2023-go/07"
	"testing"
)

var part1Input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var part2Input = part1Input

func TestPart1(t *testing.T) {
	want := 6440
	got := day07.Part1(bytes.NewReader([]byte(part1Input)))

	if want != got {
		t.Errorf(`Part1 = %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 5905
	got := day07.Part2(bytes.NewReader([]byte(part2Input)))

	if want != got {
		t.Errorf(`Part2 = %d, want %d`, got, want)
	}
}
