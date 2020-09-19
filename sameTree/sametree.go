package sameTree

func IsSameTreeMethod1(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTreeMethod1(p.Left, q.Left) && IsSameTreeMethod1(p.Right, q.Right)

}
