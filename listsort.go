package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	swapped := true
	for swapped {
		swapped = false
		current := l
		for current.Next != nil {
			if current.Data > current.Next.Data {
				// Swap values
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
	}

	return l
}
