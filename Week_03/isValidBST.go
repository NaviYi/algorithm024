/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import(
     "fmt"
 )
var mid []int
func isValidBST(root *TreeNode) bool {
    mid = make([]int,0)
    if root == nil{
        return true
    }
    inorder(root)
    fmt.Println(mid)
    max := mid[0]
    //验证是否单调递增
    result := true
    for i := 1;i<len(mid);i++{
        if mid[i] > max{
            max = mid[i]
        }else{
            result = false
        }
    }
    return result
}
func inorder(root *TreeNode){
    if root == nil{
        return
    }
    inorder(root.Left)
    mid = append(mid,root.Val)
    inorder(root.Right)
}