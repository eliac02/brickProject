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

// da completare con controlli nel caso il nodo sia il primo o l'ultimo
func deleteNode(list *linkedList, node *listNode) {
	current := list.head
	for current != nil {
		if current == node {
			current.prev.next = current.next
			current.next.prev = current.prev
			return
		}
		current = current.next
	}
}

// da completare con controlli nel caso il nodo sia il primo o l'ultimo
func insertNode(list *linkedList, node *listNode, id string) { // id è il nome del nodo dopo il quale inserire node
	current := list.head
	name := current.data.sigma
	for current != nil {
		if name == id {
			node.next = current.next
			node.prev = current
			current.next.prev = node
			current.next = node
			return
		}
		current = current.next
		name = current.data.sigma
	}
}

func searchNode(list *linkedList, id string) *listNode {
	current := list.head
	for current != nil {
		if current.data.sigma == id {
			return current
		}
		current = current.next
	}
	return nil
}
