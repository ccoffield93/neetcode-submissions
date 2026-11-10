/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Traverse root, looking for anywhere that the val of a node 
	// matches the val of the subroot

	// if the val of a subroot is found, do a check on that branch
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// process item at zero 
		curr := queue[0]
		if curr.Val == subRoot.Val {
			check := sameTree(curr, subRoot)
			if check == true {
				return true
			}
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}

		// remove processed item 
		queue = queue[1:]
	}

	return false 
}

func sameTree(first *TreeNode, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	}
	if first.Val != second.Val {
		return false
	}
	if first.Left != second.Left && (first.Left == nil || second.Left == nil) {
		// one is nil, the other isn't
		return false 
	}
	if first.Right != second.Right && (first.Right == nil || second.Right == nil) {
		// one is nil, the other isn't
		return false 
	}

	return (sameTree(first.Left, second.Left) && sameTree(first.Right, second.Right))
}