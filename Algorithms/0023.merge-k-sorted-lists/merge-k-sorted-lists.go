package problem0023

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(lists []*ListNode) *ListNode {
	length := len(lists)
	half := length / 2

	if length == 1 {
		return lists[0]
	}

	if length == 2 {
		var (
			l0, l1   = lists[0], lists[1]
			res, cur *ListNode
		)

		if l0 == nil {
			return l1
		}
		if l1 == nil {
			return l0
		}

		if l0.Val < l1.Val {
			res, cur, l0 = l0, l0, l0.Next
		} else {
			res, cur, l1 = l1, l1, l1.Next
		}

		for l0 != nil && l1 != nil {
			if l0.Val < l1.Val {
				cur.Next, l0 = l0, l0.Next
			} else {
				cur.Next, l1 = l1, l1.Next
			}
			cur = cur.Next
		}

		if l0 != nil {
			cur.Next = l0
		}
		if l1 != nil {
			cur.Next = l1
		}

		return res
	}

	return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func merge(list []*ListNode) *ListNode {
	length := len(list)
	half := length / 2

	if length == 1 {
		return list[0]
	}

	if length == 2 {
		l1, l2 := list[0], list[1]
		var res, cur *ListNode

		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}
		if l1.Val < l2.Val {
			res = l1
			cur = l1
			l1 = l1.Next
		} else {
			res = l2
			cur = l2
			l1 = l2.Next
		}

		for l1 != nil && l2 != nil {

			if l1.Val > l2.Val {
				cur.Next = l2
				l2 = l2.Next
			} else {
				cur.Next = l1
				l1 = l1.Next
			}
			cur = cur.Next
		}

		if l1 != nil {
			cur.Next = l1
		}
		if l2 != nil {
			cur.Next = l2
		}
		return res
	}

	return mergeKLists([]*ListNode{mergeKLists(list[half:]), mergeKLists(list[:half])})
}
