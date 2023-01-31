package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem {
		return root
	} else {
		if root.Data < elem {
			if root.Right != nil {
				return BTreeSearchItem(root.Right, elem)
			} else {
				return nil
			}
		} else {
			if root.Left != nil {
				return BTreeSearchItem(root.Left, elem)
			} else {
				return nil
			}
		}
	}
}
