package ds

type Stack struct {
	list LinkedList
}

func (s *Stack) StackPush(value string) { // add new head node
	// Insert at front (Head = 0)
	s.list.InsertAt(0, value)
}

func (s *Stack) StackPop() (string, bool) { // remove the head
	// Check if list is empty (We can use funcs from Linked List)
	if s.list.IsEmpty() {
		return "", false
	}

	// Get data from Head
	data := s.list.Head.data

	// Remove Head
	s.list.RemoveAt(0)

	// Return data and status
	return data, true
}
