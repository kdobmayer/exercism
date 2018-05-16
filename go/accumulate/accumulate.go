package accumulate

func Accumulate(input []string, f func(string) string) []string {
	output := make([]string, len(input))
	for i, v := range input {
		output[i] = f(v)
	}
	return output
}