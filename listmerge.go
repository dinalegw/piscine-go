package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil { // If l1 is empty, just assign l2 to it
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	if l2.Head == nil { // If l2 is empty, nothing to merge
		return
	}

	// Link l1's tail to l2's head
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
