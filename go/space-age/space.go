package space

const yearInSeconds = 365.25 * 24 * 60 * 60 // day * hour * minute * second

var orbitalPeriod = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Planet represents a planet in the solar system.
type Planet string

// Age says how old someone in different planets.
func Age(seconds float64, p Planet) float64 {
	return seconds / (orbitalPeriod[p] * yearInSeconds)
}