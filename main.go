package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Reverse a linked list and returns new linked list
func reverseAndClone(listNode *node) *node {
	var head *node

	for listNode != nil {
		n := &node{data: listNode.data}
		n.next = head
		head = n
		listNode = listNode.next
	}

	return head
}

// Checks if the nodes of first linked list equals to the nodes of the second linked list
func equals(node1 *node, node2 *node) bool {
	for node1 != nil && node2 != nil {
		if node1.data != node2.data {
			return false
		}

		node1 = node1.next
		node2 = node2.next
	}

	return node1 == nil && node2 == nil
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Checks if the elements of the linked list form a palindrome
func (list *linkedList) isPalindrome() bool {
	reversed := reverseAndClone(list.head)
	result := equals(list.head, reversed)
	return result
}

func main() {
	node1 := &node{data: 'a'}
	node2 := &node{data: 'n'}
	node3 := &node{data: 'a'}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)

	isPalindrome := linkedList1.isPalindrome()

	fmt.Println("Is it a palindrome?: ", isPalindrome)
}
