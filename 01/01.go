package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
	"unicode"
)

func Part1(file io.Reader) int {
	calibration_value := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		ptr := 0

		for !unicode.IsDigit(rune(line[ptr])) {
			ptr += 1
		}

		c1 := rune(line[ptr]) - 48

		ptr = len(line) - 1

		for !unicode.IsDigit(rune(line[ptr])) {
			ptr -= 1
		}

		c2 := rune(line[ptr]) - 48

		calibration_value += int(c1)*10 + int(c2)
	}

	return calibration_value
}

func Part2(file io.Reader) int {
	calibration_value := 0

	digit_strings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		ptr := 0
		d := -1
		found := false

		for !unicode.IsDigit(rune(line[ptr])) && !found {
			for digit := range digit_strings {
				if ptr+len(digit) > len(line) {
					continue
				}
				if string(line[ptr:ptr+len(digit)]) == digit {
					d = digit_strings[digit]
					found = true
				}
			}
			ptr += 1
		}

		c1 := rune(line[ptr]) - 48
		if d != -1 {
			c1 = rune(d)
		}

		ptr = len(line) - 1
		d = -1
		found = false

		for !unicode.IsDigit(rune(line[ptr])) && !found {
			for digit := range digit_strings {
				if ptr+len(digit) > len(line) {
					continue
				}
				if string(line[ptr:ptr+len(digit)]) == digit {
					d = digit_strings[digit]
					found = true
				}
			}
			ptr -= 1
		}

		c2 := rune(line[ptr]) - 48
		if d != -1 {
			c2 = rune(d)
		}

		calibration_value += int(c1)*10 + int(c2)
	}

	return calibration_value
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
