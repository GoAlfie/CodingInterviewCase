package linkedlist

import "fmt"

// 构建一个链表
func InitList(start, end int) *ListNode {
	pre := &ListNode{}
	cur := pre
	if start < end {
		for i := start; i <= end; i++ {
			cur.Next = &ListNode{Val: i}
			cur = cur.Next
		}
		return pre.Next
	}
	for i := start; i >= end; i-- {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	return pre.Next
}

// 打印链表
func ListPrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	fmt.Printf("List: %v\n", res)
	return res
}

// 反转链表, 迭代方式
func RevertList1(head *ListNode) *ListNode {
	var (
		pre  *ListNode
		next *ListNode
	)
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// 反转链表，递归方式，递归到最后一个节点再修改指向
func RevertList2(head *ListNode) *ListNode {
	// 退出条件，这里一定要head.Next == nil，因为我们要设置下一个节点指向当前节点
	if head == nil || head.Next == nil {
		return head
	}
	// 先递归到最后一个节点再修改指向
	res := RevertList2(head.Next)

	// 下一个节点的Next指向当前节点
	head.Next.Next = head
	// 当前节点指向nil
	head.Next = nil

	return res
}

// 反转链表，递归方式，先修改指向再递归
func RevertList3(head *ListNode) *ListNode {

	var revert func(pre, head *ListNode) *ListNode
	revert = func(pre, head *ListNode) *ListNode {
		// 退出调整，这里一定要head.Next == nil，因为我们要设置head.Next
		if head == nil || head.Next == nil {
			head.Next = pre
			return head
		}
		next := head.Next
		head.Next = pre
		return revert(head, next)
	}

	return revert(nil, head)
}

// 判断是否有环, 先走再判断是否相同
func IsHasRing1(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 判断是否有环, 先判断是否相同再走
func IsHasRing2(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 找到环入口
func RingNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 判断环的长度
func RingLong(head *ListNode) int {
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	res := 1
	slow = slow.Next
	fast = fast.Next.Next
	for fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
		res += 1
	}
	return res
}

// 判断两个链表是否相交
func HasSameNode(h1, h2 *ListNode) bool {
	for h1.Next != nil {
		h1 = h1.Next
	}
	for h2.Next != nil {
		h2 = h2.Next
	}
	return h1 == h2
}

// 判断两个链表相交的节点
func SameNode(h1, h2 *ListNode) *ListNode {
	cur1,cur2 := h1,h2
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
		if cur1 == nil {
			cur1 = h2
		}
		if cur2 == nil {
			cur2 = h1
		}
	}
	return cur1
}
