package problem0215

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  int
}{

	{
		[]int{3, 3, 3, 3, 3, 3, 3, 3, 3},
		1,
		3,
	},

	{
		[]int{3, 2, 1, 5, 6, 4, 3, 5, 23, 6, 2, 4, 6, 5, 2, 4, 6, 1, 3, 43, 6, 2, 4, 6, 3, 1, 8, 6, 3, 4, 5, 23, 767, 52, 3, 4, 5, 23, 1, 352, 34, 5, 12, 354, 6, 56, 4, 2, 3, 4, 7, 54, 4, 23, 4, 6, 745, 6},
		2,
		5,
	},

	// 可以有多个 testcase
}

func Test_findKthLargest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findKthLargest(tc.nums, tc.k), "输入:%v", tc)
	}
}
func Test_heap(t *testing.T) {
	ast := assert.New(t)

	h := new(highHeap)

	i := 5
	h.Push(i)
	ast.Equal(i, h.Pop(), "Pop() after Push(%d)", i)

}
func Benchmark_findKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			//findKthLargest(tc.nums, tc.k)
			findKthLargest2(tc.nums, tc.k)
		}
	}
}
