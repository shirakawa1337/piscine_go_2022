package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	id := 0
	var queue []*TreeNode

	queue = append(queue, root)
	for id < len(queue) {
		cur := queue[id]
		_, _ = f(cur.Data)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		id++
	}
}
