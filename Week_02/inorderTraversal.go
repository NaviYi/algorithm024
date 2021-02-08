/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var result []int
func inorderTraversal(root *TreeNode) []int {
    result = make([]int,0)
    if root == nil{
        return []int{}
    }
    traversal(root)
    return result
}
func traversal(root *TreeNode){
    if root == nil{
        return
    }
    traversal(root.Left)
    result = append(result,root.Val)
    traversal(root.Right)
}