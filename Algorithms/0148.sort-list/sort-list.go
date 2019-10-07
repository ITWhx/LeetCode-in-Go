package problem0148

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode 是题目预定义的数据类型
type ListNode = kit.ListNode

//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	left, right := split(head)
//	return merge(sortList(left), sortList(right))
//}
//
//// 从中间位置，切分 list
//func split(head *ListNode) (left, right *ListNode) {
//	// head.Next != nil
//	// 因为， sortList 已经帮忙检查过了
//
//	// fast 的变化速度是 slow 的两倍
//	// 当 fast 到达末尾的时候，slow 正好在 list 的中间
//	slow, fast := head, head
//	var slowPre *ListNode
//
//	for fast != nil && fast.Next != nil {
//		slowPre, slow = slow, slow.Next
//		fast = fast.Next.Next
//	}
//
//	// 斩断 list
//	slowPre.Next = nil
//	// 使用 slowPre 是为了确保当 list 的长度为 2 时，会均分为两个长度为 1 的子 list
//
//	left, right = head, slow
//	return
//}
//
//// 把已经排序好了的两个 list left 和 right
//// 进行合并
//func merge(left, right *ListNode) *ListNode {
//	// left != nil , right != nil
//	// 因为， sortList 已经帮忙检查过了
//
//	cur := &ListNode{}
//	headPre := cur
//	for left != nil && right != nil {
//		// cur 总是接上较小的 node
//		if left.Val < right.Val {
//			cur.Next, left = left, left.Next
//		} else {
//			cur.Next, right = right, right.Next
//		}
//		cur = cur.Next
//	}
//
//	// 处理 left 或 right 中，剩下的节点
//	if left == nil {
//		cur.Next = right
//	} else {
//		cur.Next = left
//	}
//
//	return headPre.Next
//}

func sortList(head *ListNode) *ListNode {
	// 如果 head为空或者head就一位,直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// 定义快慢俩指针,当快指针到末尾的时候,慢指针肯定在链表中间位置
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// 把链表拆分成两段,所以设置中间位置即慢指针的next为nil
	n := slow.Next
	slow.Next = nil
	// 递归排序
	return merge(sortList(head), sortList(n))
}

func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	// 设置一个空链表,
	node := &ListNode{Val: 0}
	current := node
	// 挨个比较俩链表的值,把小的值放到新定义的链表里,排好序
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			current.Next, node1 = node1, node1.Next
		} else {
			current.Next, node2 = node2, node2.Next
		}
		current = current.Next
	}

	// 两链表可能有一个没走完,所以要把没走完的放到链表的后面
	// 注意,此处跟 数组不一样的是, 数组为什么要循环,因为数组可能一个数组全部走了(比如 12345与6789比较, 前面的全部走完,后面一个没走),另一个可能有多个没走..
	// 链表虽然也有这种可能,但是 node1和node2已经是有序的了,如果另外一个没有走完,直接把next指向node1或者node2就行,因为这是链表
	if node1 != nil {
		current.Next, node1 = node1, node1.Next
	}
	if node2 != nil {
		current.Next, node2 = node2, node2.Next
	}
	return node.Next
}
