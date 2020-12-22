# Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
#
#
# For example:
# Given binary tree [3,9,20,null,null,15,7],
#
#     3
#    / \
#   9  20
#     /  \
#    15   7
#
#
#
# return its zigzag level order traversal as:
#
# [
#   [3],
#   [20,9],
#   [15,7]
# ]
#
#


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        if root == None:
            return []
        self.zigzag = True
        queue = [root]
        res = []
        while queue:
            nextLevel = []
            level = [node.val for node in queue]
            if self.zigzag:
                res.append(level)
            else:
                res.append(level[::-1])
            for node in queue:
                if node.left: 
                    nextLevel.append(node.left)
                if node.right: 
                    nextLevel.append(node.right)
            if self.zigzag == True:
                self.zigzag = False
            else:
                self.zigzag = True
            queue = nextLevel
        return res
