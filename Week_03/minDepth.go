/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    //空节点
    if root == nil{
        return 0
    }
    //只有一个节点
    if root.Left == nil && root.Right == nil{
        return 1
    }
    left := minDepth(root.Left)
    right := minDepth(root.Right)
    //当节点的左子树和右子树有一个为空，说明left和right必定有一个为0
    if root.Left == nil || root.Right == nil{
        return left+right+1
    }
    return min(left,right) + 1
}
func min(x,y int) int {
    if x < y {
        return x
    }else{
        return y
    }
}