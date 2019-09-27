package queue

import "errors"

// Queue .
type Queue []int

// Push append item at end
func (q *Queue) Push(item int) {
	*q = append(*q, item)
}

// Pop return the first item
func (q *Queue) Pop() (int, error) {
	if len(*q) == 0 {
		return -1, errors.New("queue is empty")
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item, nil
}
