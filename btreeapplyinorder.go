package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	_, _ = f(root.Data)
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}
}
