package hw04lrucache

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type ItemBody struct {
	Value interface{}
	Key   Key
}

func newListItem(val interface{}) *ListItem {
	return &ListItem{
		Value: val,
		Prev:  nil,
		Next:  nil,
	}
}

func NewItemBody(val interface{}, key Key) *ItemBody {
	return &ItemBody{
		Value: val,
		Key:   key,
	}
}
