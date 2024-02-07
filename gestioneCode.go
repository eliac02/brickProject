package main

// queue represents a queue data structure
type queue struct {
	head *queueElement
	tail *queueElement
}

// queueElement represents an element in the queue.
type queueElement struct {
	value string
	next  *queueElement
}

// newQueue creates a new empty queue.
//
// @return A pointer to the newly created queue.
func newQueue() *queue {
	return &queue{nil, nil}
}

// newQueueElement creates a new element with the given value.
//
// @param value The value of the new element.
// @return A pointer to the newly created element.
func newQueueElement(value string) *queueElement {
	return &queueElement{value, nil}
}

// enqueue is a method that adds a new element with the given value to the end of the queue.
//
// @param value The value to be added to the queue.
func (q *queue) enqueue(value string) {
	if q.head == nil {
		q.head = newQueueElement(value)
		q.tail = q.head
	} else {
		q.tail.next = newQueueElement(value)
		q.tail = q.tail.next
	}
}

// dequeue is a method that removes the first element from the queue.
func (q *queue) dequeue() {
	if q.head == nil {
		return
	}
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
}

// isEmpty is a method that checks if the queue is empty.
//
// @return true if the queue is empty, false otherwise.
func (q *queue) isEmpty() bool {
	return q.head == nil
}

// bottom is a method that returns the value of the first element in the queue.
//
// @return The value of the first element, or an empty string if the queue is empty.
func (q *queue) bottom() string {
	if q.head == nil {
		return ""
	}
	return q.head.value
}

// top is a method that returns the value of the last element in the queue.
//
// @param q The queue.
// @return The value of the last element, or an empty string if the queue is empty.
/*func top(q *queue) string {
	if q.tail == nil {
		return ""
	}
	return q.tail.value
}*/
