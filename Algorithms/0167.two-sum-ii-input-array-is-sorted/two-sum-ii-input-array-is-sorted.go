package problem0167

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if m[target-n] != 0 {
			return []int{m[target-n], i + 1}
		}
		m[n] = i + 1
	}

	return nil
}

func twoSum2(nums []int, target int) []int {

	for i, j := 0, len(nums)-1; i < j; {
		switch {
		case nums[i]+nums[j] > target:
			j--
		case nums[i]+nums[j] < target:
			i++
		default:
			return []int{i, j}
		}
	}
	return nil
}
