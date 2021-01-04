package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low, high int) int {
	// 递归
	// if root != nil {
	// 	val := 0
	// 	if root.Val >= low && root.Val <= high {
	// 		val = root.Val
	// 	}
	// 	return val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	// }
	// return 0

	// 中序遍历非递归
	stack, sum := make([]*TreeNode, 0), 0
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			l := len(stack)
			root = stack[l-1]
			stack = stack[:l-1]
			if root.Val >= low && root.Val <= high {
				sum += root.Val
			}
			root = root.Right
		}
	}
	return sum
}
