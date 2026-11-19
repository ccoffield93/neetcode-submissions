/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reorder such that [0, 1, ... n] -> [0, n-1, 1, n-2...]
func reorderList(head *ListNode) {
    index := make(map[int]*ListNode)

	track := head
	i := 0
	for track != nil {
		index[i] = track
		track = track.Next
		i++
	}
	fmt.Println(index)
	n := i
	left, right := 0, n-1
	for left < right {
		index[left].Next = index[right]
		left++

		// Stop if pointers met after moving left forward
		if left == right {
			break
		}

		index[right].Next = index[left]
		right--
	}

	// Terminate the reordered list to prevent a cycle
	index[left].Next = nil
}
