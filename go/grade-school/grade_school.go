package school

import "sort"

type Grade struct {
	grade    int
	students []string
}

type School []*Grade

func New() *School {
	return &School{}
}

func (s *School) Add(name string, grade int) {
	for _, g := range *s {
		if g.grade == grade {
			g.students = append(g.students, name)
			return
		}
	}
	*s = append(*s, &Grade{grade, []string{name}})
}

func (s *School) Grade(n int) []string {
	for _, g := range *s {
		if g.grade == n {
			return g.students
		}
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	var roaster []Grade
	for _, g := range *s {
		roaster = append(roaster, *g)
	}
	sort.Slice(roaster, func(i, j int) bool { return roaster[i].grade < roaster[j].grade })
	for _, g := range roaster {
		sort.Strings(g.students)
	}
	return roaster
}