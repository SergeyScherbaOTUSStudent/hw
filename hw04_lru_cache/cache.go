package main

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		item.Value = value

		return true
	}

	if c.capacity == c.queue.Len() {
		item := c.queue.Back()
		delete(c.items, item.Key)
		c.queue.Remove(item)
	}

	item := c.queue.PushFront(value, key)
	c.items[key] = item

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)

		return item.Value, ok
	}

	return nil, false
}

func (c *lruCache) Clear() {
	for k, item := range c.items {
		c.queue.Remove(item)
		delete(c.items, k)
	}
}
