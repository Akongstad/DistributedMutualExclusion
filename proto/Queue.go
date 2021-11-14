package proto

import (
	"container/list"
	"fmt"
)

type Queue struct {
	queue *list.List
}

func (c *Queue) newQueue() Queue{
	return Queue{
		queue: list.New(),
	}
}

func (c *Queue) Enqueue(value int) {
	c.queue.PushBack(value)
}

func (c *Queue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *Queue) Front() (int, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(int); ok {
			return val, nil
		}
		return -1, fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return -1, fmt.Errorf("Peep Error: Queue is empty")
}

func (c *Queue) Size() int {
	return c.queue.Len()
}

func (c *Queue) Empty() bool {
	return c.queue.Len() == 0
}
