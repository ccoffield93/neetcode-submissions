/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    a := list1
	b := list2
	// edge case handling
	if a == nil {
		return b
	}
	if b == nil {
		return a 
	}
	var merged *ListNode 
	if a.Val <= b.Val {
		merged = a
		a = a.Next 
	} else {
		merged = b 
		b = b.Next 
	}
	head := merged 

	for (a != nil && b != nil) {
		if a.Val <= b.Val {
			merged.Next = a 
			merged = a 
			a = a.Next 
		} else {
			merged.Next = b 
			merged = b
			b = b.Next 
		}
	}

	// now whatever's left in the case of uneven lists 
	for a != nil {
		merged.Next = a 
		merged = a 
		a = a.Next 
	}

	for b != nil {
		merged.Next = b 
		merged = b
		b = b.Next
	}

	return head
}
