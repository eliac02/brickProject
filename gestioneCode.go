package main

type queue struct {
	head *queueElement
	tail *queueElement
}

type queueElement struct {
	value string
	next  *queueElement
}

func newQueue() *queue {
	return &queue{nil, nil}
}

func newQueueElement(value string) *queueElement {
	return &queueElement{value, nil}
}

func enqueue(q *queue, value string) {
	if q.head == nil {
		q.head = newQueueElement(value)
		q.tail = q.head
	} else {
		q.tail.next = newQueueElement(value)
		q.tail = q.tail.next
	}
}

func dequeue(q *queue) {
	if q.head == nil {
		return
	}
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
}

func isEmpty(q *queue) bool {
	return q.head == nil
}

func bottom(q *queue) string {
	if q.head == nil {
		return ""
	}
	return q.head.value
}

/*func top(q *queue) string {
	if q.tail == nil {
		return ""
	}
	return q.tail.value
}*/
