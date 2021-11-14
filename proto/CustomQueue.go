package proto

import (
	"sync"
)

type customQueue struct {
	queue []int
	lock  sync.RWMutex
}

func (c *queue) Enqueue(name int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue, name)
}

func (c *queue) Dequeue() {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.queue = c.queue[1:]
	}
}

func (c *queue) Front() int {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.queue[0]
	}
	return -1
}

func (c *queue) Size() int {
	return len(c.queue)
}

func (c *queue) Empty() bool {
	return len(c.queue) == 0
}
