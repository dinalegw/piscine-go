package piscine

func ListLast(l *List) interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	}

	// If the list is empty, return nil
	return nil
}
