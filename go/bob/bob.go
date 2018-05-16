package bob

import "strings"

const (
	toQuestion     = "Sure."
	toAnythingElse = "Whatever."
	toYell         = "Whoa, chill out!"
	toYellQuestion = "Calm down, I know what I'm doing!"
	toNothing      = "Fine. Be that way!"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case isNothing(remark):
		return toNothing
	case isQuestion(remark) && isYelling(remark):
		return toYellQuestion
	case isQuestion(remark):
		return toQuestion
	case isYelling(remark):
		return toYell
	default:
		return toAnythingElse
	}
}

func isNothing(remark string) bool {
	return remark == ""
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYelling(remark string) bool {
	return remark == strings.ToUpper(remark) &&
		remark != strings.ToLower(remark)
}