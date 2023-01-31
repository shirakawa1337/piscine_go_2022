package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else {
			changeValue := BTreeMin(root.Right).Data
			changeNode := BTreeSearchItem(root.Right, changeValue)

			if changeNode.Right != nil {
				changeNode.Parent.Left = changeNode.Right
				changeNode.Right.Parent = changeNode.Parent
			} else {
				changeNode.Parent.Left = nil
			}

			root.Data = changeValue
			return root
		}
	}
	return root
}
