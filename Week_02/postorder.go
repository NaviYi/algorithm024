/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
var result []int
func postorder(root *Node) []int {
    result = make([]int,0)
    if root == nil{
        return result
    }

    order(root)
    return result
}

func order(root *Node){
    if root == nil{
        return
    }
    for _,node := range root.Children{
        order(node)
    }
    result = append(result,root.Val)
}
返回该题