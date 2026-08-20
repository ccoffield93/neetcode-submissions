/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
  * Algorithm for DFS: 
  
  Algorithm DFS(G, v)
    if v is already visited
        return        
    Mark v as visited.
    // Perform some operation on v.
    for all neighbors x of v
        DFS(G, x)
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
    visited := make(map[*TreeNode]bool)
	dfsHelper(root, visited)
	return root
}

func dfsHelper(node *TreeNode, visited map[*TreeNode]bool) {
	// If we've already seen this node, stop, otherwise mark it as seen
	if visited[node] {
		return
	}
	visited[node] = true

	// get left and right children, swap them
	l := node.Left
	r := node.Right
	node.Left = r
	node.Right = l

	// perform operation recursively on children
	if l != nil && !visited[l]{
		dfsHelper(l, visited)
	}

	if r!= nil && !visited[r] {
		dfsHelper(r, visited)
	}
}