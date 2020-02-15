package problem0113

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		// 根据 level 来更新 path
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		// 到达 leaf
		// 并且，此条路径符合要求
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			// copy 不会复制 path[len(temp):]，如果 len(path) > len(temp) 的话
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	return res
}

//这种写法是错的，当加入4结点时，slice刚好扩容，递归调用5，1结点时传入的slice都是len=3，cap=4，5结点满足条件
//加入到res，1结点在append时，会改变5结点的slice
func pathSum2(root *TreeNode, sum int) [][]int {
	var res [][]int

	var dfs func(*TreeNode, []int, int)
	dfs = func(node *TreeNode, tmp []int, sum int) {
		if node == nil {
			return
		}
		tmp = append(tmp, node.Val)
		sum -= node.Val
		if node.Left == nil && node.Right == nil && sum == 0 {
			res = append(res, tmp)
			return
		}
		dfs(node.Left, tmp, sum)
		dfs(node.Right, tmp, sum)
	}

	dfs(root, []int{}, sum)
	return res
}
