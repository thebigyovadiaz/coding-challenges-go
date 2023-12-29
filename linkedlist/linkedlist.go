package linkedlist

import "fmt"

// Node represents a node in the linked list.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a singly linked list.
type LinkedList struct {
	Head *Node
}

// Function to reverse a linked list.
func reverseLinkedList(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	return prev
}

// Function to detect a cycle in a linked list using Floyd's cycle-finding algorithm.
func hasCycle(head *Node) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true // Cycle detected
		}
	}

	return false // No cycle
}

func Linked() {
	// Creating a linked list: 1 -> 2 -> 3 -> 4 -> 5
	list := LinkedList{
		Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5}}}}},
	}

	// Reversing the linked list
	reversedList := reverseLinkedList(list.Head)

	// Printing the reversed linked list
	fmt.Println("Reversed Linked List:")
	current := reversedList
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")

	// Creating a cyclic linked list for demonstration purposes
	list.Head.Next.Next.Next.Next.Next = list.Head.Next.Next

	// Checking if the linked list has a cycle
	hasCycle := hasCycle(list.Head)
	fmt.Printf("Does the linked list have a cycle? %t\n", hasCycle)
}
