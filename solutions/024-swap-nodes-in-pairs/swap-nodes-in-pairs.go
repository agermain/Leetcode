// Given a linked list, swap every two adjacent nodes and return its head.
//
// You may not modify the values in the list's nodes. Only nodes itself may be changed.
//
//  
// Example 1:
//
//
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]
//
//
// Example 2:
//
//
// Input: head = []
// Output: []
//
//
// Example 3:
//
//
// Input: head = [1]
// Output: [1]
//
//
//  
// Constraints:
//
//
// 	The number of nodes in the list is in the range [0, 100].
// 	0 <= Node.val <= 100
//
//


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
    // We swap the first two, and will use this head 
    // to append the rest to it and return it
	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
    // Append the recursive'd list to the head
	newHead.Next = head
	return newHead
}
