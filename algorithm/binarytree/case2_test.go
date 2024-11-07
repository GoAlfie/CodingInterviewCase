package binarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initSymmetricTree() *TreeNode {
	n1 := &TreeNode{Val: 1}
	n2_l := &TreeNode{Val: 2}
	n2_r := &TreeNode{Val: 2}
	n3_l := &TreeNode{Val: 3}
	n4_l := &TreeNode{Val: 4}
	n3_r := &TreeNode{Val: 3}
	n4_r := &TreeNode{Val: 4}

	n1.Left = n2_l
	n1.Right = n2_r

	n2_l.Left = n4_l
	n2_l.Right = n3_l

	n2_r.Left = n3_r
	n2_r.Right = n4_r

	return n1
}

func Test_IsBalanceTree(t *testing.T) {
	head := initTree()
	deep, balance := isBalanceTree(head)
	fmt.Println(deep, balance)
	assert.Equal(t, balance, true)

	head2 := initSearchTree()
	deep2, balance2 := isBalanceTree(head2)
	fmt.Println(deep2, balance2)
	assert.Equal(t, balance2, true)
}

func Test_IsSearchTree(t *testing.T) {
	head1 := initTree()
	_, _, search1 := isSearchTree(head1)
	assert.Equal(t, search1, false)

	search2 := isSearchTree2(head1, nil, nil)
	assert.Equal(t, search2, false)

	assert.Equal(t, search1, search2)

	head2 := initSearchTree()
	max, min, search3 := isSearchTree(head2)
	fmt.Println(max, min, search2)
	assert.Equal(t, search3, true)

	search4 := isSearchTree2(head2, nil, nil)
	assert.Equal(t, search4, true)

	assert.Equal(t, search3, search4)
}

func Test_IsSameTree(t *testing.T) {
	head1 := initTree()
	head2 := initSearchTree()
	head3 := initTree()

	res1 := isSameTree(head1, head1)
	assert.Equal(t, res1, true)

	res2 := isSameTree(head1, head2)
	assert.Equal(t, res2, false)

	res3 := isSameTree(head1, head3)
	assert.Equal(t, res3, true)
}

func Test_IsSymmetric(t *testing.T) {
	head1 := initTree()
	res1 := isSymmetric(head1)
	assert.Equal(t, res1, false)

	head2 := initSymmetricTree()
	res2 := isSymmetric(head2)
	assert.Equal(t, res2, true)
}
