// Given the root of a Binary Search Tree (BST), return the minimum absolute difference between the values of any two different nodes in the tree.
//
//  
// Example 1:
//
//
// Input: root = [4,2,6,1,3]
// Output: 1
//
//
// Example 2:
//
//
// Input: root = [1,0,48,null,null,12,49]
// Output: 1
//
//
//  
// Constraints:
//
//
// 	The number of nodes in the tree is in the range [2, 104].
// 	0 <= Node.val <= 105
//
//
//  
// Note: This question is the same as 783: https://leetcode.com/problems/minimum-distance-between-bst-nodes/
//


/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    private int min = Integer.MAX_VALUE;
    Integer prev = null;
    public int getMinimumDifference(TreeNode root) {
        if (root == null) return min;
        getMinimumDifference(root.left);
        if (prev != null) {
            min = Math.min(min, root.val - prev);
        }
        prev = root.val;
        getMinimumDifference(root.right);
        return min;
    }
}
