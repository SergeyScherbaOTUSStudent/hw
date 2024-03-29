package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("set item with same key", func(t *testing.T) {
		c := NewCache(4)

		c.Set("A", "aA")
		c.Set("B", "bB")
		c.Set("A", "zZzZ")

		item, _ := c.Get("A")

		require.Equal(t, "zZzZ", item)
	})

	t.Run("set item to full capacity Cache", func(t *testing.T) {
		c := NewCache(3)

		c.Set("A", "aA")
		c.Set("B", "bB")
		c.Set("C", "cC")
		c.Set("D", "dD")

		_, ok := c.Get("A")

		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("get", func(t *testing.T) {
		c := NewCache(4)

		c.Set("A", 300)
		val, ok := c.Get("A")

		require.Equal(t, 300, val)
		require.True(t, ok)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(4)

		c.Set("A", "aA")
		c.Set("B", "bB")
		c.Set("A", "zZzZ")

		c.Clear()
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
