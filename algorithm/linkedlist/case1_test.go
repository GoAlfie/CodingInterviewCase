package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResertList1(t *testing.T) {
	var (
		// head       = &ListNode{}
		expecthead = &ListNode{}
	)
	expecthead = InitList(5, 1)
	expectnum := ListPrint(expecthead)

	// 构造长度5的链表
	head1 := InitList(1, 5)
	// 打印链表
	_ = ListPrint(head1)

	// 反转链表，迭代
	res1 := RevertList1(head1)
	// 打印链表
	num1 := ListPrint(res1)

	assert.Equal(t, num1, expectnum)

	// 构造长度5的链表
	head2 := InitList(1, 5)
	// 打印链表
	_ = ListPrint(head2)

	// 反转链表，递归
	res2 := RevertList2(head2)
	num2 := ListPrint(res2)
	assert.Equal(t, num2, expectnum)

	// 构造长度5的链表
	head3 := InitList(1, 5)
	// 打印链表
	_ = ListPrint(head3)

	// 反转链表，递归
	res3 := RevertList3(head3)
	num3 := ListPrint(res3)
	assert.Equal(t, num3, expectnum)
}

func Test_IsHadRing(t *testing.T) {
	head := InitList(1, 5)
	_ = ListPrint(head)

	// 无环
	res1 := IsHasRing1(head)
	assert.Equal(t, res1, false)
	res11 := IsHasRing2(head)
	assert.Equal(t, res11, false)

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head.Next.Next

	// 有环
	res2 := IsHasRing1(head)
	assert.Equal(t, res2, true)
	res22 := IsHasRing2(head)
	assert.Equal(t, res22, true)
}

func Test_RingNode(t *testing.T) {
	head := InitList(1, 10)
	_ = ListPrint(head)

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head.Next.Next

	res1 := RingNode(head)
	assert.Equal(t, res1.Val, 3)

	cur.Next = head.Next
	res2 := RingNode(head)
	assert.Equal(t, res2.Val, 2)

	cur.Next = head.Next.Next.Next
	res3 := RingNode(head)
	assert.Equal(t, res3.Val, 4)
}

func Test_RingLong(t *testing.T) {
	head := InitList(1, 6)
	_ = ListPrint(head)

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = head.Next.Next
	res1 := RingLong(head)
	assert.Equal(t, res1, 6-2)

	cur.Next = head.Next
	res2 := RingLong(head)
	assert.Equal(t, res2, 6-1)
}

func Test_HasSameNode(t *testing.T) {
	head1 := InitList(2, 5)
	head2 := InitList(3, 7)
	res1 := HasSameNode(head1, head2)
	assert.Equal(t, res1, false)

	cur1 := head1
	for cur1.Next != nil {
		cur1 = cur1.Next
	}
	cur1.Next = head2.Next.Next

	res2 := HasSameNode(head1, head2)
	assert.Equal(t, res2, true)

	cur1.Next = head2
	res3 := HasSameNode(head1, head2)
	assert.Equal(t, res3, true)

	cur1.Next = head2.Next
	res4 := HasSameNode(head1, head2)
	assert.Equal(t, res4, true)
}

func Test_SameNode(t *testing.T) {
	head1 := InitList(2, 5)
	head2 := InitList(7, 10)

	cur1 := head1
	for cur1.Next != nil {
		cur1 = cur1.Next
	}
	cur1.Next = head2.Next.Next

	res1 := SameNode(head1, head2)
	assert.Equal(t, res1.Val, 9)

	cur1.Next = head2.Next
	res2 := SameNode(head1, head2)
	assert.Equal(t, res2.Val, 8)

	cur1.Next = head2.Next.Next.Next
	res3 := SameNode(head1, head2)
	assert.Equal(t, res3.Val, 10)

	cur1.Next = head2
	res4 := SameNode(head1, head2)
	assert.Equal(t, res4.Val, 7)

	res5 := SameNode(head1, head1)
	assert.Equal(t, res5.Val, 2)
}
