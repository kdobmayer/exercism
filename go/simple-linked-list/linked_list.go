package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(a []int) *List {
	l := new(List)
	for _, x := range a {
		l.Push(x)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(x int) {
	e := &Element{data: x}
	if l.head == nil {
		l.head = e
		l.size++
		return
	}
	i := l.head
	for i.next != nil {
		i = i.next
	}
	i.next = e
	l.size++
}

func (l *List) Pop() (int, error) {
	var x int
	switch l.size {
	case 0:
		return 0, errors.New("empty list")
	case 1:
		x = l.head.data
		l.head = nil
	default:
		e := l.head
		for i := 1; i < l.size-1; i++ {
			e = e.next
		}
		x = e.next.data
		e.next = nil
	}
	l.size--
	return x, nil
}

func (l *List) Array() []int {
	var a []int
	for e := l.head; e != nil; e = e.next {
		a = append(a, e.data)
	}
	return a

}

func (l *List) Reverse() *List {
	a := l.Array()
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return New(a)
}