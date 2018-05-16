package gigasecond

import "time"

// AddGigasecond calculates the moment after 10^9 seconds the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}