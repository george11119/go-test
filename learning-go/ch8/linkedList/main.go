package main

import "fmt"

type List[T comparable] struct {
	head *Node[T]
	size int
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func (l *List[T]) Add(value T) {
	newNode := Node[T]{value, nil}
	l.size += 1

	if l.head == nil {
		l.head = &newNode
		return
	}

	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &newNode
}

func (l *List[T]) Insert(value T, index int) {
	tmp := l.head
	if index == 0 {
		l.head = &Node[T]{value, tmp}
		return
	}

	for range index - 1 {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}

	oldNext := tmp.next
	tmp.next = &Node[T]{value, oldNext}
}

func (l List[T]) Index(value T) int {
	tmp := l.head
	i := 0
	for tmp != nil {
		if tmp.value == value {
			return i
		}
		tmp = tmp.next
		i++
	}
	return -1
}

func (l List[T]) Display() {
	if l.head == nil {
		fmt.Println("<empty>")
		return
	}

	tmp := l.head
	for tmp != nil {
		fmt.Print(tmp.value)
		if tmp.next != nil {
			fmt.Print(" -> ")
		}
		tmp = tmp.next
	}
	fmt.Println()
}

func main() {
	l := List[int]{nil, 0}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Insert(5, 0)
	l.Display()
	fmt.Println(l.Index(1))
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(3))
	fmt.Println(l.Index(20))
}
