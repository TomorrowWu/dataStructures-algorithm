package src

//转化为求两个字符串最大公共子序列的问题
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)

	// dp[i][j] == k 表示 word1[:i] 和 word2[:j] 的最大公共子序列的长度为 k
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// 总长度-公共子字符串长度
	return n1 + n2 - dp[n1][n2]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
