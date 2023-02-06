package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val == 1 {
			return true
		} else {
			return false
		}
	} else {
		l := evaluateTree(root.Left)
		r := evaluateTree(root.Right)
		if root.Val == 3 {
			return l && r
		} else {
			return l || r
		}
	}
}
