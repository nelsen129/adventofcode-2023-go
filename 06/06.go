package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
	"unicode"
)

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

func parseNumFromLineOnlyDigits(line []byte) (num int) {
	for _, char := range line {
		if !unicode.IsDigit(rune(char)) {
			continue
		}

		num *= 10
		num += int(char - byte('0'))
	}

	return num
}

func QuadraticRealZeroes(a, b, c float64) (float64, float64, bool) {
	d := b*b - 4*a*c

	if d < 0 {
		return 0, 0, false
	}

	s := math.Sqrt(d)

	z1 := (-b - s) / (2 * a)
	z2 := (-b + s) / (2 * a)

	return z1, z2, true
}

func getValidTimes(time, dist int) int {
	zero1, zero2, _ := QuadraticRealZeroes(-1, float64(time), -1*float64(dist))
	start := max(0, 1+int(math.Floor(min(zero1, zero2))))
	end := min(time, -1+int(math.Ceil(max(zero1, zero2))))

	return end - start + 1
}

func Part1(file io.Reader) (result int) {
	result = 1
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Bytes()
	times := parseNumsFromLine(line)

	scanner.Scan()
	line = scanner.Bytes()
	dists := parseNumsFromLine(line)

	for i := range len(times) {
		result *= getValidTimes(times[i], dists[i])
	}

	return result
}

func Part2(file io.Reader) (result int) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Bytes()
	time := parseNumFromLineOnlyDigits(line)

	scanner.Scan()
	line = scanner.Bytes()
	dist := parseNumFromLineOnlyDigits(line)

	result = getValidTimes(time, dist)

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
