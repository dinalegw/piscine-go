package piscine

// ListForEachIf applies function f to nodes that satisfy cond
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}

// IsPositiveNode checks if the node's data is a positive int
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int:
		return node.Data.(int) > 0
	default:
		return false
	}
}

// IsAlNode checks if the node's data is not a numeric type
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int:
		return false
	default:
		return true
	}
}
