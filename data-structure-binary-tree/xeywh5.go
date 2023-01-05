package data_structure_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 所谓什么序,是指“何时遍历根节点的顺序”,前序为根左右,中序左根右,后序左右根
//func front(node *TreeNode, answer *[]int) {
//	*answer = append(*answer, node.Val)
//	if node.Left != nil {
//		front(node.Left, answer)
//	}
//	if node.Right != nil {
//		front(node.Right, answer)
//	}
//}
//
//func preorderTraversal(root *TreeNode) (answer []int) {
//	if root != nil {
//		front(root, &answer)
//	}
//	return
//}

func preorderTraversal(root *TreeNode) (answer []int) {
	stack := []*TreeNode{}
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		answer = append(answer, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}
