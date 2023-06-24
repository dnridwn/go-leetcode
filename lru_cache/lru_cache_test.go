package lru_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_1(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	assert.Equal(t, 1, lruCache.Get(1))

	lruCache.Put(3, 3)
	assert.Equal(t, -1, lruCache.Get(2))

	lruCache.Put(4, 4)
	assert.Equal(t, -1, lruCache.Get(1))
	assert.Equal(t, 3, lruCache.Get(3))
	assert.Equal(t, 4, lruCache.Get(4))
}

func TestLRUCache_2(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(2, 1)
	lruCache.Put(2, 2)
	assert.Equal(t, 2, lruCache.Get(2))

	lruCache.Put(1, 1)
	lruCache.Put(4, 1)

	assert.Equal(t, -1, lruCache.Get(2))
}

func TestLRUCache_3(t *testing.T) {
	lruCache := Constructor(2)
	assert.Equal(t, -1, lruCache.Get(2))

	lruCache.Put(2, 6)
	assert.Equal(t, -1, lruCache.Get(1))

	lruCache.Put(1, 5)
	lruCache.Put(1, 2)
	assert.Equal(t, 2, lruCache.Get(1))
	assert.Equal(t, 6, lruCache.Get(2))
}
