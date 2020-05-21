package algo

import (
	"fmt"
)

func Selection(data []int64) []int64 {
	for i := 0; i < len(data); i++ {
		tmpIndex := i

		for j := i + 1; j < len(data); j++ {
			if data[tmpIndex] > data[j] {
				tmpIndex = j
			}
		}

		data[i], data[tmpIndex] = data[tmpIndex], data[i]
	}

	return data
}

func Insert(data []int64) []int64 {
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		j := i - 1
		for ; j >= 0; j-- {
			if data[j] > tmp {
				data[j+1] = data[j]
			} else {
				break
			}
		}

		data[j+1] = tmp
		fmt.Println(data)
	}
	return data
}

func Bubble(data []int64) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println(data)
}

func Merge(data []int64) []int64 {
	mid := len(data) / 2
	fmt.Println(data)
	if mid > 0 {
		fmt.Println(mid)
		data1 := Merge(data[:mid])
		data2 := Merge(data[mid:])
		fmt.Println(mid, data1, data2)
		merge(data1, data2)
	}

	return data
}

func merge(data1 []int64, data2 []int64) []int64 {
	res := []int64{}
	i := 0
	j := 0

	for {
		if i > len(data1) {
			res = append(res, data2[j:]...)
			break
		}

		if j > len(data2) {
			res = append(res, data1[i:]...)
			break
		}

		if data1[i] > data2[j] {
			res = append(res, data2[j])
			j++
		}

		if data1[i] <= data2[j] {
			res = append(res, data1[i])
			i++
		}

	}
	return res
}

func QuickSort(data []int64) {
	fmt.Println(quickSort(data))
}

func quickSort(data []int64) []int64 {
	mid := 0

	data, mid = partition(data)
	fmt.Println(data, mid)
	if (mid - 1) > 0 {
		quickSort(data[:mid-1])
	}

	fmt.Println(data[mid:])
	if len(data) > (mid - 1) {
		quickSort(data[mid-1:])
	}

	return data
}

func partition(data []int64) ([]int64, int) {
	if len(data) <= 1 {
		return data, 0
	}

	i := 0
	j := 0
	tmp := data[len(data)-1]

	for {
		for ; i < len(data); i++ {
			if data[i] > tmp {
				break
			}
		}

		for ; j < len(data); j++ {
			if data[j] <= tmp && j > i {
				break
			}
		}

		if i > len(data)-1 {
			break
		}

		if j > len(data)-1 {
			break
		}

		data[i], data[j] = data[j], data[i]

	}

	return data, i
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

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	qsort(head, nil)
	return head
}

// [head, tail)
func qsort(head, tail *ListNode) {
	if head == tail {
		return
	}
	pivot := partition2(head, tail)
	qsort(head, pivot)
	qsort(pivot.Next, tail)
}

// [head, tail)
func partition2(head, tail *ListNode) *ListNode {
	pivot := head.Val
	p := head
	q := head.Next
	//for q != tail {
	for q != nil {
		if q.Val < pivot {
			p = p.Next
			p.Val, q.Val = q.Val, p.Val
		}
		q = q.Next
	}
	head.Val, p.Val = p.Val, head.Val
	return p
}
