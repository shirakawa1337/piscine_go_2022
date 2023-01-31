package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(BTreeLevelCount(root.Left), BTreeLevelCount(root.Right))
}

func max(count int, count2 int) int {
	if count > count2 {
		return count
	}
	return count2
}
