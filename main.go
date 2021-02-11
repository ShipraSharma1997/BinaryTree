package main

import "fmt"

type Node struct {
    data        string
    left, right *Node
}

func main() {
    nodeG := Node{data: "10", left: nil, right: nil}
    nodeF := Node{data: "20", left: &nodeG, right: nil}
    nodeE := Node{data: "30", left: nil, right: nil}
    nodeD := Node{data: "40", left: &nodeE, right: nil}
    nodeC := Node{data: "50", left: nil, right: nil}
    nodeB := Node{data: "60", left: &nodeD, right: &nodeF}
    nodeA := Node{data: "70", left: &nodeB, right: &nodeC}
    fmt.Println("Preorder:")
    nodeA.PrintPre()
    fmt.Println("Inorder:")
	nodeA.PrintIn()
    fmt.Println("Postorder:")
    nodeA.PrintPost()
}

// Preorder (Root, Left, Right)
func (root *Node) PrintPre() {
    fmt.Println(root.data)
    if root.left != nil {
        root.left.PrintPre()
    }
    if root.right != nil {
        root.right.PrintPre()
    }
}

// Inorder (Left, Root, Right)
func (root *Node) PrintIn() {
    if root.left != nil {
        root.left.PrintIn()
    }
    fmt.Println(root.data)
    if root.right != nil {
        root.right.PrintIn()
    }
}

// Postorder (Left, Right, Root)
func (root *Node) PrintPost() {
    if root.left != nil {
        root.left.PrintPost()
    }
    if root.right != nil {
		root.right.PrintPost()
    }
    fmt.Println(root.data)
}