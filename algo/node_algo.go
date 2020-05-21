package algo

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this ListNode) String() string {
	s := ""
	head := &this

	for {
		s = s + strconv.Itoa(head.Val) + ","
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
	s = s + "\n"

	return s
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	current := head

	if head.Next == nil {
		return nil
	}

	delCurrent := head.Next
	currentIndex := 0

	for {
		if currentIndex < n {
			currentIndex = currentIndex + 1
		} else {
			current = current.Next
		}

		if delCurrent.Next == nil {
			if currentIndex < n {
				head = head.Next
			} else {
				current.Next = current.Next.Next
			}

			break
		}

		delCurrent = delCurrent.Next

	}
	return head
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Val: head.Val}

	if head.Next == nil {
		return newHead
	}
	currentNode := head.Next

	for {
		newNextNode := &ListNode{Val: currentNode.Val}
		newNextNode.Next = newHead
		newHead = newNextNode

		if currentNode.Next != nil {
			currentNode = currentNode.Next
		} else {
			break
		}

	}
	return newHead
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//current := l1
	//current2 := l2
	//head:=l1
	//
	//for {
	//	if current.Val < current2.Val {
	//		if current.Next != nil {
	//			nextNode = current.Next
	//		}
	//
	//		current.Next = current2.Next
	//	} else {
	//		if current2.Next != nil {
	//			nextNode = current2.Next
	//		}
	//
	//		current2.Next = current.Next
	//	}
	//}
	return nil
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	one := head
	two := head

	for {
		if two.Next == nil || two.Next.Next == nil {
			return false
		}

		if one.Next == nil {
			return false
		}

		one = one.Next
		two = two.Next.Next

		if one.Next == two.Next {
			return true
		}

	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	current := l3
	carry := 0

	for {
		l1V := 0
		l2V := 0

		if l1 != nil {
			l1V = l1.Val
		}

		if l2 != nil {
			l2V = l2.Val
		}

		tmp := l1V + l2V + carry

		if tmp < 10 {
			carry = 0
		} else {
			tmp = tmp % 10
			carry = 1
		}

		current.Val = tmp
		if l1 == nil || l1.Next == nil {
			l1 = nil
		} else {
			l1 = l1.Next
		}

		if l2 == nil || l2.Next == nil {
			l2 = nil
		} else {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if carry >= 1 {
				current.Next = &ListNode{Val: carry}
			}
			return l3
		}

		current.Next = &ListNode{}
		current = current.Next
	}
}
