package beer

import (
	"errors"
	"fmt"
)

const (
	lastVerse = `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`
	lineFormat = `%v bottle%[3]s of beer on the wall, %[1]v bottle%[3]s of beer.
Take %[5]s down and pass it around, %[2]v bottle%[4]s of beer on the wall.
`
)

func Verse(n int) (string, error) {
	switch {
	case n < 0 || n > 99:
		return "", errors.New("invalid verse")
	case n == 0:
		return lastVerse, nil
	default:
		return verse(n), nil
	}
}

func verse(n int) string {
	currPlural, nextPlural, take := "s", "s", "one"
	endBeers := interface{}(n - 1)
	if n == 1 {
		currPlural = ""
		take = "it"
		endBeers = "no more"
	} else if n == 2 {
		nextPlural = ""
	}
	return fmt.Sprintf(lineFormat, n, endBeers, currPlural, nextPlural, take)
}

func Verses(upper, lower int) (string, error) {
	if upper > 100 || lower < 0 || lower > upper {
		return "", errors.New("invalid verses")
	}
	return verses(upper, lower), nil

}

func verses(upper, lower int) string {
	var vs string
	for i := upper; i >= lower; i-- {
		vs += verse(i) + "\n"
	}
	return vs
}

func Song() string {
	return verses(99, 0)
}