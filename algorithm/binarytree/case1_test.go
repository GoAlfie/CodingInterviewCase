package binarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initTree() *TreeNode {
	head := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}

	head.Left = n2
	head.Right = n3

	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7

	return head
}

func initSearchTree() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}

	n4.Left = n2
	n4.Right = n6

	n2.Left = n1
	n2.Right = n3

	n6.Left = n5
	n6.Right = n7

	return n4
}

func Test_BuildBinaryTreeByPre(t *testing.T) {
	prelist := []int{1, 2, 4, 5, 3, 6, 7}
	inlist := []int{4, 2, 5, 1, 6, 3, 7}
	res1 := BuildBinaryTreeByPre1(prelist, inlist)
	assert.Equal(t, res1.Val, 1)

	res2 := BuildBinaryTreeByPre2(prelist, inlist)
	assert.Equal(t, res2.Val, 1)
	assert.Equal(t, res1, res2)

	res3 := PrePrint1(res1)
	fmt.Println(res3)
	assert.Equal(t, res3, prelist)

	res4 := MidPrint1(res1)
	fmt.Println(res4)
	assert.Equal(t, res4, inlist)
}

func Test_BuildBinaryTreeByPost(t *testing.T) {
	postlist := []int{4, 5, 2, 6, 7, 3, 1}
	inlist := []int{4, 2, 5, 1, 6, 3, 7}

	res1 := BuildBinaryTreeByPost(postlist, inlist)
	assert.Equal(t, res1.Val, 1)

	res2 := PostPrint1(res1)
	fmt.Println(res2)
	assert.Equal(t, res2, postlist)

	res3 := MidPrint1(res1)
	fmt.Println(res3)
	assert.Equal(t, res3, inlist)
}

func Test_PreOrder(t *testing.T) {
	head := initTree()
	res1 := PrePrint1(head)
	fmt.Println(res1)
	res2 := PrePrint2(head)
	fmt.Println(res2)

	assert.Equal(t, res1, res2)
}

func Test_InOrder(t *testing.T) {
	head := initTree()
	res1 := MidPrint1(head)
	fmt.Println(res1)
	res2 := MidPrint2(head)
	fmt.Println(res2)

	assert.Equal(t, res1, res2)
}

func Test_PostOrder(t *testing.T) {
	head := initTree()
	res1 := PostPrint1(head)
	fmt.Println(res1)
	res2 := PostPrint2(head)
	fmt.Println(res2)

	assert.Equal(t, res1, res2)
}
