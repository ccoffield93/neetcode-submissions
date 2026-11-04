/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    currMaxDiameter := 0

	// anonymous function to get the depth (height) of a node 
	// we have to declare it outside the func itself so that 
	// it's in scope for recursion
	var dfs func(root *TreeNode) int 
	dfs = func(root *TreeNode) int{
		// base case
		if root == nil {
			return 0 
		}

		// recursively find the depth of left and right 
		left := dfs(root.Left) 
		right := dfs(root.Right)

		// check the max diameter-- diameter is the sum of 
		// left and right children's depths
		currMaxDiameter = max(currMaxDiameter, left + right)
		
		// depth is one (current node) plus the depth of its 
		// longest chain of children
		return 1+max(left, right)

	}

	dfs(root)
	
	return currMaxDiameter
}

