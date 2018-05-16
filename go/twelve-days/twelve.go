package twelve

import "strings"

func Song() string {
	return lyrics
}

func Verse(day int) string {
	return strings.Split(lyrics, "\n")[day-1]
}