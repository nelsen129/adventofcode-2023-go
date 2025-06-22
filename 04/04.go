package main

import (
	"github.com/nelsen129/adventofcode-2023-go/algorithm"

	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"unicode"
)

type card struct {
	winning []int
	have    []int
}

func cardFromLine(line []byte) *card {
	c := &card{
		winning: []int{},
		have:    []int{},
	}
	idx := 0
	for idx < len(line) && line[idx] != byte(':') {
		idx += 1
	}
	idx += 1

	for idx < len(line) && line[idx] != byte('|') {
		if !unicode.IsDigit(rune(line[idx])) {
			idx += 1
			continue
		}

		number, numlen := parseIntFromLine(idx, line)
		idx += numlen + 1
		c.winning = append(c.winning, number)
	}
	idx += 1

	for idx < len(line) {
		if !unicode.IsDigit(rune(line[idx])) {
			idx += 1
			continue
		}

		number, numlen := parseIntFromLine(idx, line)
		idx += numlen + 1
		c.have = append(c.have, number)
	}

	c.sort()
	return c
}

func (c *card) sort() {
	c.winning = algorithm.QuickSort(c.winning)
	c.have = algorithm.QuickSort(c.have)
}

func (c *card) score() int {
	overlap := algorithm.FindOverlapInSortedLists(c.winning, c.have)

	if len(overlap) == 0 {
		return 0
	}
	return 1 << (len(overlap) - 1)
}

func (c *card) winCount() int {
	return len(algorithm.FindOverlapInSortedLists(c.winning, c.have))
}

func parseIntFromLine(idx int, line []byte) (int, int) {
	end := idx
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end += 1
	}

	number, _ := strconv.Atoi(string(line[idx:end]))
	return number, end - idx
}

func Part1(file io.Reader) (result int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		c := cardFromLine(line)
		result += c.score()
	}

	return result
}

func Part2(file io.Reader) (result int) {
	scanner := bufio.NewScanner(file)
	bonusCards := []int{}

	for scanner.Scan() {
		line := scanner.Bytes()
		c := cardFromLine(line)
		scratches := 1
		if len(bonusCards) > 0 {
			scratches += bonusCards[0]
			bonusCards = bonusCards[1:]
		}
		result += scratches

		wins := c.winCount()
		i := 0
		for i < len(bonusCards) && i < wins {
			bonusCards[i] += scratches
			i += 1
		}
		for i < wins {
			bonusCards = append(bonusCards, scratches)
			i += 1
		}
	}

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
