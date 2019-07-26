package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	m := make(map[int]struct{})
	m[2] = struct{}{}
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//从当前字符后一个，数组最后一个字符开始计算
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			//每计算一个，根据sum值判断向后移动l，还是向前移动r
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				//结果等于零，则计算当前字符是否还有和等于0的其他组合
				l, r = next(nums, l, r)
			}
		}
	}
	return res
}
func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}

func removeDuplicates(nums []int) int {
	temp := []int{}
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			temp = append(temp, v)
			m[v] = struct{}{}
		}
	}

	nums = temp
	return len(temp)
}
