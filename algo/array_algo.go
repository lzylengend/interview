package algo

import (
	"fmt"
)

func MergeNums(nums1 []int, m int, nums2 []int, n int) {
	end1 := len(nums1) - 1
	m1 := m - 1
	n1 := n - 1

	if m1 < 0 {
		for i, _ := range nums2 {
			nums1[i] = nums2[i]
		}
		return
	}

	for {
		if n1 < 0 {
			break
		}

		if m1 < 0 {
			if end1 < 0 {
				break
			}

			for i := end1; i >= 0; i-- {
				nums1[i] = nums2[n1]
				n1--
			}
			break
		}

		if nums1[m1] > nums2[n1] {
			nums1[end1] = nums1[m1]
			m1--

		} else {
			nums1[end1] = nums2[n1]
			n1--
		}

		end1--
	}

	fmt.Println(nums1)
}

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		if DicSearch(matrix[i], target) {
			return true
		}
	}

	return false
}

func DicSearch(data []int, target int) bool {
	mid := len(data) / 2

	for {
		if target == data[mid] {
			return true
		}

		if mid >= len(data)-1 || mid == 0 {
			break
		}

		if target > data[mid] {
			mid = (mid + len(data)) / 2
		} else {
			mid = mid / 2
		}

	}

	return false
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i1 := 0
	i2 := 0

	length := len(nums1) + len(nums2)
}
