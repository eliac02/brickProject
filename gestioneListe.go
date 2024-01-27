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

// crea una nuova lista
func newList() *linkedList {
	return &linkedList{nil, nil}
}

// crea un nuovo nodo
func newNode(data mattoncino) *listNode {
	return &listNode{nil, nil, data}
}

// insertisce alla fine della lista un nuovo nodo
func addNode(list *linkedList, node listNode) {
	if list.head == nil {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		node.prev = list.tail
		list.tail = &node
	}
}

func deleteNode(list *linkedList, node *listNode) {
}

func insertNode(list *linkedList, node *listNode) {
}