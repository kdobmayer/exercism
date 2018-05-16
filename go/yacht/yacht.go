package yacht

import "sort"

var scores = map[string]int{
	"ones":   1,
	"twos":   2,
	"threes": 3,
	"fours":  4,
	"fives":  5,
	"sixes":  6,
}

// sortDice sorts a slice of dice in decreasing order based on the count and the value.
func sortDice(dice []int) {
	counts := make(map[int]int)
	for _, n := range dice {
		counts[n]++
	}
	sort.Slice(dice, func(i, j int) bool {
		return counts[dice[i]] > counts[dice[j]] ||
			counts[dice[i]] == counts[dice[j]] && dice[i] > dice[j]
	})
}

func isStraight(dice []int) bool {
	for i := 1; i < len(dice); i++ {
		if dice[i] != dice[i-1]-1 {
			return false
		}
	}
	return true
}

func sum(dice []int) int {
	var sum int
	for _, n := range dice {
		sum += n
	}
	return sum
}

// Score returns the score of the dice for the given category.
func Score(dice []int, category string) int {
	sortDice(dice)
	switch category {
	case "yacht":
		if dice[0] == dice[4] {
			return 50
		}
	case "choice":
		return sum(dice)
	case "big straight":
		if isStraight(dice) && dice[0] == 6 {
			return 30
		}
	case "little straight":
		if isStraight(dice) && dice[0] == 5 {
			return 30
		}
	case "four of a kind":
		if dice[0] == dice[3] {
			return dice[0] * 4
		}
	case "full house":
		if dice[0] == dice[2] && dice[3] == dice[4] && dice[0] != dice[4] {
			return sum(dice)
		}
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		var score int
		for _, n := range dice {
			if n == scores[category] {
				score += scores[category]
			}
		}
		return score
	}
	return 0
}