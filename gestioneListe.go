package main

import (
	"fmt"
)

// listNode represents a node in the linked list.
type listNode struct {
	next  *listNode
	prev  *listNode
	segno byte
	data  mattoncino
}

// linkedList reoresents a doubly linked list.
type linkedList struct {
	head *listNode
	tail *listNode
}

// newList creates a new empty linked list.
//
// @return A pointer to the newly created linked list.
func newList() *linkedList {
	return &linkedList{nil, nil}
}

// newNode creates a new node with the given data and sign.
//
// @param data The data of the new node, that are the name, left shape and right shape of a brick.
// @param segno The sign of the brick, "+" if it's given normal, or "-" if it's given reversed.
// @return A pointer to the newly created node.
func newNode(data mattoncino, segno byte) *listNode {
	return &listNode{nil, nil, segno, data}
}

// addNode appends a new node to the end of the list.
//
// @param list The linked list to which the node will be added.
// @param node The node to be added.
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

// deleteNode removes a node from the list.
//
// @param list The linked list from which the node will be removed.
// @param node The node to be removed.
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

// insertNode inserts a new node after a specified node in the list.
//
// @param list The linked list in which the node will be inserted.
// @param node The node to be inserted.
// @param id The name of the node after which the new node will be inserted.
/*func insertNode(list *linkedList, node *listNode, id string) {
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

// searchNode searches for a node with the given name in the list.
//
// @param list The linked list to be searched.
// @param name The name of the node.
// @return A pointer to the found node, or nil if not found.
func searchNode(list *linkedList, name string) *listNode {
	current := list.head
	for current != nil {
		if current.data.sigma == name {
			return current
		}
		current = current.next
	}
	return nil
}

// printList prints the contents of the linked list.
//
// @param list The linked list to be printed.
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
