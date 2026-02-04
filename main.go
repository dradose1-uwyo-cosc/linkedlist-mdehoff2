//Mason DeHoff
//COSC 3750
//[Date Here]
//
/*
	Don't forget to run your go mod init command in your terminal
	Review the assignment instructions for running your code
	All the code you need to write should be put in the /ds/ package files
	Uncomment the import statement for your module name
	you can uncomment the tests in main as you go to test
	The code in main is not an extensive test, you should add more and test your code as needed
*/
package main

import (
	"fmt"
	"hw1-linked-lists/ds"
)

func main() {
	fmt.Println("Linked List Test:")

	linkedlist := &ds.LinkedList{}
	linkedlist.InsertAt(0, "first")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("first")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth")
	linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	fmt.Println("Stack Test:")
	stack := &ds.Stack{}
	stack.StackPush("first")
	stack.StackPush("second")
	stack.StackPush("third")
	data, _ := stack.StackPop()
	println("Popped from stack:", data)
	fmt.Println("-------------")

	fmt.Println("Queue Test:")
	queue := &ds.Queue{}
	queue.QueuePush("first")
	queue.QueuePush("second")
	queue.QueuePush("third")
	data, _ = queue.QueuePop()
	println("Popped from queue:", data)
	fmt.Println("-------------")
}
