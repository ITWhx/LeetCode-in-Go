package problem0002

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

type question struct {
	p para
	a ans
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	substring := lengthOfLongestSubstring("abba")
	fmt.Println(substring)
	qs := []question{
		question{
			p: para{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
		question{
			p: para{
				one: makeListNode([]int{9, 8, 7, 6, 5}),
				two: makeListNode([]int{1, 1, 2, 3, 4}),
			},
			a: ans{
				one: makeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		question{
			p: para{
				one: makeListNode([]int{0}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers(p.one, p.two), "输入:%v", p)
	}
}

func lengthOfLongestSubstring(s string) int {
	start, maxLen := 0, 0

	table := [256]int{}

	for i, _ := range table {
		table[i] = -1
	}

	for i, v := range s {
		//注意这里是v，不是下标i
		if table[v] > -1 {
			start = table[v] + 1
		}
		table[v] = i
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
