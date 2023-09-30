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

		l.PushFront(15, "")

		require.Equal(t, 1, l.Len())
		require.Equal(t, 15, l.Back().Value)
		require.Equal(t, 15, l.Front().Value)
	})

	t.Run("remove first and last items", func(t *testing.T) {
		l := NewList()

		fd0 := l.PushFront(10, "1")
		l.PushFront(15, "2")
		l.PushFront(20, "3")
		l.PushFront(25, "4")
		fd1 := l.PushFront(30, "5")

		l.Remove(fd1)
		l.Remove(fd0)

		require.Equal(t, 25, l.Front().Value)
		require.Equal(t, 15, l.Back().Value)
		require.Equal(t, 3, l.Len())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10, "")
		l.PushBack(20, "")
		l.PushBack(30, "")
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next
		l.Remove(middle)
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v, "")
			} else {
				l.PushBack(v, "")
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

		l.MoveToFront(newListItem(10, ""))

		require.Empty(t, l.Front())
	})

	t.Run("Push back. First item", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())

		l.PushBack(10, "")

		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 1, l.Len())
	})

	t.Run("move to front from tail", func(t *testing.T) {
		l := NewList()

		l.PushBack(10, "")
		l.PushBack(20, "")
		fm := l.PushBack(30, "")

		l.MoveToFront(fm)

		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, 30, l.Front().Value)
	})

	t.Run("remove all items", func(t *testing.T) {
		l := NewList()

		fr := l.PushFront(10, "")
		l.Remove(fr)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Back())
		require.Nil(t, l.Front())
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()

		a := l.PushFront("aaa", "1")
		b := l.PushFront("bbb", "2")

		l.MoveToFront(a)
		require.Equal(t, "aaa", l.Front().Value)

		l.MoveToFront(b)
		require.Equal(t, "bbb", l.Front().Value)
	})
}
