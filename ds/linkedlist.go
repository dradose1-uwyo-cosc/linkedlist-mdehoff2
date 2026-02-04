package ds

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Insert(value string) { // insert at the tail
	// Create new node with value from string parameter
	newNode := &Node{data: value}

	// Check if List is empty, and if so insert node at both head and tail
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else { // Add to tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++ // Increment size of List
}

func (l *LinkedList) InsertAt(position int, value string) error { // Insert at a position in the list
	if position < 0 || position > l.Size {
		return fmt.Errorf("ERROR: Index out of bounds of list")
	}
	// New Node
	newNode := &Node{data: value}

	// Insert at Head
	if position == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		if l.Tail == nil {
			l.Tail = newNode
		}
		l.Size++
		return nil
	}

	// Insert at Tail
	if position == l.Size {
		l.Tail.Next = newNode
		l.Tail = newNode
		l.Size++
		return nil
	}

	// Insert in the middle
	current := l.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	l.Size++

	return nil
}

func (l *LinkedList) Remove(value string) error { // Removes first occurrence of value parameter
	// Check if the list is empty
	if l.Head == nil {
		return fmt.Errorf("ERROR: Cannot delete, list is empty")
	}

	// Check if head is given value
	if l.Head.data == value {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		l.Size--
		return nil
	}

	// Traverse to find the value in list
	current := l.Head
	for current.Next != nil {
		if current.Next.data == value {
			current.Next = current.Next.Next // Skip over the node

			// Check if it was tail that was removed
			if current.Next == nil {
				l.Tail = current
			}

			l.Size--
			return nil
		}
		current = current.Next
	}
	return fmt.Errorf("ERROR: Value not found in list")
}

func (l *LinkedList) RemoveAll(value string) error { // Remove all occurences of value
	// Check if the list is empty
	if l.Head == nil {
		return fmt.Errorf("ERROR: Cannot delete, list is empty")
	}

	// Traverse through list and find all values that match parameter
	for l.Head != nil && l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
	}

	// Check if list is empty after removing heads
	if l.Head == nil {
		l.Tail = nil
		return nil
	}

	// Traverse through the rest of the list to find any remaining matches of value
	current := l.Head
	for current.Next != nil {
		if current.Next.data == value {
			current.Next = current.Next.Next
			l.Size--
		} else {
			current = current.Next
		}
	}

	// Update tail and return
	l.Tail = current
	return nil
}

func (l *LinkedList) RemoveAt(pos int) error {
	// Check if the list is empty
	if l.Head == nil {
		return fmt.Errorf("ERROR: Cannot delete, list is empty")
	}

	// Check if pos given is within bounds of the list
	if pos < 0 || pos >= l.Size {
		return fmt.Errorf("ERROR: Pos value is out of bounds of list")
	}

	// Check if pos = 0, then just remove the Head of the list
	if pos == 0 {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		l.Size--
		return nil
	}

	// Traverse through the list til one before the node to remove
	current := l.Head
	for i := 0; i < pos-1; i++ {
		current = current.Next
	}

	// Remove the node
	current.Next = current.Next.Next

	if pos == l.Size-1 {
		l.Tail = current
	}
	l.Size--
	return nil
}

func (l *LinkedList) IsEmpty() bool { // checks if the linked list is empty
	if l.Head == nil {
		return true
	}
	return false
}

func (l *LinkedList) GetSize() int { // get size of list
	return l.Size
}

func (l *LinkedList) Reverse() { //reverse the list
	// No need for reverse if list is empty or just 1 element
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	// Reverse list
	var prev *Node = nil // Initialize prev variable for reverse process
	current := l.Head
	l.Tail = l.Head // Flip Head and Tail
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev // Now the new head (last node in original list)
}

func (l *LinkedList) PrintList() { //print the list
	// Traverse through the list and print out each and every element
	current := l.Head
	for current != nil {
		fmt.Print(current.data)
		if current.Next != nil {
			fmt.Print("->") // Spacer between elements and show's the relationship
		}
		current = current.Next
	}
	fmt.Println()
}
