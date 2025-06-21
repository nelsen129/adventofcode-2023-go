package main_test

import (
	"bytes"
	day1 "github.com/nelsen129/adventofcode-2023-go/01"
	"testing"
)

var part1Input = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var part2Input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPart1(t *testing.T) {
	want := 142
	got := day1.Part1(bytes.NewReader([]byte(part1Input)))

	if want != got {
		t.Errorf(`Part1 = %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 281
	got := day1.Part2(bytes.NewReader([]byte(part2Input)))

	if want != got {
		t.Errorf(`Part2 = %d, want %d`, got, want)
	}
}
