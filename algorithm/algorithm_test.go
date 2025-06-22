package algorithm_test

import (
	"github.com/nelsen129/adventofcode-2023-go/algorithm"
	"testing"

	"slices"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{`empty`, []int{}, []int{}},
		{`sorted`, []int{0, 1, 2, 3}, []int{0, 1, 2, 3}},
		{`unsorted`, []int{3, 0, 2, 1}, []int{0, 1, 2, 3}},
		{`duplicates`, []int{3, 0, 0, 0, 1}, []int{0, 0, 0, 1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := algorithm.QuickSort(tt.in)
			if !slices.Equal(got, tt.out) {
				t.Errorf(`got %v, want %v`, got, tt.out)
			}
		})

	}
}

func TestFindOverlapInSortedLists(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   [2][]int
		out  []int
	}{
		{`empty`, [2][]int{[]int{}, []int{}}, []int{}},
		{`none`, [2][]int{[]int{0, 1, 2, 3}, []int{4, 5, 6, 7}}, []int{}},
		{`some`, [2][]int{[]int{0, 2, 4, 6}, []int{1, 2, 3, 4}}, []int{2, 4}},
		{`all`, [2][]int{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}}, []int{0, 1, 2, 3}},
		{`duplicates`, [2][]int{[]int{0, 0, 1}, []int{0, 0, 3}}, []int{0, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := algorithm.FindOverlapInSortedLists(tt.in[0], tt.in[1])
			if !slices.Equal(got, tt.out) {
				t.Errorf(`got %v, want %v`, got, tt.out)
			}
		})

	}
}
