# Given the root of a binary tree, return the sum of values of its deepest leaves.
#  
# Example 1:
#
#
# Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
# Output: 15
#
#
# Example 2:
#
#
# Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
# Output: 19
#
#
#  
# Constraints:
#
#
# 	The number of nodes in the tree is in the range [1, 104].
# 	1 <= Node.val <= 100
#
#


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def deepestLeavesSum(self, root: TreeNode) -> int:
        queue = [root]
        while queue:
            new = []
            for node in queue:
                if node.left:
                    new.append(node.left)
                if node.right:
                    new.append(node.right)
            # Actually store last level
            # by detecting that there's no nodes after this level
            if not new:
                break
            queue = new
        return sum([node.val for node in queue])
