/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
    var toReturn *ListNode 

	list := []*ListNode{}
    curr := head
	for curr != nil {
		list = append(list, curr)
		curr = curr.Next
	}

	toReturn = list[len(list) - 1]
	curr = toReturn 
	for i := len(list) - 2; i >= 0; i-- {
		curr.Next = list[i]
		curr = curr.Next
	}

	// end of the list 
	curr.Next = nil 

	return toReturn
}
