package main

import (
	"fmt"
)

type Item struct {
	next, prev *Item
	data       interface{}
}

func (n *Item) Value() interface{} {
	return n.data
}

func (n *Item) Next() *Item {
	return n.next
}

func (n *Item) Prev() *Item {
	return n.prev
}

type List struct {
	last, first *Item
	len         int
}

// return length of list
func (l *List) Len() int {
	return l.len
}

// Returns first item
func (l *List) First() *Item {
	return l.first
}

// Returns last item
func (l *List) Last() *Item {
	return l.last
}

// adds item to the beginning of the list
func (l *List) PushFront(v interface{}) *Item {
	newitem := Item{data: v}
	if l.last == nil {
		l.last = &newitem
	}
	if l.first != nil {
		newitem.next, l.first.prev = l.first, &newitem
	}
	l.first = &newitem
	l.len++
	return &newitem

}

// adds item to the end of the list
func (l *List) PushBack(v interface{}) *Item {
	newitem := Item{data: v}
	if l.last != nil {
		newitem.prev, l.last.next = l.last, &newitem
	}
	if l.first == nil {
		l.first = &newitem
	}

	l.last = &newitem
	l.len++
	return &newitem
}

func (l *List) Remove(i *Item) {
	//If the item is center in the list
	if i.prev != nil && i.next != nil {
		i.prev.next, i.next.prev = i.next, i.prev
	}
	//If the item is first in the list
	if i.prev == nil && i.next != nil {
		i.next.prev, l.first = nil, i.next
	}
	//If the item is last in the list
	if i.prev != nil && i.next == nil {
		i.prev.next, l.last = nil, i.prev
	}

	l.len--
}

func main() {
	testlist := new(List)
	testlist.PushBack(1)
	second := testlist.PushBack(2)
	testlist.PushFront(0)
	testlist.PushBack(3)
	testlist.Remove(second)

	//Iterates through the list
	for i := testlist.First(); i != nil; i = i.Next() {
		fmt.Println(i.Value())
	}
	fmt.Println("Количество элементов:", testlist.Len())
}
