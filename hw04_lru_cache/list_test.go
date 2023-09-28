package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("add first item and check general list", func(t *testing.T) {
		l := NewList()

		l.PushFront(15)

		require.Equal(t, 1, l.Len())
		require.Equal(t, 15, l.Back().Value)
		require.Equal(t, 15, l.Front().Value)
	})

	t.Run("delete last item", func(t *testing.T) {
		l := NewList()

		l.Remove(l.PushFront(15))

		require.Equal(t, 0, l.Len())
	})

	t.Run("remove first and last items", func(t *testing.T) {
		l := NewList()

		fd0 := l.PushFront(10)
		l.PushFront(15)
		l.PushFront(20)
		l.PushFront(25)
		fd1 := l.PushFront(30)

		l.Remove(fd1)
		l.Remove(fd0)

		require.Equal(t, 25, l.Front().Value)
		require.Equal(t, 15, l.Back().Value)
		require.Equal(t, 3, l.Len())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)
		l.PushBack(20)
		l.PushBack(30)
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next
		l.Remove(middle)
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		}

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front())
		l.MoveToFront(l.Back())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("move to front in empty list", func(t *testing.T) {
		l := NewList()

		l.MoveToFront(newListItem(10))

		require.Empty(t, l.Front())
	})
}
