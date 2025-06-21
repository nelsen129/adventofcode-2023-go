package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func Part1(file io.Reader) (int, error) {
	valid_id_sum := 0

	total_cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")
		valid := true

		for _, round := range rounds {
			if !valid {
				break
			}
			colors := strings.Split(round, ", ")
			for _, color := range colors {
				color_draw := strings.Split(color, " ")
				color_count, err := strconv.Atoi(color_draw[0])
				if err != nil {
					return 0, err
				}

				if color_count > total_cubes[color_draw[1]] {
					valid = false
					break
				}
			}
		}

		if valid {
			game_id, err := strconv.Atoi(strings.Split(game[0], " ")[1])
			if err != nil {
				return 0, err
			}
			valid_id_sum += game_id
		}
	}

	return valid_id_sum, nil
}

func Part2(file io.Reader) (int, error) {
	total_power := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")
		game_min_cubes := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}

		for _, round := range rounds {
			colors := strings.Split(round, ", ")
			for _, color := range colors {
				color_draw := strings.Split(color, " ")
				color_count, err := strconv.Atoi(color_draw[0])
				if err != nil {
					return 0, err
				}

				game_min_cubes[color_draw[1]] = max(game_min_cubes[color_draw[1]], color_count)
			}
		}

		power := 1
		for _, v := range game_min_cubes {
			power *= v
		}
		total_power += power
	}

	return total_power, nil
}

func main() {
	start := time.Now()

	args := os.Args[1:]
	file_path := args[0]

	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}

	p1, err := Part1(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", p1)

	file.Seek(0, 0)

	p2, err := Part2(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 2:", p2)

	duration := time.Since(start)

	fmt.Println("Program execution time:", duration)
}
