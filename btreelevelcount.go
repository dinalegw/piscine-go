package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLevel := BTreeLevelCount(root.Left)
	rightLevel := BTreeLevelCount(root.Right)

	if leftLevel > rightLevel {
		return leftLevel + 1
	}
	return rightLevel + 1
}
