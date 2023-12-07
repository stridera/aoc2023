package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const testData = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

// ERROR:  {[8 11 11 7 10] 877 1} {[7 9 5 12 7] 394 1}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

const (
	JOKER = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

type Hand struct {
	display string
	cards   []int
	bid     int
	rank    int
}

func toHand(cards []rune, jokers bool) []int {
	hand := []int{}
	for _, card := range cards {
		switch card {
		case '2':
			hand = append(hand, TWO)
		case '3':
			hand = append(hand, THREE)
		case '4':
			hand = append(hand, FOUR)
		case '5':
			hand = append(hand, FIVE)
		case '6':
			hand = append(hand, SIX)
		case '7':
			hand = append(hand, SEVEN)
		case '8':
			hand = append(hand, EIGHT)
		case '9':
			hand = append(hand, NINE)
		case 'T':
			hand = append(hand, TEN)
		case 'J':
			if jokers {
				hand = append(hand, JOKER)
			} else {
				hand = append(hand, JACK)
			}
		case 'Q':
			hand = append(hand, QUEEN)
		case 'K':
			hand = append(hand, KING)
		case 'A':
			hand = append(hand, ACE)
		}
	}
	return hand
}

func parseHands(data string, jokers bool) []Hand {
	hands := []Hand{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		cards := strings.Split(line, " ")
		bid, _ := strconv.Atoi(cards[1])
		hand := Hand{cards[0], toHand([]rune(cards[0]), jokers), bid, HighCard}
		hands = append(hands, hand)
	}
	return hands
}

func calc(hands []Hand) int {

	for i := range hands {
		dict := map[int]int{}
		for _, card := range hands[i].cards {
			dict[card]++
		}
		counts := []int{}
		joker_count := 0
		for val, count := range dict {
			if val == JOKER {
				joker_count = count
			} else {
				counts = append(counts, count)
			}
		}
		sort.Ints(counts)
		if joker_count > 0 {
			if len(counts) == 0 {
				counts = append(counts, joker_count)
			} else {
				counts[len(counts)-1] += joker_count
			}
		}
		if len(counts) == 1 {
			hands[i].rank = FiveKind
		} else if len(counts) == 2 {
			if counts[0] == 1 {
				hands[i].rank = FourKind
			} else {
				hands[i].rank = FullHouse
			}
		} else if len(counts) == 3 {
			if counts[2] == 3 {
				hands[i].rank = ThreeKind
			} else if counts[2] == 2 {
				hands[i].rank = TwoPair
			}
		} else if len(counts) == 4 {
			hands[i].rank = OnePair
		} else {
			hands[i].rank = HighCard
		}
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank != hands[j].rank {
			return hands[i].rank > hands[j].rank
		}
		for ii := 0; ii < len(hands[i].cards); ii++ {
			if hands[i].cards[ii] != hands[j].cards[ii] {
				return hands[i].cards[ii] > hands[j].cards[ii]
			}
		}
		fmt.Println("ERROR: ", hands[i], hands[j])
		return false
	})

	sum := 0
	for i, hand := range hands {
		rank := len(hands) - i
		// fmt.Println("Hand: ", hand, " Rank: ", rank, " Bid: ", hand.bid, " Winnings: ", hand.bid*rank)
		sum += hand.bid * rank
	}
	return sum
}

func test() {
	hands := parseHands(testData, false)
	p1 := calc(hands)
	fmt.Println("TEST Part 1: ", p1, "Expects: 6440 Pass? ", p1 == 6440)
	hands = parseHands(testData, true)
	p2 := calc(hands)
	fmt.Println("TEST Part 2: ", p2, "Expects: 5905 Pass? ", p2 == 5905)
}

func main() {
	test()
	// return

	// read file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	hands := parseHands(string(data), false)
	p1 := calc(hands)
	fmt.Println("Part 1: ", p1)
	hands = parseHands(string(data), true)
	p2 := calc(hands)
	fmt.Println("Part 2: ", p2)
}
