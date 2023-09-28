package main

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func newListItem(val interface{}) *ListItem {
	var item = new(ListItem)
	item.Value = val

	return item
}
