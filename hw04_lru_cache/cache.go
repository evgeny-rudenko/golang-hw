package hw04lrucache

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

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	cItem := cacheItem{key, value}
	// уже есть значение
	if item, ok := c.items[key]; ok {
		c.items[key].Value = cItem
		c.queue.MoveToFront(item)
		return true
	}
	// нет значения и вышли за границы размера кэша
	if c.queue.Len() >= c.capacity {
		itemForDel := c.queue.Back()
		c.queue.Remove(itemForDel)
		delete(c.items, itemForDel.Value.(cacheItem).key)
	}
	c.queue.PushFront(cItem)
	c.items[key] = c.queue.Front()
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return c.queue.Front().Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
