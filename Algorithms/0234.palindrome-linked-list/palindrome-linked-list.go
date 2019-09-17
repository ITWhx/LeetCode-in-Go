package problem0234

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func isPalindrome(head *ListNode) bool {
	//// 获取 list 中的值
	//nums := make([]int, 0, 64)
	//for head != nil {
	//	nums = append(nums, head.Val)
	//	head = head.Next
	//}
	//
	//// 按照规则对比值
	//l, r := 0, len(nums)-1
	//for l < r {
	//	if nums[l] != nums[r] {
	//		return false
	//	}
	//	l++
	//	r--
	//}
	//return true
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	var pre *ListNode = nil
	for head != slow {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	if fast != nil {
		slow = slow.Next
	}

	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}
