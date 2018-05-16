package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(a []string) FreqMap {
	ch := make(chan FreqMap)

	for _, s := range a {
		go func(s string) { ch <- Frequency(s) }(s)
	}

	freq := FreqMap{}
	for range a {
		for r, v := range <-ch {
			freq[r] += v
		}
	}
	return freq
}