package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func toInt(data string) int {
	res := 0
	for i := 0; i < len(data); i++ {
		res *= 10
		res += int(data[i] - '0')
	}
	return res
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if toInt(root.Data) <= toInt(data) {
		if root.Right != nil {
			BTreeInsertData(root.Right, data)
		} else {
			rightNode := &TreeNode{Data: data, Parent: root}
			root.Right = rightNode
		}
	} else {
		if root.Left != nil {
			BTreeInsertData(root.Left, data)
		} else {
			leftNode := &TreeNode{Data: data, Parent: root}
			root.Left = leftNode
		}
	}
	return root
}
