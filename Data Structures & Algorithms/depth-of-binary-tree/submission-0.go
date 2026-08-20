/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /* 

 Algorithm BFS(G, v)
    Q ← new empty FIFO queue
    Mark v as visited.
    Q.enqueue(v)
    while Q is not empty
        a ← Q.dequeue()
        // Perform some operation on a.
        for all unvisited neighbors x of a
            Mark x as visited.
            Q.enqueue(x)
*/ 

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root} // use a slice and just move it 
	depth := 0

	for len(queue) > 0 {
		depth++
		
		// Process all nodes currently in the queue (one full level)
		levelSize := len(queue) 
		for i := 0; i < levelSize; i++ {
			// "dequeue" (take item from beginning of slice and remove it)
			check := queue[0]
			queue = queue[1:]

			if check.Left != nil {
				queue = append(queue, check.Left)
			}
			if check.Right != nil {
				queue = append(queue, check.Right)
			}
		}

		// after this, the queue will have the 'next level' of breadth
		// which will increase depth and go again
	}

	return depth
}
	