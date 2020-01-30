package problem0300

import "sort"

func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))

	for _, n := range nums {
		at := sort.SearchInts(tails, n)
		if at == len(tails) {
			// n 比 tails 中所有的数都大
			// 把 n 放入 tails 的尾部
			tails = append(tails, n)
		} else if tails[at] > n {
			// tails[at-1] < n < tails[at]
			// tails[at] = n, 不会改变 tail 的递增性，却增加了加入更多数的可能性
			tails[at] = n
		}
	}

	return len(tails)
}

func lengthOfLIS2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		tmp := 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				tmp = max(1+dp[j], tmp)
			}
		}
		dp[i] = tmp
		res = max(tmp, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
