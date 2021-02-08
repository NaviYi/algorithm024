/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var result []int
func preorderTraversal(root *TreeNode) []int {
    result = make([]int,0)
    if root == nil{
        return result
    }
    traversal(root)
    return result
}
func traversal(root *TreeNode){
    if root == nil{
        return
    }
    result = append(result,root.Val)
    traversal(root.Left)
    traversal(root.Right)
}