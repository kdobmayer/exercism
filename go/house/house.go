package house

import "fmt"

var lines = []struct{ subject, predicate string }{
	{"house that Jack built", "lay in"},
	{"malt", "ate"},
	{"rat", "killed"},
	{"cat", "worried"},
	{"dog", "tossed"},
	{"cow with the crumpled horn", "milked"},
	{"maiden all forlorn", "kissed"},
	{"man all tattered and torn", "married"},
	{"priest all shaven and shorn", "woke"},
	{"rooster that crowed in the morn", "kept"},
	{"farmer sowing his corn", "belonged to"},
	{"horse and the hound and the horn", ""},
}

func Verse(n int) string {
	return fmt.Sprintf("This is the %s%s", lines[n-1].subject, verse(n-1))
}

func verse(n int) string {
	if n == 0 {
		return "."
	}
	return fmt.Sprintf("\nthat %s the %s%s", lines[n-1].predicate, lines[n-1].subject, verse(n-1))
}

func Song() string {
	song, sep := "", ""
	for i := 1; i <= 12; i++ {
		song += sep + Verse(i)
		sep = "\n\n"
	}
	return song
}