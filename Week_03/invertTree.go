/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    invert(root)
    return root
}

func invert(root *TreeNode){
    if root == nil{
        return
    }
    if root.Left == nil && root.Right == nil{
        return
    }

    //what to do
    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp
    //down
    invert(root.Left)
    invert(root.Right)
}