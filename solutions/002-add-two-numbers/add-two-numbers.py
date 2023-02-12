# You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
#
# You may assume the two numbers do not contain any leading zero, except the number 0 itself.
#
#  
# Example 1:
#
#
# Input: l1 = [2,4,3], l2 = [5,6,4]
# Output: [7,0,8]
# Explanation: 342 + 465 = 807.
#
#
# Example 2:
#
#
# Input: l1 = [0], l2 = [0]
# Output: [0]
#
#
# Example 3:
#
#
# Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
# Output: [8,9,9,9,0,0,0,1]
#
#
#  
# Constraints:
#
#
# 	The number of nodes in each linked list is in the range [1, 100].
# 	0 <= Node.val <= 9
# 	It is guaranteed that the list represents a number that does not have leading zeros.
#
#


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        n1 = str(l1.val)
        n2 = str(l2.val)
        while l1.next:
            n1 += str(l1.next.val)
            l1 = l1.next
        while l2.next:
            n2 += str(l2.next.val)
            l2 = l2.next
        answer = int(n1[::-1]) + int(n2[::-1])
        ans = [int(n) for n in list(str(answer))]
        ans.reverse()
        head = ListNode(ans[0])
        # Check boundaries of answer
        if len(ans) > 1:
            cur = head
            for i in range(1, len(ans)):
                cur.next = ListNode(ans[i])
                cur = cur.next
        return head

