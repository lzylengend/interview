package algo

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val = t2.Val + t1.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	if equiTree(s, t) {
		return true
	}

	if IsSubtree(s.Left, t) {
		return true
	}

	return IsSubtree(s.Right, t)
}

func equiTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	if !equiTree(s.Left, t.Left) {
		return false
	}
	if !equiTree(s.Right, t.Right) {
		return false
	}

	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	return false
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func maxDepth(root *TreeNode) int {
	//if root == nil {
	//	return 0
	//}
	//
	//return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	res := 1

	queue = append(queue, root)

	for {
		if len(queue) == 0 {
			break
		}

		tmp := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}

			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}

		}

		res = res + 1
		queue = tmp
	}
	return res
}
