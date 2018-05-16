// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(n int) bool {
	if n%4 == 0 && n%25 != 0 || n%16 == 0 {
		return true
	}
	return false
}