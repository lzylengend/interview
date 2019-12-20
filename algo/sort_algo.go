package algo

import (
	"fmt"
)

func Selection(data []int64) {

}

func ClimbStairs(n int) int {
	a := 1
	b := 2
	tmp := a + b

	if n < 1 {
		return 0
	}

	if n == 1 {
		return a
	}

	if n == 2 {
		return b
	}

	for i := 3; i <= n; i++ {
		tmp = a + b
		a, b = b, tmp
	}

	return tmp
}

func MaxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	count := 0

	for i := 0; i < len(prices); i++ {
		if i+1 == len(prices) {
			return count
		}

		if prices[i+1] < prices[i] {
			continue
		}

		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > count {
				count = prices[j] - prices[i]
			}
		}
	}

	return count
}

func MaxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	maxCount := nums[0]
	maxTmp := nums[0]

	for i := 1; i < len(nums); i++ {
		if maxTmp <= 0 {
			maxTmp = nums[i]
		} else {
			maxTmp += nums[i]
		}

		if maxTmp > maxCount {
			maxCount = maxTmp
		}
	}

	return maxCount
}

func Rob(nums []int) int {
	return 0
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 1
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != current {
			nums[index] = nums[i]
			index++
			current = nums[i]
		}
	}

	fmt.Println(nums)

	return index
}

func MaxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	isBuy := false
	buy := prices[0]
	count := 0

	for i := 0; i < len(prices); i++ {
		if isBuy && prices[i] > buy {
			count = count + (prices[i] - buy)
			buy = prices[i]
			isBuy = false
		}

		if i+1 == len(prices) {
			return count
		}

		if !isBuy && prices[i+1] > prices[i] {
			isBuy = true
			buy = prices[i]
		}

	}

	return count
}

func Rotate(nums []int, k int) {
	for i := 0; i < len(nums); i++ {

	}

}

func ContainsDuplicate(nums []int) bool {
	tmp := make(map[int]bool)
	for _, v := range nums {
		_, ok := tmp[v]
		if ok {
			return true
		}

		tmp[v] = false
	}

	return false
}

func SingleNumber(nums []int) int {
	t := 0

	for _, v := range nums {
		t = t ^ v
	}

	return t
}

func MoveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}

func PlusOne(digits []int) []int {
	tmp := 1

	for i := len(digits) - 1; i >= 0; i-- {
		if tmp != 1 {
			break
		}

		digits[i] = digits[i] + 1
		if digits[i] >= 10 {
			tmp = 1
			digits[i] = digits[i] - 10
		} else {
			tmp = 0
		}
	}

	if tmp == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i, v := range nums {
		m[v] = i
	}

	res := make([]int, 0)
	for i, v := range nums {
		_, ok := m[target-v]
		if ok {
			if i == m[target-v] {
				continue
			}

			res = append(res, i, m[target-v])
			break
		}
	}

	return res
}
