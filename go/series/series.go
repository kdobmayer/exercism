package series

func All(n int, s string) []string {
	var series []string
	for i := 0; i <= len(s)-n; i++ {
		series = append(series, s[i:i+n])
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (string, bool) {
	all := All(n, s)
	if len(all) == 0 {
		return "", false
	}
	return all[0], true
}