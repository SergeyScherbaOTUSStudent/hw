package main

type ListItem struct {
	Value interface{}
	Key   Key
	Next  *ListItem
	Prev  *ListItem
}

func newListItem(val interface{}, key Key) *ListItem {
	return &ListItem{
		Value: val,
		Key:   key,
		Prev:  nil,
		Next:  nil,
	}
}
