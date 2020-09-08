package leetcode

func nthUglyNumber(n int) int {
	// 初始化三个下标
	a, b, c := 0, 0, 0

	dp := make([]int, n)
	// 丑数 1 加入
	dp[0] = 1

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5

		// 将当前 dp 位置设置为最小的数
		dp[i] = n2
		if dp[i] > n3 {
			dp[i] = n3
		}
		if dp[i] > n5 {
			dp[i] = n5
		}

		// 并将选中的数字的下标 +1
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}
