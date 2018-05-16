package robotname

import (
	"math/rand"
	"time"
)

var names []string

type Robot struct {
	name string
}

func randomName() string {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	name := ""
	for i := 0; i < 5; i++ {
		if i < 2 {
			name += string(rnd.Intn(26) + 65)
			continue
		}
		name += string(rnd.Intn(10) + 48)
	}
	return name
}

// Name is two uppercase letters followed by three digits.
func (r *Robot) Name() string {
	if r.name == "" {
		for {
			newName := randomName()
			if contains(names, newName) {
				continue
			}
			r.name = newName
			break
		}
		names = append(names, r.name)
	}
	return r.name
}

func contains(names []string, new string) bool {
	for _, n := range names {
		if n == new {
			return true
		}
	}
	return false
}

func (r *Robot) Reset() {
	r.name = ""
}