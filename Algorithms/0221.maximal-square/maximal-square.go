package problem0221

//
//func maximalSquare(matrix [][]byte) int {
//	m := len(matrix)
//	if m == 0 {
//		return 0
//	}
//	n := len(matrix[0])
//	if n == 0 {
//		return 0
//	}
//
//	maxEdge := 0
//	// dp[i][j] == 以 (i,j) 点为右下角点的符合题意的最大正方形的边长
//	// TODO: 由于 dp[i][j] 只与上，左上，左的数据有关，可以把 dp 压缩成一维的
//	dp := make([][]int, m)
//	for i := range dp {
//		dp[i] = make([]int, n)
//		if matrix[i][0] == '1' {
//			dp[i][0] = 1
//			maxEdge = 1
//		}
//	}
//	for j := 1; j < n; j++ {
//		if matrix[0][j] == '1' {
//			dp[0][j] = 1
//			maxEdge = 1
//		}
//	}
//
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			if matrix[i][j] == '1' {
//				dp[i][j] = 1 +
//					min(
//						dp[i-1][j-1],
//						min(
//							dp[i-1][j],
//							dp[i][j-1],
//						),
//					)
//				maxEdge = max(maxEdge, dp[i][j])
//			}
//		}
//	}
//
//	return maxEdge * maxEdge
//}
//
//func min(a int, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//在原数组上修改，占用内存更少
func maximalSquare(matrix [][]byte) int {
	l1 := len(matrix)
	if l1 == 0 {
		return 0
	}
	l2 := len(matrix[0])
	if l2 == 0 {
		return 0
	}
	res := 0
	for i := 0; i < l1; i++ {
		if matrix[i][0] == '1' {
			res = 1
			break
		}
	}

	for i := 0; i < l2; i++ {
		if matrix[0][i] == '1' {
			res = 1
			break
		}
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if matrix[i][j] == '1' {
				matrix[i][j] = 1 + min(matrix[i][j-1], min(matrix[i-1][j], matrix[i-1][j-1]))
				res = max(res, int(matrix[i][j]-'0'))
			}
		}
	}

	return res * res
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
