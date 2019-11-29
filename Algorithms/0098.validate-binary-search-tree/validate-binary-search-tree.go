package problem0098

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {

	//Go int 类型的最小值与最大值
	MIN, MAX := -1<<63, 1<<63-1

	return recur(MIN, MAX, root)
}

// 以递归的方式，检查 root.Val 是否在 (min, max) 范围内。
func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max &&
		recur(min, root.Val, root.Left) &&
		recur(root.Val, max, root.Right)
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}
