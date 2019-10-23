/**
*
*@author 吴昊轩
*@create 2019-10-0721:13
 */
package heap_sort

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ints := []int{3, 42, 2, 5, 7, 32, 4, 4, 1}
	fmt.Println(ints)
	heapSort(ints)
	fmt.Println(ints)
}

//在n个数值中找出m个最大的数
func Test2(t *testing.T) {
	ints := []int{232, 3, 2, 4, 1, 3, 4, 123, 324, 4534, 123, 5, 63, 2}
	nums := findKNums(ints, 7)
	fmt.Println(nums)
}
