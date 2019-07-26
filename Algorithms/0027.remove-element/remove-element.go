package problem0027

func removeElement(nums []int, val int) int {
	// j指向最后一个不为val的位置
	// i指向第一个值为val的位置
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= 0 && nums[j] == val {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}

func removeElement1(nums []int, val int) int {
	//双指针，l1是慢指针，i是快指针。
	l1, res := 0, 0
	for i, v := range nums {

		if v == val {
			continue
		}
		//当前元素不等于val，进行元素交换，l1指针加1，变为下一次交换的位置
		nums[l1], nums[i] = nums[i], nums[l1]
		l1++
		res = l1

	}
	return res
}
