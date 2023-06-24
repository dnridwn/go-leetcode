package lru_cache

type LRUCache struct {
	capacity     int
	caches       map[int]int
	orderMapping []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:     capacity,
		caches:       make(map[int]int),
		orderMapping: make([]int, 0),
	}
}

func (c *LRUCache) Get(key int) int {
	value, ok := c.caches[key]
	if !ok {
		return -1
	}
	c.reorder(key)
	return value
}

func (c *LRUCache) Put(key int, value int) {
	if len(c.caches) == c.capacity {
		_, ok := c.caches[key]
		if !ok {
			if len(c.orderMapping) > 0 {
				delete(c.caches, c.orderMapping[0])
				c.orderMapping = c.orderMapping[1:]
			}
		}
	}
	c.reorder(key)
	c.caches[key] = value
}

func (c *LRUCache) reorder(latest int) {
	exists := false
	for k, v := range c.orderMapping {
		if v == latest {
			exists = true
			if latest != c.orderMapping[len(c.orderMapping)-1] {
				c.orderMapping = append(c.orderMapping[0:k], c.orderMapping[k+1:]...)
				c.orderMapping = append(c.orderMapping, latest)
			}
		}
	}

	if !exists {
		c.orderMapping = append(c.orderMapping, latest)
	}
}
