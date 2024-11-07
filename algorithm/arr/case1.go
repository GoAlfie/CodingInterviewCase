package arr

func hIndex(citations []int) int {
	for i := len(citations); i > 0; i-- {
		found := 0
		for _, v := range citations {
			if v >= i {
				found += 1
			}
		}
		if found >= i {
			return i
		}
	}
	return 0
}

// 去除重复元素
func removeDuplicateItem(list []any) []any {
	if len(list) == 0 {
		return list
	}

	i := 0
	// [0,i]是保证没有重复元素的列表
	for j := 1; j < len(list); j++ {
		if list[i] != list[j] {
			i++ // 必须是先 i++ 再赋值，把list[j]放在i的下一个位置
			list[i] = list[j]
		}
	}

	return list[:i+1]
}

// 中心下标，下标左侧元素的和等于右侧元素的和
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	// 开始的时候i在最左侧，lsum = 0, rsum = 列表的和
	lsum := 0
	rsum := 0
	for _, num := range nums {
		rsum += num
	}

	// 开始检查i的左侧元素和是否等于i的右侧元素和，如果相等，返回i
	for i, num := range nums {
		lsum += num
		if lsum == rsum {
			return i
		}
		rsum -= num
	}
	return -1
}
