package main_test

import (
	day05 "github.com/nelsen129/adventofcode-2023-go/05"
	"testing"

	"bytes"
	"slices"
)

var part1Input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

var part2Input = part1Input

func TestAlmanacRangesTidy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   day05.AlmanacRanges
		out  day05.AlmanacRanges
	}{
		{
			`sort`,
			day05.AlmanacRanges{day05.AlmanacRange{10, 1}, day05.AlmanacRange{5, 2}, day05.AlmanacRange{1, 3}},
			day05.AlmanacRanges{day05.AlmanacRange{1, 3}, day05.AlmanacRange{5, 2}, day05.AlmanacRange{10, 1}},
		}, {
			`merge`,
			day05.AlmanacRanges{day05.AlmanacRange{1, 3}, day05.AlmanacRange{4, 2}, day05.AlmanacRange{7, 1}},
			day05.AlmanacRanges{day05.AlmanacRange{1, 5}, day05.AlmanacRange{7, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.Tidy()
			if !slices.Equal(got, tt.out) {
				t.Errorf(`got %v, want %v`, got, tt.out)
			}
		})

	}
}

func TestConvertMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		inSrc []int
		inMap [3]int
		outS  []int
		outD  []int
	}{
		{`seed-to-soil 1`, []int{13, 14, 55, 79}, [3]int{50, 98, 2}, []int{13, 14, 55, 79}, []int{}},
		{`seed-to-soil 2`, []int{13, 14, 55, 79}, [3]int{52, 50, 48}, []int{13, 14}, []int{57, 81}},
		{`soil-to-fert 1`, []int{13, 14, 57, 81}, [3]int{0, 15, 37}, []int{13, 14, 57, 81}, []int{}},
		{`soil-to-fert 2`, []int{13, 14, 57, 81}, [3]int{37, 52, 2}, []int{13, 14, 57, 81}, []int{}},
		{`soil-to-fert 3`, []int{13, 14, 57, 81}, [3]int{39, 0, 15}, []int{57, 81}, []int{52, 53}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotS, gotD := day05.ConvertMap(tt.inSrc, tt.inMap[0], tt.inMap[1], tt.inMap[2])
			if !slices.Equal(gotS, tt.outS) || !slices.Equal(gotD, tt.outD) {
				t.Errorf(`got %v, %v, want %v, %v`, gotS, gotD, tt.outS, tt.outD)
			}
		})

	}
}

func TestPart1(t *testing.T) {
	want := 35
	got := day05.Part1(bytes.NewReader([]byte(part1Input)))

	if want != got {
		t.Errorf(`Part1 = %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 46
	got := day05.Part2(bytes.NewReader([]byte(part2Input)))

	if want != got {
		t.Errorf(`Part2 = %d, want %d`, got, want)
	}
}
