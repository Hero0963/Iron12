package invertBinaryTree

func InvertTreeMethod1(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTreeMethod1(root.Left)
	InvertTreeMethod1(root.Right)

	return root
}
