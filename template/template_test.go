package main_test

import (
	"bytes"
	daytemplate "github.com/nelsen129/adventofcode-2023-go/template"
	"testing"
)

var part1Input = `Input
Line 2`

var part2Input = `Input
Line 2`

func TestPart1(t *testing.T) {
	want := 0
	got := daytemplate.Part1(bytes.NewReader([]byte(part1Input)))

	if want != got {
		t.Errorf(`Part1 = %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 0
	got := daytemplate.Part2(bytes.NewReader([]byte(part2Input)))

	if want != got {
		t.Errorf(`Part2 = %d, want %d`, got, want)
	}
}
