package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func cardLess(a, b rune, jokers bool) bool {
	if jokers && b == 'J' {
		return false
	}
	if jokers && a == 'J' {
		return true
	}
	cardVal := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	return cardVal[a] < cardVal[b]
}

type CamelCard struct {
	Hand []rune
	Bid  int
}

func camelCardFromLine(line string) *CamelCard {
	cc := &CamelCard{}

	lineSplit := strings.Split(line, " ")
	cc.Hand = []rune(lineSplit[0])
	cc.Bid, _ = strconv.Atoi(lineSplit[1])

	return cc
}

func (cc CamelCard) handType(jokers bool) int {
	counts := map[rune]int{}

	for _, char := range cc.Hand {
		if _, ok := counts[char]; !ok {
			counts[char] = 1
		} else {
			counts[char] += 1
		}
	}

	jokerCount := counts['J']

	if jokers {
		delete(counts, 'J')
	}

	var m1, m2 int
	for _, c := range counts {
		if c > m1 {
			m1, m2 = c, m1
		} else if c > m2 {
			m2 = c
		}
	}

	if jokers {
		m1 += jokerCount
	}

	if m1 == 5 {
		return 10
	} else if m1 == 4 {
		return 8
	} else if m1 == 3 && m2 == 2 {
		return 6
	} else if m1 == 3 {
		return 4
	} else if m1 == 2 && m2 == 2 {
		return 3
	} else if m1 == 2 {
		return 2
	} else if m1 == 1 {
		return 0
	}
	return -10
}

func (cc CamelCard) Less(comp CamelCard, jokers bool) bool {
	if cc.handType(jokers) < comp.handType(jokers) {
		return true
	}
	if cc.handType(jokers) > comp.handType(jokers) {
		return false
	}
	for i := range 5 {
		if cc.Hand[i] != comp.Hand[i] {
			return cardLess(cc.Hand[i], comp.Hand[i], jokers)
		}
	}

	return false
}

type CamelCards []CamelCard

func (cc CamelCards) Len() int           { return len(cc) }
func (cc CamelCards) Less(i, j int) bool { return cc[i].Less(cc[j], false) }
func (cc CamelCards) Swap(i, j int)      { cc[i], cc[j] = cc[j], cc[i] }

type CamelCardsJoker []CamelCard

func (cc CamelCardsJoker) Len() int           { return len(cc) }
func (cc CamelCardsJoker) Less(i, j int) bool { return cc[i].Less(cc[j], true) }
func (cc CamelCardsJoker) Swap(i, j int)      { cc[i], cc[j] = cc[j], cc[i] }

func Part1(file io.Reader) (result int) {
	camelCards := CamelCards{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cc := camelCardFromLine(line)
		camelCards = append(camelCards, *cc)
	}
	sort.Sort(camelCards)

	for i, cc := range camelCards {
		result += (i + 1) * cc.Bid
	}

	return result
}

func Part2(file io.Reader) (result int) {
	camelCards := CamelCardsJoker{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cc := camelCardFromLine(line)
		camelCards = append(camelCards, *cc)
	}
	sort.Sort(camelCards)

	for i, cc := range camelCards {
		result += (i + 1) * cc.Bid
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
