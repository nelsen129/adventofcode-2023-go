package main_test

import (
	"bytes"
	day06 "github.com/nelsen129/adventofcode-2023-go/06"
	"testing"
)

var part1Input = `Time:      7  15   30
Distance:  9  40  200`

var part2Input = part1Input

func TestPart1(t *testing.T) {
	want := 288
	got := day06.Part1(bytes.NewReader([]byte(part1Input)))

	if want != got {
		t.Errorf(`Part1 = %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 71503
	got := day06.Part2(bytes.NewReader([]byte(part2Input)))

	if want != got {
		t.Errorf(`Part2 = %d, want %d`, got, want)
	}
}
