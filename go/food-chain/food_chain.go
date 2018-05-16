package foodchain

import (
	"fmt"
	"strings"
)

const (
	first   = "I know an old lady who swallowed a %s."
	general = "She swallowed the %s to catch the %s."
	spider  = "wriggled and jiggled and tickled inside her."
)

var animals = []struct{ Kind, Comment string }{
	{"", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It " + spider},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

func Verse(n int) string {
	lines := []string{fmt.Sprintf(first, animals[n].Kind), animals[n].Comment}
	for i := n; i >= 2; i-- {
		lines = append(lines, fmt.Sprintf(general, animals[i].Kind, animals[i-1].Kind))
	}
	lines[n-1] = strings.Replace(lines[n-1], "spider.", "spider that "+spider, 1)
	lines = append(lines, animals[1].Comment)
	if n == 1 || n == 8 {
		lines = lines[:2]
	}
	return strings.Join(lines, "\n")
}

func Verses(n, m int) string {
	var vs []string
	for i := n; i <= m; i++ {
		vs = append(vs, Verse(i))
	}
	return strings.Join(vs, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}