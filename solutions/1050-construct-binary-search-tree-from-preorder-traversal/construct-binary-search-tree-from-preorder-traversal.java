// Return the root node of a binary search tree that matches the given preorder traversal.
//
// (Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)
//
// It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.
//
// Example 1:
//
//
// Input: [8,5,1,7,10,12]
// Output: [8,5,10,1,7,null,12]
//
//
//
//  
// Constraints:
//
//
// 	1 <= preorder.length <= 100
// 	1 <= preorder[i] <= 10^8
// 	The values of preorder are distinct.
//
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
    private int i;
    public TreeNode bstFromPreorder(int[] preorder) {
        return helper(preorder, Integer.MAX_VALUE);
    }
    public TreeNode helper(int[] preorder, int n) {
        if (i == preorder.length || preorder[i] > n) return null;
        TreeNode node = new TreeNode(preorder[i++]);
        node.left = helper(preorder, node.val);
        node.right = helper(preorder, n);
        return node;
    }
}
