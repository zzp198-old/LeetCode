package data_structure_binary_tree

func dfs(ans [][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	ans[level] = append(ans[level], node.Val)
	dfs(ans, node.Left, level+1)
	dfs(ans, node.Left, level+1)
}

func levelOrder(root *TreeNode) (ans [][]int) {
	dfs(ans, root, 0)
	return
}
