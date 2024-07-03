package main

import (
	"advent/internal/pkg/reader"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// hand types
const (
	FIVE       int = 7
	FOUR           = 6
	FULL_HOUSE     = 5
	THREE          = 4
	TWO_PAIR       = 3
	PAIR           = 2
	HIGH_CARD      = 1
)

type hand struct {
	hand string
	bet  int
	t    int
}

func main() {
	lines := reader.FileToArray("data/2023/day7.txt")
	fmt.Println(solveProblem(lines, false))
	fmt.Println(solveProblem(lines, true))
}

func solveProblem(lines []string, j bool) int {
	hands := parseHands(lines, j)
	slices.SortFunc(hands, func(a, b hand) int {
		return a.compare(b, j)
	})

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bet
	}
	return winnings
}

func (h hand) compare(hand hand, j bool) int {
	if hand.t < h.t {
		return 1
	}
	if hand.t == h.t {
		return compareRunes(h.hand, hand.hand, j)
	}
	return -1
}

func compareRunes(a, b string, j bool) int {
	for i := 0; i < len(a); i++ {
		valA := cardValue(a[i], j)
		valB := cardValue(b[i], j)
		if valB < valA {
			return 1
		}
		if valB > valA {
			return -1
		}
	}
	return 0
}

func cardValue(b byte, j bool) int {
	switch b {
	case 65:
		return 14
	case 75:
		return 13
	case 81:
		return 12
	case 74:
		if j {
			return 1
		}
		return 11
	case 84:
		return 10
	default:
		num, err := strconv.Atoi(fmt.Sprintf("%c", b))
		if err != nil {
			panic(err)
		}
		return num
	}
}

func getType(h hand, j bool) int {
	cards := map[rune]int{}
	jokers := 0
	for _, card := range h.hand {
		if j && card == 74 {
			jokers++
			continue
		}
		cards[card]++
	}
	switch len(cards) {
	case 1, 0:
		return FIVE
	case 2:
		for _, v := range cards {
			if v == 1 {
				return FOUR
			}
		}
		return FULL_HOUSE
	case 3:
		for _, v := range cards {
			if v == 3 {
				return THREE
			}
			if v == 2 {
				return TWO_PAIR + jokers
			}
		}
		return PAIR + jokers
	case 4:
		return PAIR
	default:
		return HIGH_CARD
	}
}

func parseHands(c []string, j bool) []hand {
	hands := []hand{}
	for _, line := range c {
		l := strings.Split(line, " ")
		bet, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}

		hand := hand{hand: l[0], bet: bet}
		hand.t = getType(hand, j)
		hands = append(hands, hand)
	}
	return hands
}
