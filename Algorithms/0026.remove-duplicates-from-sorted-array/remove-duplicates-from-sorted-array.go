package problem0026

//双双指针，一个快指针，一个慢指针
func removeDuplicates(a []int) int {
	left, right, size := 0, 1, len(a)
	//left是每次需要替换元素的下标
	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left]
	}
	return left + 1
}

//更快
func removeDuplicates1(nums []int) int {
	l1, l2 := 0, 1
	for ; l2 < len(nums); l2++ {
		if nums[l2] == nums[l1] {
			continue
		}
		l1++
		nums[l1], nums[l2] = nums[l2], nums[l1]
	}
	l1++
	return l1
}
