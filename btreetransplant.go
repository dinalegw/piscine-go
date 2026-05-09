package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// If node is the root
	if node.Parent == nil {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	// Determine if node is left or right child of its parent
	if node.Parent.Left == node {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
