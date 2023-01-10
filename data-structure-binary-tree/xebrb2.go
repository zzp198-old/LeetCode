package data_structure_binary_tree

// 后序遍历特殊，非递归需要前序的“根右左”逆序变为”左右根“
func postorderTraversal(root *TreeNode) (answer []int) {
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		answer = append(answer, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// sort是降序,不是逆序
	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}
	return answer
}
