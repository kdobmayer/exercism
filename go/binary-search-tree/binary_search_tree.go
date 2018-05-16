package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

func (s *SearchTreeData) Insert(n int) {
	if s.data >= n {
		if s.left == nil {
			s.left = Bst(n)
			return
		}
		s.left.Insert(n)
		return
	}
	if s.right == nil {
		s.right = Bst(n)
		return
	}
	s.right.Insert(n)
}

func (s *SearchTreeData) MapString(f func(int) string) []string {
	var mapped []string

	if s.left != nil {
		mapped = append(mapped, s.left.MapString(f)...)
	}
	mapped = append(mapped, f(s.data))
	if s.right != nil {
		mapped = append(mapped, s.right.MapString(f)...)
	}

	return mapped
}

func (s *SearchTreeData) MapInt(f func(int) int) []int {
	var mapped []int

	if s.left != nil {
		mapped = append(mapped, s.left.MapInt(f)...)
	}
	mapped = append(mapped, f(s.data))
	if s.right != nil {
		mapped = append(mapped, s.right.MapInt(f)...)
	}

	return mapped
}