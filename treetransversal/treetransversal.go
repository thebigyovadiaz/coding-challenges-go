package treetransversal

import "fmt"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal performs an inorder traversal of a binary tree.
func InorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)
	fmt.Printf("%d ", root.Value)
	InorderTraversal(root.Right)
}

// PreorderTraversal performs a preorder traversal of a binary tree.
func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Value)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

// PostorderTraversal performs a postorder traversal of a binary tree.
func PostorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left)
	PostorderTraversal(root.Right)
	fmt.Printf("%d ", root.Value)
}

func TreeTransversal() {
	// Example usage:
	// Constructing a binary tree:       1
	//                                   / \
	//                                  2   3
	//                                 / \
	//                                4   5

	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 4,
			},
			Right: &TreeNode{
				Value: 5,
			},
		},
		Right: &TreeNode{
			Value: 3,
		},
	}

	fmt.Println("Inorder Traversal:")
	InorderTraversal(root)
	fmt.Println()

	fmt.Println("Preorder Traversal:")
	PreorderTraversal(root)
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	PostorderTraversal(root)
	fmt.Println()
}
