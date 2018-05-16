package poker

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

// Ranks of face cards.
const (
	J = 11
	Q = 12
	K = 13
	A = 14
)

// Card suits.
const (
	Hearts   = '\u2661' // ♡
	Diamonds = '\u2662' // ♢
	Spades   = '\u2664' // ♤
	Clubs    = '\u2667' // ♧
)

// Hand ranks.
const (
	HighCard = iota + 1
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

// Card represents a poker card.
type Card struct {
	Rank int
	Suit rune
}

// SetRank sets the rank of the card to r if it is a valid rank.
func (c *Card) SetRank(r rune) error {
	if '2' <= r && r <= '9' {
		c.Rank = int(r - '0')
		return nil
	}

	switch r {
	case '0':
		c.Rank = 10
	case 'J':
		c.Rank = J
	case 'Q':
		c.Rank = Q
	case 'K':
		c.Rank = K
	case 'A':
		c.Rank = A
	default:
		return errors.New("invalid rank")
	}
	return nil
}

// SetSuit sets the suit of the card to r if it is a valid suit.
func (c *Card) SetSuit(r rune) error {
	switch r {
	case Hearts, Diamonds, Spades, Clubs:
		c.Suit = r
	default:
		return errors.New("invalid suit")
	}
	return nil
}

func parseCard(s string) (Card, error) {
	if s[0] == '1' {
		s = s[1:]
	}

	var c Card
	if err := c.SetRank(rune(s[0])); err != nil {
		return Card{}, fmt.Errorf("set rank (%c): %s", s[0], err)
	}
	r, size := utf8.DecodeRuneInString(s[1:])
	if err := c.SetSuit(r); err != nil {
		return Card{}, fmt.Errorf("set suit (%c): %s", r, err)
	}
	if s[size+1:] != "" {
		return Card{}, fmt.Errorf("characters after suit (%s)", s[size+1:])
	}
	return c, nil
}

// Hand represents a poker hand.
type Hand struct {
	Rank  int
	Cards []Card
}

func parseHand(str string) (Hand, error) {
	var cards []Card
	for _, s := range strings.Fields(str) {
		card, err := parseCard(s)
		if err != nil {
			return Hand{}, fmt.Errorf("parse card (%s): %s", s, err)
		}
		cards = append(cards, card)
	}
	if len(cards) != 5 {
		return Hand{}, fmt.Errorf("invalid number of cards (%d)", len(cards))
	}

	counts := make(map[int]int)
	for _, card := range cards {
		counts[card.Rank]++
	}
	sort.Slice(cards, func(i, j int) bool {
		if counts[cards[i].Rank] > counts[cards[j].Rank] {
			return true
		}
		if counts[cards[i].Rank] == counts[cards[j].Rank] {
			return cards[i].Rank > cards[j].Rank
		}
		return false
	})

	h := Hand{Cards: cards}
	switch counts[cards[0].Rank] {
	case 4:
		h.Rank = FourOfAKind
	case 3:
		if counts[cards[3].Rank] == 2 {
			h.Rank = FullHouse
			break
		}
		h.Rank = ThreeOfAKind
	case 2:
		if counts[cards[2].Rank] == 2 {
			h.Rank = TwoPair
			break
		}
		h.Rank = Pair
	case 1:
		if isStraight(cards) && isFlush(cards) {
			h.Rank = StraightFlush
			break
		}
		if isStraight(cards) {
			h.Rank = Straight
			break
		}
		if isFlush(cards) {
			h.Rank = Flush
			break
		}
		h.Rank = HighCard
	}
	return h, nil
}

func (h Hand) isBetter(than Hand) bool {
	if h.Rank != than.Rank {
		return h.Rank > than.Rank
	}

	if h.Rank == Straight || h.Rank == StraightFlush {
		if h.Cards[0].Rank == A && h.Cards[1].Rank == 5 {
			return false
		}
		if than.Cards[0].Rank == A && than.Cards[1].Rank == 5 {
			return true
		}
	}

	for i := 0; i < len(h.Cards); i++ {
		if h.Cards[i].Rank != than.Cards[i].Rank {
			return h.Cards[i].Rank > than.Cards[i].Rank
		}
	}

	return false
}

func isStraight(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		rank, prank := cards[i].Rank, cards[i-1].Rank
		if rank != prank-1 {
			if i == 1 && rank == 5 && prank == A {
				continue
			}
			return false
		}
	}
	return true
}

func isFlush(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[i].Suit != cards[i-1].Suit {
			return false
		}
	}
	return true
}

// BestHand select the strongest hands from multiple hands.
func BestHand(a []string) ([]string, error) {
	var bestCards []string
	var bestHand Hand

	for _, s := range a {
		hand, err := parseHand(s)
		if err != nil {
			return nil, fmt.Errorf("parse hand (%s): %s", s, err)
		}

		if hand.isBetter(bestHand) {
			bestCards, bestHand = []string{s}, hand
			continue
		}
		if bestHand.isBetter(hand) {
			continue
		}
		bestCards = append(bestCards, s)
	}

	return bestCards, nil
}