/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	originalToCopy := make(map[*Node]*Node)

	// Pass 1: create all copied nodes
	for curr := head; curr != nil; curr = curr.Next {
		originalToCopy[curr] = &Node{Val: curr.Val}
	}

	// Pass 2: assign Next and Random pointers on the copied nodes
	for curr := head; curr != nil; curr = curr.Next {
		copyNode := originalToCopy[curr]
		copyNode.Next = originalToCopy[curr.Next]
		copyNode.Random = originalToCopy[curr.Random]
	}

	return originalToCopy[head]
}
