package piscine

// CompStr compares two interfaces for equality
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind returns a pointer to the Data of the first node where comp(node.Data, ref) is true
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head

	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}

	// Not found
	return nil
}
