package problem0101

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(node1, node2 *TreeNode) bool {
	if node1 == nil {
		return node2 == nil
	}
	if node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return helper(node1.Left, node2.Right) && helper(node1.Right, node2.Left)
}
