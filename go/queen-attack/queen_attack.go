package queenattack

import (
	"errors"
	"math"
	"regexp"
)

var re = regexp.MustCompile("^[a-h][1-8]$")

func CanQueenAttack(white, black string) (bool, error) {
	if !re.MatchString(white) || !re.MatchString(black) || white == black {
		return false, errors.New("invalid input")
	}

	colDiff := math.Abs(float64(white[0]) - float64(black[0]))
	rowDiff := math.Abs(float64(white[1]) - float64(black[1]))
	if white[0] == black[0] || white[1] == black[1] || colDiff == rowDiff {
		return true, nil
	}
	return false, nil
}