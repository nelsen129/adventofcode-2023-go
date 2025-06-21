package main_test

import (
	"bytes"
	day2 "github.com/nelsen129/adventofcode-2023-go/02"
	"testing"
)

var part1Input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var part2Input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestPart1(t *testing.T) {
	want := 8
	got, err := day2.Part1(bytes.NewReader([]byte(part1Input)))

	if err != nil {
		t.Errorf(`Got err %v, expected none`, err)
	}

	if want != got {
		t.Errorf(`Part1 = %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2286
	got, err := day2.Part2(bytes.NewReader([]byte(part2Input)))

	if err != nil {
		t.Errorf(`Got err %v, expected none`, err)
	}

	if want != got {
		t.Errorf(`Part2 = %d, want %d`, got, want)
	}
}
