package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
)

type Table struct{ Played, Win, Draw, Loss, Points int }

func Tally(r io.Reader, w io.Writer) error {
	var teams = map[string]*Table{
		"Allegoric Alaskians":     {},
		"Courageous Californians": {},
		"Blithering Badgers":      {},
		"Devastating Donkeys":     {},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		if !valid(line) {
			return fmt.Errorf("Invalid line: %s\n", line)
		}
		fields := strings.Split(line, ";")
		home := fields[0]
		away := fields[1]
		result := fields[2]
		switch result {
		case "win":
			teams[home].Played++
			teams[away].Played++
			teams[home].Win++
			teams[home].Points += 3
			teams[away].Loss++
		case "draw":
			teams[home].Played++
			teams[away].Played++
			teams[home].Draw++
			teams[away].Draw++
			teams[home].Points += 1
			teams[away].Points += 1
		case "loss":
			teams[home].Played++
			teams[away].Played++
			teams[home].Loss++
			teams[away].Win++
			teams[away].Points += 3
		default:
			return errors.New("Invalid result")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	var names []string
	for k := range teams {
		names = append(names, k)
	}
	sort.SliceStable(names, func(i, j int) bool {
		pi, pj := teams[names[i]].Points, teams[names[j]].Points
		return pi > pj || (pi == pj && names[i] < names[j])
	})

	tw := tabwriter.NewWriter(w, 0, 0, 8, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "Team\t MP |  W |  D |  L |  P")
	for _, name := range names {
		fmt.Fprintf(tw, "%s\t  %d |  %d |  %d |  %d |  %d\n", name, teams[name].Played, teams[name].Win, teams[name].Draw, teams[name].Loss, teams[name].Points)
	}
	tw.Flush()
	return nil
}

func valid(line string) bool {
	if strings.Count(line, ";") != 2 {
		return false
	}
	if !strings.Contains(line, "win") &&
		!strings.Contains(line, "draw") &&
		!strings.Contains(line, "loss") {
		return false
	}
	return true
}