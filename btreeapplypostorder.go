package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyPostorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyPostorder(root.Right, f)
	}
	_, _ = f(root.Data)
}
