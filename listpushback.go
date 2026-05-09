package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		// List is empty, new node is both head and tail
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Append at the end
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
