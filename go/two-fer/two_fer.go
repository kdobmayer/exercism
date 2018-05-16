package twofer

import "fmt"

// ShareWith gives one for name and one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}