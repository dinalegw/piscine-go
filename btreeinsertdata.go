package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// Case where the tree is empty
	if root == nil {
		return &TreeNode{Data: data}
	}

	current := root

	for {
		if data < current.Data { // go left
			if current.Left == nil {
				current.Left = &TreeNode{
					Data:   data,
					Parent: current,
				}
				return root
			}
			current = current.Left
		} else { // go right
			if current.Right == nil {
				current.Right = &TreeNode{
					Data:   data,
					Parent: current,
				}
				return root
			}
			current = current.Right
		}
	}
}
