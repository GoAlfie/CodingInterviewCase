package binarytree

import (
	"math"
	"strconv"
)

// 判断是否是平衡树
func isBalanceTree(head *TreeNode) (deep int, balance bool) {
	if head == nil {
		return 0, true
	}
	ldeep, lbalance := isBalanceTree(head.Left)
	rdeep, rbalance := isBalanceTree(head.Right)

	if ldeep > rdeep {
		return ldeep + 1, ldeep-rdeep < 2 && lbalance && rbalance
	}
	return rdeep + 1, rdeep-ldeep < 2 && lbalance && rbalance
}

// 判断是否是搜索树1
func isSearchTree(head *TreeNode) (max, min int, search bool) {
	if head == nil {
		return math.MinInt, math.MaxInt, true
	}

	lmax, min, lsearch := isSearchTree(head.Left)
	max, rmin, rsearch := isSearchTree(head.Right)

	search = lmax < head.Val && rmin > head.Val && lsearch && rsearch

	if head.Left == nil {
		min = head.Val
	}
	if head.Right == nil {
		max = head.Val
	}

	return
}

// 判断是否是搜索树2
func isSearchTree2(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	// 设定每个子树的最大值与最小值
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	return isSearchTree2(node.Left, min, &node.Val) && isSearchTree2(node.Right, &node.Val, max)
}

// 判断是两个相同的树
func isSameTree(head1, head2 *TreeNode) bool {
	if head1 == nil {
		return head2 == nil
	}
	if head2 == nil {
		return head1 == nil
	}
	if head1.Val != head2.Val {
		return false
	}

	return isSameTree(head1.Left, head2.Left) && isSameTree(head1.Right, head2.Right)
}

// 判断是否是对称树
func isSymmetric(head *TreeNode) bool {
	if head == nil {
		return true
	}

	// 判断两个数相互对称
	var f func(head1, head2 *TreeNode) bool
	f = func(head1, head2 *TreeNode) bool {
		if head1 == nil || head2 == nil {
			return head1 == head2
		}
		if head1.Val != head2.Val {
			return false
		}

		return f(head1.Left, head2.Right) && f(head1.Right, head2.Left)
	}

	return f(head.Left, head.Right)
}

// 所有左叶子节点之和1
func SumOfLeftLeaves1(root *TreeNode) int {
	var f func(head *TreeNode, isLeft bool) int
	// 计算当前子树下所有左叶子节点之和
	f = func(head *TreeNode, isLeft bool) int {
		res := 0
		if head == nil {
			return res
		}
		if isLeft && head.Left == nil && head.Right == nil {
			res += head.Val
		}

		res += f(head.Left, true)
		res += f(head.Right, false)

		return res
	}
	return f(root.Left, true) + f(root.Right, false)
}

// 所有左叶子节点之和2
func SumOfLeftLeaves2(root *TreeNode) int {
	res := 0
	if root == nil {
		return 0
	}

	// 判断左节点是不是叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}

	// 递归累加左子树和右子树
	res += SumOfLeftLeaves2(root.Left)
	res += SumOfLeftLeaves2(root.Right)

	return res
}

// 二叉树节点种类数量
func NumColorTree(root *TreeNode) int {
	set := make(map[int]struct{})

	var f func(head *TreeNode)
	f = func(head *TreeNode) {
		if head == nil {
			return
		}
		set[head.Val] = struct{}{}
		f(head.Left)
		f(head.Right)
	}

	f(root)

	return len(set)
}

// 二叉树的所有路径 https://leetcode.cn/problems/binary-tree-paths
func BinaryTreePaths(root *TreeNode) []string {
	res := []string{}

	var dfs func(head *TreeNode, path string)
	dfs = func(head *TreeNode, path string) {
		if head == nil {
			return
		}
		// 拼接路径字符串
		path += strconv.Itoa(head.Val)
		// 到叶子节点收集结果
		if head.Left == nil && head.Right == nil {
			res = append(res, path)
		}
		path += "->"
		dfs(head.Left, path)
		dfs(head.Right, path)
	}

	dfs(root, "")
	return res
}

// 二叉树路径总和 https://leetcode.cn/problems/path-sum-ii
func BinaryTreePathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)

	var dfs func(head *TreeNode, targetSum int, path []int)
	dfs = func(head *TreeNode, targetSum int, path []int) {
		if head == nil {
			return
		}
		// 记录已经走过的路径
		path = append(path, head.Val)
		// 还需要的目标路径和减去当前值（这样可以省略一个变量记录已经走过的路径和）
		targetSum -= head.Val

		// 如果是叶子节点且满足路径总和要求，则收集结果
		if head.Left == nil && head.Right == nil && targetSum == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		dfs(head.Left, targetSum, path)
		dfs(head.Right, targetSum, path)
	}

	dfs(root, targetSum, []int{})

	return res
}

// 所有可能的二叉搜索树 https://leetcode.cn/problems/unique-binary-search-trees-ii
func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var f func(start, end int) []*TreeNode
	f = func(start, end int) []*TreeNode {
		var res = make([]*TreeNode, 0)

		if start > end {
			// 这里是大坑，一定要返回列表中nil值，才能进去收集结果的循环
			return append(res, nil)
		}

		for i := start; i <= end; i++ {
			left := f(start, i-1)
			right := f(i+1, end)

			for _, l := range left {
				for _, r := range right {
					head := &TreeNode{Val: i}
					head.Left = l
					head.Right = r

					res = append(res, head)
				}
			}
		}
		return res
	}

	return f(1, n)
}

// 二叉树锯齿型遍历 https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
func ZigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	type queueItem struct {
		item *TreeNode
		deep int
	}
	queue := []*queueItem{{root, 0}}
	path := []int{}
	curdeep := 0

	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		if n.deep > curdeep {
			if curdeep%2 == 1 {
				for i := 0; i < (len(path) / 2); i++ {
					path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
				}
			}
			res = append(res, path)
			curdeep = n.deep
			path = []int{}
		}

		path = append(path, n.item.Val)
		if n.item.Left != nil {
			queue = append(queue, &queueItem{n.item.Left, n.deep + 1})
		}
		if n.item.Right != nil {
			queue = append(queue, &queueItem{n.item.Right, n.deep + 1})
		}

	}

	if curdeep%2 == 1 {
		for i := 0; i < (len(path) / 2); i++ {
			path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
		}
	}
	res = append(res, path)

	return res
}

