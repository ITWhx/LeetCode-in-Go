package problem0215

import "container/heap"

func findKthLargest(nums []int, k int) int {
	temp := highHeap(nums)
	h := &temp
	heap.Init(h)

	if k == 1 {
		return (*h)[0]
	}

	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}

	return (*h)[0]
}

// highHeap 实现了 heap 的接口
type highHeap []int

func (h highHeap) Len() int {
	return len(h)
}
func (h highHeap) Less(i, j int) bool {
	// h[i] > h[j] 使得 h[0] == max(h...)
	return h[i] > h[j]
}
func (h highHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *highHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

func findKthLargest2(nums []int, k int) int {
	return find(nums, 0, len(nums)-1, k)
}

func find(nums []int, low, high, k int) int {

	start, end := low, high
	tmp := nums[low]

	for low < high {

		for low < high && nums[high] <= tmp {
			high--
		}

		nums[low] = nums[high]

		for low < high && nums[low] >= tmp {
			low++
		}
		nums[high] = nums[low]

	}
	if low == k-1 {
		return tmp
	}
	if low > k-1 {
		return find(nums, start, low-1, k)
	}
	return find(nums, low+1, end, k)

}
