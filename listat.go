package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		current = current.Next
		index++
	}

	// Position out of bounds
	return nil
}
