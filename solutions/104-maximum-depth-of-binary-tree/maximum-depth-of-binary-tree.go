// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//  
// Example 1:
//
//
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
//
//
// Example 2:
//
//
// Input: root = [1,null,2]
// Output: 2
//
//
// Example 3:
//
//
// Input: root = []
// Output: 0
//
//
// Example 4:
//
//
// Input: root = [0]
// Output: 1
//
//
//  
// Constraints:
//
//
// 	The number of nodes in the tree is in the range [0, 104].
// 	-100 <= Node.val <= 100
//


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    return max(left, right)+1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
