package main

type listNode struct {
	next *listNode
	prev *listNode
	data mattoncino
}

type linkedList struct {
	head *listNode
	tail *listNode
}

func newNode(data mattoncino) *listNode {
	return &listNode{nil, nil, data}
}

func addNode(row *linkedList, brick listNode) {
	if row.head == nil {
		row.head = &brick
		row.tail = &brick
	} else {

		//controlla questa parte della funzione
		row.tail.next = &brick
		brick.prev = row.tail
		row.tail = &brick
	}
}
