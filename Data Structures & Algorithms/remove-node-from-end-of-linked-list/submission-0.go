/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    last := head 
	nthFromLast := head
	nminusoneFromLast := head
	for i := 0; i < n; i++ {
		last = last.Next 
		if last == nil && i != n-1 {
			return nil
		}
	}

	for last != nil{
		last = last.Next
		nthFromLast = nthFromLast.Next
		if last != nil {
			nminusoneFromLast = nminusoneFromLast.Next
		}
	}

    if nminusoneFromLast == nthFromLast {
		// we have to remove first element
		return head.Next
	}

	nminusoneFromLast.Next = nthFromLast.Next

	return head
}
