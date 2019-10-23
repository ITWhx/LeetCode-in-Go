/**
*
*@author 吴昊轩
*@create 2019-08-2319:06
 */
package quick_sort

func quick_sort(s []int, left, right int) {
	if left < right {
		par := getPar(s, left, right)
		quick_sort(s, left, par-1)
		quick_sort(s, par+1, right)
	}
}

func getPar(s []int, left, right int) int {
	fast := left
	last := right

	leftLen := 0
	rightLen := 0
	l1 := left
	r2 := right
	tmp := s[left]
	for left < right {
		for right > left && s[right] >= tmp {
			if s[right] == tmp {
				s[right], s[r2] = s[r2], s[right]
				rightLen++
				r2--
			}
			right--
		}
		s[left] = s[right]
		for right > left && s[left] <= tmp {
			if s[left] == tmp {
				s[l1], s[left] = s[left], s[l1]
				leftLen++
				l1++
			}
			left++
		}
		s[right] = s[left]
	}
	s[left] = tmp

	i := left - 1
	j := fast
	for j < l1 && s[i] != tmp {
		s[j], s[i] = s[i], s[j]
		j++
		i--
	}

	i2 := left + 1
	j2 := last

	for j2 > r2 && s[i2] != tmp {
		s[i2], s[j2] = s[j2], s[i2]
		i2++
		j2--
	}

	return left
}

func setPartion(s []int, l, r int) {
	mid := (l + r) / 2

	if s[mid] > s[r] {
		s[mid], s[r] = s[r], s[mid]
	}
	if s[l] > s[mid] {
		s[l], s[mid] = s[mid], s[l]
	}

	if s[mid] < s[r] {
		s[mid], s[l] = s[l], s[mid]
	}
}

//快速排序：通过一趟排序将待排序记录分割成独立的两部分，其中一部分记录的关键字均比另一部分关键字小，则分别
//对这两部分继续进行排序，直到整个序列有序。
func quikSort2(nums []int, low, high int) {
	for low < high {
		middle := getMiddle(nums, low, high)
		//对左子序列进行排序
		quick_sort(nums, low, middle+1)
		//对右子序列进行排序
		quick_sort(nums, middle+1, high)
	}
}

func getMiddle(nums []int, low, high int) int {
	//当前数组的第一个元素作为中轴（基准）
	tmp := nums[low]
	for low < high {
		//这里temp <= nums[high]中等号的情况相当于数组中出现了两个相等的数字，循环比较依然能够继续
		for low < high && nums[high] >= tmp {
			high--
		}
		//将大于基准值的high
		nums[low] = nums[high]
		for low < high && nums[low] <= tmp {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = tmp
	return low
}
