package piscine

func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		// empty list or single element, nothing to reverse
		return
	}

	var prev *NodeL
	current := l.Head
	l.Tail = l.Head // the old head will become the new tail

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev // prev is the new head
}
