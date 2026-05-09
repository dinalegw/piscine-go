package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "", false, false)
}

// min, max = boundaries
// minSet, maxSet = indicate whether boundaries should be applied
func isBST(node *TreeNode, min, max string, minSet, maxSet bool) bool {
	if node == nil {
		return true
	}

	// If a min bound exists, enforce: min <= node.Data
	if minSet && node.Data < min {
		return false
	}

	// If a max bound exists, enforce: node.Data < max
	// Strict max because duplicates go on the RIGHT in piscine
	if maxSet && node.Data >= max {
		return false
	}

	return isBST(node.Left, min, node.Data, minSet, true) &&
		isBST(node.Right, node.Data, max, true, maxSet)
}
