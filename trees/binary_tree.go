package algae

import "fmt"

type Node struct {
    content Container
    Parent *Node
    Left *Node
    Right *Node
}

func (n *Node) SetLeft(left *Node) {
    n.Left = left
}

func (n *Node) SetRight(right *Node) {
    n.Right = right
}

func (n *Node) SetParent(parent *Node) {
    n.Parent = parent
}

func (n Node) GetContent() Container {
    return n.content
}

func (n Node) InsertNode(other *Node) {
    if n.content.IsGreater(other.GetContent()) {
        n.checkLeft(other)
    } else if n.content.IsLesser(other.GetContent()) {
        n.checkRight(other)
    }
}

func (n Node) checkLeft(other *Node) {
    if n.Left == nil {
        n.SetLeft(other)
    } else {
        n.Left.InsertNode(other)
    }
}

func (n Node) checkRight(other *Node) {
    if n.Right == nil {
        n.SetRight(other)
    } else {
        n.Right.InsertNode(other)
    }
}

func (n Node) TraverseInorder() {
    if n.Left != nil {n.Left.TraverseInorder()}
    fmt.Println(n.content.GetValue())
    if n.Right != nil {n.Right.TraverseInorder()}
}

func (n Node) TraversePostorder() {
    if n.Left != nil {n.Left.TraversePostorder()}
    if n.Right != nil {n.Right.TraversePostorder()}
    fmt.Println(n.content.GetValue())
}

func (n Node) TraversePreorder() {
    fmt.Println(n.content.GetValue())
    if n.Left != nil {n.Left.TraversePreorder()}
    if n.Right != nil {n.Right.TraversePreorder()}
}

type Container struct {
    value string
}

func (con Container) GetValue() string {
    return con.value
}

func (con Container) IsGreater(other Container) bool {
    return con.GetValue() > other.GetValue()
}

func (con Container) IsGreaterOrEqual(other Container) bool {
    return con.GetValue() >= other.GetValue()
}

func (con Container) IsEqual(other Container) bool {
    return con.GetValue() == other.GetValue()
}

func (con Container) IsNotEqual(other Container) bool {
    return con.GetValue() != other.GetValue()
}

func (con Container) IsLesserOrEqual(other Container) bool {
    return con.GetValue() <= other.GetValue()
}

func (con Container) IsLesser(other Container) bool {
    return con.GetValue() < other.GetValue()
}