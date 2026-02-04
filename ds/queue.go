package ds

import "fmt"

type Queue struct {
	list LinkedList
}

func (q *Queue) QueuePush(value string) { // add tail node
	q.list.InsertAt(q.list.GetSize(), value)
}

func (q *Queue) QueuePop() (string, error) { // remove the head
	// Check if Queue is empty
	if q.list.IsEmpty() {
		return "", fmt.Errorf("ERROR: Queue is empty so Pop does not work")
	}

	// Get Head data
	data := q.list.Head.data

	// Remove Head node
	q.list.RemoveAt(0)

	// Return data and status
	return data, nil
}
