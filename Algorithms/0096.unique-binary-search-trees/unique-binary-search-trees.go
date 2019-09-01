package problem0096

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 5
	}

	res := 0
	// 左右对称，所以只做一半
	for i := 1; i <= n/2; i++ {
		res += numTrees(i-1) * numTrees(n-i)
	}
	res *= 2
	// 处理 n 为奇数的情况
	if n%2 == 1 {
		temp := numTrees(n / 2)
		res += temp * temp
	}

	return res
}

//递归解法
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	return getAns(1, n)
}

func getAns(start, end int) []*TreeNode {

	var res []*TreeNode
	if start > end {
		return append(res, nil)
	}

	if start == end {
		return append(res, &TreeNode{Val: start})
	}

	for i := start; i < end+1; i++ {
		left := getAns(start, i-1)
		right := getAns(i+1, end)
		for _, v := range left {
			for _, v2 := range right {
				r := &TreeNode{Val: i}
				r.Left, r.Right = v, v2
				res = append(res, r)
			}
		}
	}

	return res
}
