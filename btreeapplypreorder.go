package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	_, _ = f(root.Data)
	if root.Left != nil {
		BTreeApplyPreorder(root.Left, f)
	}
	if root.Right != nil {
		BTreeApplyPreorder(root.Right, f)
	}
}
