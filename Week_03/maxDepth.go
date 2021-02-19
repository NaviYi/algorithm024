/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    //nothing to do

    //drill down
    l := maxDepth(root.Left)
    r := maxDepth(root.Right)
    return Max(l,r) + 1
}
func Max(x,y int) int {
    if x > y{
        return x
    }else{
        return y
    }
}