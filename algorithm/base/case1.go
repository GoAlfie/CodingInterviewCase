package base

import "fmt"

// 求素数的个数，枚举法
func primeCount1(n int) int {
	count := 0
	isPrime := func(i int) bool {
		// j < 根号i即可，所以 j * j <= i
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				return false
			}
		}
		return true
	}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// 求素数的个数，过滤排除法
func primeCount2(n int) int {
	count := 0
	// 列表位标记是否是素数，不用于表示最终结果
	isPrime := make([]bool, n+1) // false 表示是素数
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			for j := i; j*i <= n; j++ {
				isPrime[j*i] = true
			}
			count++
		}
	}
	fmt.Printf("isPrime: %v\n", isPrime)
	return count
}

// 滑动窗口
// func lengthOfLongestSubstring(arr []int) {
// 	window := []int{}
// 	left := 0

// 	for right := 0; right < len(arr); right++ {
// 		window = append(window, arr[right])
// 		if !condition(window) {
// 			window = window[1:]
// 			left++
// 		}
// 	}
// }