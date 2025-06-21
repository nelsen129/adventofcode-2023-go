package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Part1(file io.Reader) (result int) {
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }

	return result
}

func Part2(file io.Reader) (result int) {
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }

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
