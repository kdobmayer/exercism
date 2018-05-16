package brackets

var pairs = map[rune]rune{']': '[', '}': '{', ')': '('}

func Bracket(input string) (bool, error) {
	var stack []rune
	for _, r := range input {
		switch r {
		case '[', '{', '(':
			stack = append(stack, r)
		case ']', '}', ')':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false, nil
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0, nil
}