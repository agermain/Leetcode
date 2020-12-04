// Given the root of a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only one right child.
//
//  
// Example 1:
//
//
// Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//
//
// Example 2:
//
//
// Input: root = [5,1,7]
// Output: [1,null,5,null,7]
//
//
//  
// Constraints:
//
//
// 	The number of nodes in the given tree will be in the range [1, 100].
// 	0 <= Node.val <= 1000
//


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	arr := []*TreeNode{}
	inOrder(root, &arr)
	// We now have the nodes in the right order
	// However, there are still "Left" links.
	// Delete the "Left" links and link "Right" to next in the array
	for i := 0; i < len(arr)-1; i++ {
		arr[i].Left = nil
		arr[i].Right = arr[i+1]
        // Deal with the "Left" of the last node in the tree
		if i == len(arr)-2 {
			arr[i+1].Left = nil
		}
	}
	return arr[0]
}

func inOrder(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root)
	inOrder(root.Right, arr)
}

