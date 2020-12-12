// Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.
//
//  
// Example 1:
//
//
// Input
// ["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
// [[[1, 2, 3]], [], [], [], [], []]
// Output
// [null, 1, 3, 2, 2, 3]
//
// Explanation
// Solution solution = new Solution([1, 2, 3]);
// solution.getRandom(); // return 1
// solution.getRandom(); // return 3
// solution.getRandom(); // return 2
// solution.getRandom(); // return 2
// solution.getRandom(); // return 3
// // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
//
//
//  
// Constraints:
//
//
// 	The number of nodes in the linked list will be in the range [1, 104]
// 	-104 <= Node.val <= 104
// 	At most 104 calls will be made to getRandom.
//
//
//  
// Follow up:
//
//
// 	What if the linked list is extremely large and its length is unknown to you?
// 	Could you solve this efficiently without using extra space?
//
//


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    root *ListNode
    arr []ListNode
    length int
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    var arr []ListNode
    var length int
    for head != nil {
        arr = append(arr, *head)
        head = head.Next
        length++
    }
    return Solution{head, arr, length}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    return this.arr[rand.Intn(this.length)].Val
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
