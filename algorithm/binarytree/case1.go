package binarytree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode

	Val int
}

// 前序、中序构建二叉树1
func BuildBinaryTreeByPre1(prelist, inlist []int) *TreeNode {
	if len(prelist) == 0 || len(inlist) == 0 {
		return nil
	}

	// 根据前序确定头节点，根据中序确定左右节点
	var buildBinaryTreeByPre1 func(prelist []int, prestart, preend int, inlist []int, instart, inend int) *TreeNode
	buildBinaryTreeByPre1 = func(prelist []int, prestart, preend int, inlist []int, instart, inend int) *TreeNode {
		if prestart > preend || instart > inend {
			return nil
		}
		head := &TreeNode{Val: prelist[prestart]}

		// 找到头节点在中序中的索引
		var headIndex int
		for i, v := range inlist {
			if v == head.Val {
				headIndex = i
				break
			}
		}

		leftCount := headIndex - instart
		// rightCount := inend - headIndex

		head.Left = buildBinaryTreeByPre1(prelist, prestart+1, prestart+leftCount, inlist, instart, headIndex-1)
		head.Right = buildBinaryTreeByPre1(prelist, prestart+leftCount+1, preend, inlist, headIndex+1, inend)

		return head
	}

	return buildBinaryTreeByPre1(prelist, 0, len(prelist)-1, inlist, 0, len(inlist)-1)
}

// 前序、中序构建二叉树2
func BuildBinaryTreeByPre2(prelist, inlist []int) *TreeNode {
	if len(prelist) == 0 || len(inlist) == 0 {
		return nil
	}

	head := &TreeNode{Val: prelist[0]}
	// 找到头节点在中序中的索引
	var headIndex int
	for i, v := range inlist {
		if v == head.Val {
			headIndex = i
			break
		}
	}

	head.Left = BuildBinaryTreeByPre2(prelist[1:headIndex+1], inlist[:headIndex])
	head.Right = BuildBinaryTreeByPre2(prelist[headIndex+1:], inlist[headIndex+1:])

	return head
}

// 后序、中序构建二叉树
func BuildBinaryTreeByPost(postlist, inlist []int) *TreeNode {
	if len(postlist) == 0 || len(inlist) == 0 {
		return nil
	}

	head := &TreeNode{Val: postlist[len(postlist)-1]}
	var headIndex int
	for i, v := range inlist {
		if v == head.Val {
			headIndex = i
			break
		}
	}

	head.Left = BuildBinaryTreeByPost(postlist[:headIndex], inlist[:headIndex])
	head.Right = BuildBinaryTreeByPost(postlist[headIndex:len(postlist)-1], inlist[headIndex+1:])

	return head
}

// 前序1
func PrePrint1(node *TreeNode) (res []int) {
	if node == nil {
		return
	}
	res = append(res, node.Val)
	res = append(res, PrePrint1(node.Left)...)
	res = append(res, PrePrint1(node.Right)...)

	return
}

// 前序2
func PrePrint2(node *TreeNode) (res []int) {
	if node == nil {
		return
	}
	stack := []*TreeNode{node}
	for len(stack) != 0 {
		// 弹出
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, n.Val)
		// 左右节点依次入栈
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
	}

	return
}

// 中序1
func MidPrint1(node *TreeNode) (res []int) {
	if node == nil {
		return
	}
	res = append(res, MidPrint1(node.Left)...)
	res = append(res, node.Val)
	res = append(res, MidPrint1(node.Right)...)

	return
}

// 中序2
func MidPrint2(node *TreeNode) (res []int) {
	if node == nil {
		return
	}
	stack := []*TreeNode{}
	// 用一个变量记录当前的位置
	cur := node

	for len(stack) != 0 || cur != nil {
		// 只要左侧不为空就往左侧走
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 弹出
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, n.Val)

		if n.Right != nil {
			cur = n.Right
		}
	}

	return
}

// 后序1
func PostPrint1(node *TreeNode) (res []int) {
	if node == nil {
		return
	}
	res = append(res, PostPrint1(node.Left)...)
	res = append(res, PostPrint1(node.Right)...)
	res = append(res, node.Val)

	return
}

// 后序2
func PostPrint2(node *TreeNode) (res []int) {
	if node == nil {
		return
	}
	stack := []*TreeNode{node}

	// 前序:根—>左->右, 后序:左->右->根,返过来就是:根—>右->左
	for len(stack) != 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, n.Val)

		if n.Left != nil {
			stack = append(stack, n.Left)
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i += 1
		j -= 1
	}

	return
}

// 层序遍历
func bfs(head *TreeNode) (res []int) {
	if head == nil {
		return
	}
	queue := []*TreeNode{head}

	for len(queue) != 0 {
		// 队列弹出
		n := queue[0]
		queue = queue[1:]

		res = append(res, n.Val)
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		}
	}
	return
}

type queueItem struct {
	*TreeNode
	deep int
}

// 层序遍历,分别输出每一层
func bfs2(head *TreeNode) (res [][]int) {
	if head == nil {
		return
	}
	// 每个节点记录节点所在的层
	queue := []*queueItem{{head, 0}}
	// 记录每一层的结果
	path := []int{}
	// 记录当前所在的层
	curdeep := 0

	for len(queue) != 0 {
		// 队列弹出
		n := queue[0]
		queue = queue[1:]

		if n.deep > curdeep {
			// 收集结果path
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)

			// 重置变量
			curdeep = n.deep
			path = []int{}
		}

		path = append(path, n.Val)

		if n.Left != nil {
			queue = append(queue, &queueItem{n.Left, n.deep + 1})
		}
		if n.Right != nil {
			queue = append(queue, &queueItem{n.Right, n.deep + 1})
		}
	}
	// 结果加入最后一层
	res = append(res, path)

	return
}
