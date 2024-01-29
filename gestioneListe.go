package main

import (
	"fmt"
)

type listNode struct {
	next  *listNode
	prev  *listNode
	segno byte
	data  mattoncino
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
func newNode(data mattoncino, segno byte) *listNode {
	return &listNode{nil, nil, segno, data}
}

// insertisce alla fine della lista un nuovo nodo
func addNode(list *linkedList, node *listNode) {
	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
	}
}

// da completare con controlli nel caso il nodo sia il primo o l'ultimo
/*func deleteNode(list *linkedList, node *listNode) {
	current := list.head
	for current != nil {
		if current == node {
			current.prev.next = current.next
			current.next.prev = current.prev
			return
		}
		current = current.next
	}
}*/

// da completare con controlli nel caso il nodo sia il primo o l'ultimo
/*func insertNode(list *linkedList, node *listNode, id string) { // id è il nome del nodo dopo il quale inserire node
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
}*/

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

func printList(list *linkedList) {
	current := list.head
	for current != nil {
		switch current.segno {
		case '+':
			fmt.Printf("%s: %s, %s\n", current.data.sigma, current.data.alpha, current.data.beta)
		case '-':
			fmt.Printf("%s: %s, %s\n", current.data.sigma, current.data.beta, current.data.alpha)
		}
		current = current.next
	}
}
