package main

import (
    "container/list"
    "fmt"
)

type Queue struct {
    queue *list.List
}

func (c *Queue) Enqueue(value string) {
    c.queue.PushBack(value)
}

func (c *Queue) Dequeue() error {
    if c.queue.Len() > 0 {
        ele := c.queue.Front()
        c.queue.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *Queue) Front() (string, error) {
    if c.queue.Len() > 0 {
        if val, ok := c.queue.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *Queue) Size() int {
    return c.queue.Len()
}

func (c *Queue) Empty() bool {
    return c.queue.Len() == 0
}