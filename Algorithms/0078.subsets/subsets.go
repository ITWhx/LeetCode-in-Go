package problem0078

import (
	"math"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0, resSize(nums))
	rec(nums, []int{}, &res)
	return res

}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func rec(nums, temp []int, res *[][]int) {
	size := len(nums)
	if size == 0 {
		*res = append(*res, temp)
		return
	}

	// 对于 last 来说，要么存在于某个子集中，要么不存在
	nums, last := nums[:size-1], nums[size-1]

	rec(nums, temp, res)

	withLast := make([]int, 1, 1+len(temp))
	withLast[0] = last
	withLast = append(withLast, temp...)
	rec(nums, withLast, res)
}

//位移法
func subsets2(nums []int) [][]int {
	res := [][]int{}
	num := 1 << uint32(len(nums)) //执行 2 的 n 次方
	for i := 0; i < num; i++ {
		tmp := []int{}
		idx := 0 //记录当前对应数组的哪一位
		iCopy := i
		for iCopy != 0 {
			if iCopy&1 == 1 { //判断当前位是否是 1
				tmp = append(tmp, nums[idx])
			}
			idx++
			iCopy >>= 1 //右移一位
		}
		res = append(res, tmp)
	}
	return res
}
