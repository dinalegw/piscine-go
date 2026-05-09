package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	current := root

	for current != nil {
		if elem == current.Data {
			return current
		}
		if elem < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}
