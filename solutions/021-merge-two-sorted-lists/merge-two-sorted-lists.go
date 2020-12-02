// Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
//
//  
// Example 1:
//
//
// Input: l1 = [1,2,4], l2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
//
// Example 2:
//
//
// Input: l1 = [], l2 = []
// Output: []
//
//
// Example 3:
//
//
// Input: l1 = [], l2 = [0]
// Output: [0]
//
//
//  
// Constraints:
//
//
// 	The number of nodes in both lists is in the range [0, 50].
// 	-100 <= Node.val <= 100
// 	Both l1 and l2 are sorted in non-decreasing order.
//
//


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // Use tmpNode but return head.Next
    head := &ListNode{}
    tmpNode := head

    for {
        // If one of the two lists is empty, just append the other one
        if l1 == nil {
            tmpNode.Next = l2
            break
        }
        
         if l2 == nil {
            tmpNode.Next = l1
             break
        }
        
        if l1.Val < l2.Val {
            tmpNode.Next = l1
            l1 = l1.Next
        } else {
            tmpNode.Next = l2
            l2 = l2.Next
        }
        // Update the tail of the list
        tmpNode = tmpNode.Next
    }
    
    return head.Next
}
