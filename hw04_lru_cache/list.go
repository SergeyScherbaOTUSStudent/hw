package main

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}, key Key) *ListItem
	PushBack(v interface{}, key Key) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	length int
	head   *ListItem
	tail   *ListItem
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

func (l *list) PushFront(v interface{}, key Key) *ListItem {
	var item = newListItem(v, key)

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

func (l *list) PushBack(v interface{}, key Key) *ListItem {
	var item = newListItem(v, key)

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

	l.PushFront(li.Value, li.Key)
	l.Remove(li)
}

func NewList() List {
	return new(list)
}
