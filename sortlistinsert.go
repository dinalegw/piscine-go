package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	// Insert at the beginning or if list is empty
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	// Traverse list to find insertion point
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert new node
	newNode.Next = current.Next
	current.Next = newNode

	return l
}
