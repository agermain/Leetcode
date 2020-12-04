// Reverse a singly linked list.
//
// Example:
//
//
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
//
//
// Follow up:
//
// A linked list can be reversed either iteratively or recursively. Could you implement both?
//


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    s := Solution{}
    s.helper(head)
    return s.head
}

type Solution struct {
    head *ListNode
}

func (s *Solution) helper(lastHead *ListNode) {
    if lastHead == nil {
        return
    }
    // We have reached the end of the Linked List, so it should be the new head.
    if lastHead.Next == nil {
        s.head = lastHead
        return
    }
    // Get to the bottom of the call stack
    s.helper(lastHead.Next)
    
    // Set the next node's "Next" to current node.
    lastHead.Next.Next = lastHead
    
    // Remove the link between current and next node
    // as it will be set later in the call stack
    lastHead.Next = nil
}
