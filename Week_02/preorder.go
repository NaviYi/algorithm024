/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var reslut []int
func preorder(root *Node) []int {
    reslut = make([]int,0)
    if root == nil{
        return reslut
    }
    order(root)
    return reslut
}

func order(root *Node){
    if root == nil{
        return
    }
    reslut = append(reslut,root.Val)
    for _,node := range root.Children{
        order(node)
    }
}