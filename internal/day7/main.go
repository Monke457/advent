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
	lines := reader.FileToArray("data/day7.txt")

	hands := parseHands(lines)
	slices.SortFunc(hands, func(a, b hand) int {
		return a.compare(b)
	})

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bet
	}
	fmt.Println(winnings)
}

func (h hand) compare(hand hand) int {
	if hand.t > h.t {
		return 1
	}
	if hand.t == h.t {
		return compareRunes(h.hand, hand.hand)
	}
	return -1
}

func compareRunes(a, b string) int {
	for i := 0; i < len(a); i++ {
		valA := cardValue(a[i])
		valB := cardValue(b[i])
		if valB > valA {
			return 1
		}
		if valB < valA {
			return -1
		}
	}
	return 0
}

func cardValue(b byte) int {
	switch b {
	case 65:
		return 14
	case 75:
		return 13
	case 81:
		return 12
	case 74:
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

func getType(h hand) int {
	cards := map[rune]int{}
	for _, card := range h.hand {
		cards[card]++
	}
	switch len(cards) {
	case 1:
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
		}
		return TWO_PAIR
	case 4:
		return PAIR
	default:
		return HIGH_CARD
	}
}

func parseHands(c []string) []hand {
	hands := []hand{}
	for _, line := range c {
		l := strings.Split(line, " ")
		bet, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}

		hand := hand{hand: l[0], bet: bet}
		hand.t = getType(hand)
		hands = append(hands, hand)
	}
	return hands
}
