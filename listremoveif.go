package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove matching nodes at the head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If list is now empty
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Remove matching nodes in the middle or end
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	// Update Tail
	l.Tail = current
}
