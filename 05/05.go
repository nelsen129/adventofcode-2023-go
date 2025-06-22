package main

import (
	"github.com/nelsen129/adventofcode-2023-go/algorithm"

	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
	"unicode"
)

type AlmanacRange struct {
	Start  int
	Length int
}

func (ar AlmanacRange) ModifyOverlap(destRange int, conversion AlmanacRange) (AlmanacRanges, *AlmanacRange) {
	// check first if there is no overlap
	if ar.Start+ar.Length <= conversion.Start || ar.Start >= conversion.Start+conversion.Length {
		return AlmanacRanges{ar}, nil
	}

	sources := AlmanacRanges{}
	dest := AlmanacRange{}

	if ar.Start < conversion.Start {
		sources = append(sources, AlmanacRange{
			Start:  ar.Start,
			Length: conversion.Start - ar.Start,
		})
		dest.Start = conversion.Start
	} else {
		dest.Start = ar.Start
	}

	if ar.Start+ar.Length > conversion.Start+conversion.Length {
		postStart := conversion.Start + conversion.Length
		sources = append(sources, AlmanacRange{
			Start:  postStart,
			Length: ar.Start + ar.Length - postStart,
		})
		dest.Length = conversion.Start + conversion.Length - dest.Start
	} else {
		dest.Length = ar.Start + ar.Length - dest.Start
	}

	dest.Start += destRange - conversion.Start

	return sources, &dest
}

type AlmanacRanges []AlmanacRange

func (ar AlmanacRanges) Len() int           { return len(ar) }
func (ar AlmanacRanges) Less(i, j int) bool { return ar[i].Start < ar[j].Start }
func (ar AlmanacRanges) Swap(i, j int)      { ar[i], ar[j] = ar[j], ar[i] }

func (ar AlmanacRanges) Tidy() AlmanacRanges {
	sort.Sort(ar)

	i := 0
	for i < len(ar) {
		if ar[i].Length == 0 {
			ar = append(ar[:i], ar[i+1:]...)
		} else {
			i += 1
		}
	}

	i = 0
	for i < len(ar)-1 {
		if ar[i].Start+ar[i].Length < ar[i+1].Start {
			i += 1
			continue
		}

		ar[i].Length = ar[i+1].Start + ar[i+1].Length - ar[i].Start
		ar = append(ar[:i+1], ar[i+2:]...)
	}

	return ar
}

func parseNumFromLine(idx int, line []byte) (int, int) {
	end := idx
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end += 1
	}

	number, _ := strconv.Atoi(string(line[idx:end]))
	return number, end - idx
}

func parseNumsFromLine(line []byte) []int {
	nums := []int{}

	i := 0

	for i < len(line) {
		if !unicode.IsDigit(rune(line[i])) {
			i += 1
			continue
		}

		number, numlen := parseNumFromLine(i, line)
		i += numlen + 1
		nums = append(nums, number)
	}

	return nums
}

func ConvertMap(sources []int, destRange, sourceRange, sourceLength int) ([]int, []int) {
	dests := []int{}

	i, _ := algorithm.BinarySearchSorted(sourceRange, sources)
	j := 0
	for j+i < len(sources) && sources[j+i] < sourceRange+sourceLength {
		dests = append(dests, sources[j+i]-sourceRange+destRange)
		j += 1
	}

	sources = append(sources[:i], sources[i+j:]...)

	return sources, dests
}

func performAlmanacConversionRound(sources []int, scanner *bufio.Scanner) []int {
	dests := []int{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if line == nil || len(line) == 0 {
			break
		}

		conversion := parseNumsFromLine(line)
		var d []int
		sources, d = ConvertMap(sources, conversion[0], conversion[1], conversion[2])
		dests = append(dests, d...)
	}

	dests = algorithm.QuickSort(append(sources, dests...))
	return dests
}

func convertMapRanges(sources AlmanacRanges, destRange int, conversion AlmanacRange) (AlmanacRanges, AlmanacRanges) {
	dests := AlmanacRanges{}
	newSources := AlmanacRanges{}

	for _, source := range sources {
		ns, d := source.ModifyOverlap(destRange, conversion)
		newSources = append(newSources, ns...)
		if d != nil {
			dests = append(dests, *d)
		}
	}

	return newSources, dests
}

func performAlmanacRangeConversionRound(sources AlmanacRanges, scanner *bufio.Scanner) AlmanacRanges {
	dests := AlmanacRanges{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if line == nil || len(line) == 0 {
			break
		}

		conversion := parseNumsFromLine(line)
		destRange := conversion[0]
		convAr := AlmanacRange{
			Start:  conversion[1],
			Length: conversion[2],
		}
		var d AlmanacRanges
		sources, d = convertMapRanges(sources, destRange, convAr)
		dests = append(dests, d...)
	}

	dests = append(dests, sources...)
	dests = dests.Tidy()
	return dests
}

func Part1(file io.Reader) (result int) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds := parseNumsFromLine(scanner.Bytes())
	sources := algorithm.QuickSort(seeds)
	scanner.Scan()

	for scanner.Scan() {
		sources = performAlmanacConversionRound(sources, scanner)
	}
	result = sources[0]

	return result
}

func Part2(file io.Reader) (result int) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds := parseNumsFromLine(scanner.Bytes())
	sources := AlmanacRanges{}
	for i := 0; i+1 < len(seeds); i += 2 {
		sources = append(sources, AlmanacRange{
			Start:  seeds[i],
			Length: seeds[i+1],
		})
	}
	sources = sources.Tidy()
	scanner.Scan()

	for scanner.Scan() {
		sources = performAlmanacRangeConversionRound(sources, scanner)
	}
	result = sources[0].Start

	return result
}

func main() {
	start := time.Now()

	args := os.Args[1:]
	file_path := args[0]

	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", Part1(file))

	file.Seek(0, 0)

	fmt.Println("Part 2:", Part2(file))

	duration := time.Since(start)

	fmt.Println("Program execution time:", duration)
}
