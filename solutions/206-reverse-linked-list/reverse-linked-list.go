// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
//  
// Example 1:
//
//
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
//
// Example 2:
//
//
// Input: head = [1,2]
// Output: [2,1]
//
//
// Example 3:
//
//
// Input: head = []
// Output: []
//
//
//  
// Constraints:
//
//
// 	The number of nodes in the list is the range [0, 5000].
// 	-5000 <= Node.val <= 5000
//
//
//  
// Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
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
