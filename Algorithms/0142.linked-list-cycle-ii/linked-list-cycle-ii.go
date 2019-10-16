package problem0142

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

//设: 1. 从头节点到入环节点长度为a
//2. 从入环的节点到相遇节点长度为b
//3. 从相遇的节点再到达入环节点的长度为c
//证明: a = c(解法1的前提) 第一次相遇时fast超了slow一圈环的长度, 也就是b+c, 且
//fast的速度是slow的2倍 又因为slow走过的距离是a+b, 所以2(a+b) = a+b + b+c, 得
//到a=c 由a=c, 我们可以得到:使指针从环内相遇的的节点和头节点同时出发, 最终会
//在入口节点会合
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	//想不明白这里为什么从第二，三个节点开始（其实快慢指针是从头节点开始移动，先各子移动一次，避免第一次slow== fast导致循环跳出）
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}

	if slow != fast {
		return nil
	}

	for slow != head {
		slow, head = slow.Next, head.Next
	}
	return slow
}
