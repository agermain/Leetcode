// Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.
//
// Follow up:
// What if the linked list is extremely large and its length is unknown to you? Could you solve this efficiently without using extra space?
//
//
// Example:
//
// // Init a singly linked list [1,2,3].
// ListNode head = new ListNode(1);
// head.next = new ListNode(2);
// head.next.next = new ListNode(3);
// Solution solution = new Solution(head);
//
// // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
// solution.getRandom();
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
