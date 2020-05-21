package algo

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	fmt.Println(ClimbStairs(3))
}

func TestMergeNums(t *testing.T) {
	MergeNums([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func TestDicSearch(t *testing.T) {
	fmt.Println(DicSearch([]int{5}, 5))
}

func TestSearchMatrix(t *testing.T) {
	fmt.Println(SearchMatrix([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, 19))
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
}

func TestIsSubtree(t *testing.T) {
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 4}
	r.Right = &TreeNode{Val: 5}
	r.Left.Left = &TreeNode{Val: 1}
	r.Left.Right = &TreeNode{Val: 2}

	t2 := &TreeNode{Val: 4}
	t2.Left = &TreeNode{Val: 1}
	t2.Right = &TreeNode{Val: 2}

	fmt.Println(IsSubtree(r, t2))
}

func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 8}
	//l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 0}
	//l2.Next = &ListNode{Val: 6}
	//l2.Next.Next = &ListNode{Val: 4}

	fmt.Println(addTwoNumbers(l1, l2))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
