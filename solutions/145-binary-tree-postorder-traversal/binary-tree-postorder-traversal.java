// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
// Follow up: Recursive solution is trivial, could you do it iteratively?
//
//  
// Example 1:
//
//
// Input: root = [1,null,2,3]
// Output: [3,2,1]
//
//
// Example 2:
//
//
// Input: root = []
// Output: []
//
//
// Example 3:
//
//
// Input: root = [1]
// Output: [1]
//
//
// Example 4:
//
//
// Input: root = [1,2]
// Output: [2,1]
//
//
// Example 5:
//
//
// Input: root = [1,null,2]
// Output: [2,1]
//
//
//  
// Constraints:
//
//
// 	The number of the nodes in the tree is in the range [0, 100].
// 	-100 <= Node.val <= 100
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
    public List<Integer> postorderTraversal(TreeNode root) {
        List<Integer> out = new ArrayList<>();
        helper(root, out);
        return out;
    }
    private void helper(TreeNode root, List<Integer> result) {
        if (root != null) {
            helper(root.left, result);
            helper(root.right, result);
            result.add(root.val);
        }
    }
}
