package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"unicode"
)

func getGearOptions(symbols map[[2]int]byte) map[[2]int][]int {
	gearOpts := map[[2]int][]int{}
	for coord, sym := range symbols {
		if sym != byte('*') {
			continue
		}

		gearOpts[coord] = []int{}
	}

	return gearOpts
}

func getSymbols(file io.Reader) map[[2]int]byte {
	symbols := map[[2]int]byte{}

	scanner := bufio.NewScanner(file)
	lineCount := -1

	for scanner.Scan() {
		line := scanner.Bytes()
		lineCount += 1

		for col, b := range line {
			if b == byte('.') {
				continue
			}
			if unicode.IsDigit(rune(b)) {
				continue
			}

			symbols[[2]int{lineCount, col}] = b
		}
	}

	return symbols
}

func parseIntFromLine(idx int, line []byte) (int, int) {
	end := idx
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end += 1
	}

	number, _ := strconv.Atoi(string(line[idx:end]))
	return number, end - idx
}

func validPart(row, col, numlen int, symbols map[[2]int]byte) bool {
	i := row - 1
	for j := col - 1; j <= (col + numlen); j += 1 {
		if _, ok := symbols[[2]int{i, j}]; ok {
			return true
		}
	}

	if _, ok := symbols[[2]int{row, col - 1}]; ok {
		return true
	}

	if _, ok := symbols[[2]int{row, col + numlen}]; ok {
		return true
	}

	i = row + 1
	for j := col - 1; j <= (col + numlen); j += 1 {
		if _, ok := symbols[[2]int{i, j}]; ok {
			return true
		}
	}

	return false
}

func addValidParts(file io.Reader, symbols map[[2]int]byte) (result int) {
	scanner := bufio.NewScanner(file)
	lineCount := -1

	for scanner.Scan() {
		line := scanner.Bytes()
		lineCount += 1

		col := 0
		for col < len(line) {
			if !unicode.IsDigit(rune(line[col])) {
				col += 1
				continue
			}

			number, numlen := parseIntFromLine(col, line)
			if validPart(lineCount, col, numlen, symbols) {
				result += number
			}

			col += numlen + 1
		}
	}

	return result
}

func addGearOpt(row, col, numlen, number int, gearOpts map[[2]int][]int) {
	i := row - 1
	for j := col - 1; j <= (col + numlen); j += 1 {
		if _, ok := gearOpts[[2]int{i, j}]; ok {
			gearOpts[[2]int{i, j}] = append(gearOpts[[2]int{i, j}], number)
		}
	}

	if _, ok := gearOpts[[2]int{row, col - 1}]; ok {
		gearOpts[[2]int{row, col - 1}] = append(gearOpts[[2]int{row, col - 1}], number)
	}

	if _, ok := gearOpts[[2]int{row, col + numlen}]; ok {
		gearOpts[[2]int{row, col + numlen}] = append(gearOpts[[2]int{row, col + numlen}], number)
	}

	i = row + 1
	for j := col - 1; j <= (col + numlen); j += 1 {
		if _, ok := gearOpts[[2]int{i, j}]; ok {
			gearOpts[[2]int{i, j}] = append(gearOpts[[2]int{i, j}], number)
		}
	}
}

func addValidGearOpts(gearOpts map[[2]int][]int) (result int) {
	for _, nums := range gearOpts {
		if len(nums) != 2 {
			continue
		}
		result += nums[0] * nums[1]
	}
	return
}

func addGears(file io.Reader, gearOpts map[[2]int][]int) int {
	scanner := bufio.NewScanner(file)
	lineCount := -1

	for scanner.Scan() {
		line := scanner.Bytes()
		lineCount += 1

		col := 0
		for col < len(line) {
			if !unicode.IsDigit(rune(line[col])) {
				col += 1
				continue
			}

			number, numlen := parseIntFromLine(col, line)
			addGearOpt(lineCount, col, numlen, number, gearOpts)
			col += numlen + 1
		}
	}

	return addValidGearOpts(gearOpts)
}

func Part1(file io.ReadSeeker) (result int) {
	symbols := getSymbols(file)

	file.Seek(0, 0)

	result = addValidParts(file, symbols)

	return result
}

func Part2(file io.ReadSeeker) (result int) {
	symbols := getSymbols(file)
	gearOpts := getGearOptions(symbols)

	file.Seek(0, 0)

	result = addGears(file, gearOpts)

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
