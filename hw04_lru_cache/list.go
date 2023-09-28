package main

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	length int
	head   *ListItem
	tail   *ListItem
	List
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(val interface{}) *ListItem {
	var item = newListItem(val)

	if l.length != 0 {
		item.Next = l.head
		l.Front().Prev = item
	} else {
		l.tail = item
	}

	l.head = item
	l.length++

	return item
}

func (l *list) PushBack(val interface{}) *ListItem {
	var item = newListItem(val)

	if l.length != 0 {
		item.Prev = l.tail
		l.Back().Next = item
	} else {
		l.head = item
	}

	l.tail = item
	l.length++

	return item
}

func (l *list) Remove(li *ListItem) {
	left := li.Prev
	right := li.Next

	if nil != left {
		left.Next = right
	}

	if nil != right {
		right.Prev = left
	}

	if li == l.head {
		l.head = right
	}

	if li == l.tail {
		l.tail = left
	}

	l.length--
}

func (l *list) MoveToFront(li *ListItem) {
	if li == l.Front() || l.length == 0 {
		return
	}

	left := li.Prev
	right := li.Next

	left.Next = right

	if nil != right {
		right.Prev = left
	}

	li.Next = l.Front()
	l.head = li
}

func NewList() List {
	return new(list)
}
