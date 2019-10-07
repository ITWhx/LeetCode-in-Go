/**
*
*@author 吴昊轩
*@create 2019-10-0720:58
 */
package heap_sort

func heapSort(nums []int) {
	len := len(nums) - 1
	for i := len/2 - 1; i >= 0; i-- {
		buildHeap(nums, i, len)
	}
	for j := len; j >= 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		buildHeap(nums, 0, j-1)
	}
}

func buildHeap(nums []int, i, n int) {
	for {
		j1 := i*2 + 1
		if j1 > n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 <= n && nums[j2] < nums[j1] {
			j = j2
		}
		if nums[i] < nums[j] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

func findKNums(nums []int, n int) []int {

	target := nums[:n]
	l := len(target) - 1
	for i := l/2 - 1; i >= 0; i-- {
		buildHeap(target, i, l)
	}

	for i := n; i < len(nums); i++ {
		if nums[i] > target[0] {
			target[0] = nums[i]
			buildHeap(target, 0, l)
		}
	}
	return target
}
