/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	visitedMap := make(map[*ListNode]bool)

	curr := head
	for curr != nil {
		_, hit := visitedMap[curr]

		if hit {
			return true
		}
		visitedMap[curr] = true
		curr = curr.Next 
	}
	return false
}
