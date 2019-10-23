/**
*
*@author 吴昊轩
*@create 2019-10-0720:58
 */
package heap_sort

func heapSort(nums []int) {
	len := len(nums) - 1
	//构建大顶堆，只需调整一半的节点。len/2-1下标之前的节点才有子节点，
	for i := len/2 - 1; i >= 0; i-- {
		buildHeap(nums, i, len)
	}

	for j := len; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		buildHeap(nums, 0, j-1)
	}
}

//大顶堆
func buildHeap(nums []int, i, n int) {
	for {
		//左孩子节点
		j1 := i*2 + 1
		//j1 <0 int数溢出
		if j1 > n || j1 < 0 {
			break
		}
		j := j1
		//右孩子大于左孩子，将j2赋值j
		if j2 := j1 + 1; j2 <= n && nums[j2] > nums[j1] {
			j = j2
		}
		//当前节点大于左右孩子，则退出循环 这一步上次忘记了！！！！！！！！！！！！！！！！！！
		if nums[i] > nums[j] {
			break
		}
		//交换
		nums[i], nums[j] = nums[j], nums[i]
		//交换节点下标赋值当前下标，继续向下构建
		i = j
	}
}

//构建小顶堆
func buildHeap1(nums []int, i, n int) {
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

//从m个数中找到最大的n个数
func findKNums(nums []int, n int) []int {

	target := nums[:n]
	l := len(target) - 1
	for i := l/2 - 1; i >= 0; i-- {
		buildHeap1(target, i, l)
	}

	for i := n; i < len(nums); i++ {
		if nums[i] > target[0] {
			target[0] = nums[i]
			buildHeap1(target, 0, l)
		}
	}
	return target
}
