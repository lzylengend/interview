package algo

func Fib(N int) int {
	if N <= 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	return Fib(N-1) + Fib(N-2)
}

func ClimbStairs(n int) int {
	dp := []int{}

	dp = append(dp, 1)
	dp = append(dp, 2)

	for i := 2; i < n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n-1]
}

func subsets(nums []int) [][]int {

}

func subsets1(nums []int) [][]int {
	res := make([][]int, 0)
	for _, v := range nums {

	}
}
