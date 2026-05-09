package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: l.Head}

	// Update head
	l.Head = newNode

	// If the list was empty, also set tail
	if l.Tail == nil {
		l.Tail = newNode
	}
}
