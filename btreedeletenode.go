package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	// Case 1: node has no left child
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}

	// Case 2: node has no right child
	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	// Case 3: node has both children
	// Find the successor (minimum in right subtree)
	successor := BTreeMin(node.Right)

	if successor.Parent != node {
		// Replace successor with its right child
		root = BTreeTransplant(root, successor, successor.Right)
		// Attach node's right child to successor
		successor.Right = node.Right
		if successor.Right != nil {
			successor.Right.Parent = successor
		}
	}

	// Replace node with successor
	root = BTreeTransplant(root, node, successor)
	successor.Left = node.Left
	if successor.Left != nil {
		successor.Left.Parent = successor
	}

	return root
}
