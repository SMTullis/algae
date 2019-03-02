package algae

import (
    "errors"
    "fmt"
)

type Node struct {
    content Container
    Parent *Node
    Left *Node
    Right *Node
}

func (n *Node) setLeft(left *Node) {
    n.Left = left
    n.Left.setParent(n)
}

func (n *Node) setRight(right *Node) {
    n.Right = right
    n.Right.setParent(n)
}

func (n *Node) setParent(parent *Node) {
    n.Parent = parent
}

func (n Node) GetContent() *Container {
    return &n.content
}

func (n Node) setContent(content Container) {
    n.content = content
}

func (n Node) InsertNode(other *Node) {
    if n.content.IsGreater(*other.GetContent()) {
        if n.checkLeft() {
            n.Left.InsertNode(other)
        } else {
            n.setLeft(other)
        }
    } else if n.content.IsLesser(*other.GetContent()) {
        if n.checkRight() {
            n.Right.InsertNode(other)
        } else {
            n.setRight(other)
        }
    }
}

func (n Node) checkLeft() bool {
    return n.Left != nil
}

func (n Node) checkRight() bool {
    return n.Right != nil
}

func (n Node) checkParent() bool {
    return n.Parent != nil
}

func (n Node) clearLeft() {
    n.Left = nil
}

func (n Node) clearRight() {
    n.Right = nil
}

func (n Node) clearParent() {
    n.Parent = nil
}

func (n Node) TraverseInorder() {
    if n.checkLeft() {n.Left.TraverseInorder()}
    fmt.Println(n.content.GetValue())
    if n.checkRight() {n.Right.TraverseInorder()}
}

func (n Node) TraversePostorder() {
    if n.checkLeft() {n.Left.TraversePostorder()}
    if n.checkRight() {n.Right.TraversePostorder()}
    fmt.Println(n.content.GetValue())
}

func (n Node) TraversePreorder() {
    fmt.Println(n.content.GetValue())
    if n.checkLeft() {n.Left.TraversePreorder()}
    if n.checkRight() {n.Right.TraversePreorder()}
}

func (n Node) shiftLeft() {
    if n.checkRight() {
        n.Right.shiftRight()
    }
    if n.checkParent() {
        n.Right.setParent(n.Parent)
    } else {
        n.Right.clearParent()
    }
    n.Right.setLeft(&n)
    n.clearRight()
}

func (n Node) shiftRight() {
    if n.checkLeft() {
        n.Left.shiftLeft()
    }
    if n.checkParent() {
        n.Left.setParent(n.Parent)
    } else {
        n.Left.clearParent()
    }
    n.Left.setRight(&n)
    n.Left.clearLeft()
}

func (n Node) Find(value string) (*Container, error) {
    if match, err := n.Search(value); err != nil {
        return nil, err
    } else {
        return (*match).GetContent(), nil
    }
}

func (n Node) Search(value string) (*Node, error) {
    if match := n.search(value); match == nil {
        return match, errors.New("Value not found.")
    } else {
        return match, nil
    }
}

func (n Node) search(value string) *Node {
    var match *Node
    self := n.GetContent()
    if x := self.GetValue(); value == x {
        match = &n
    } else if n.checkLeft() && value < x {
        match = n.Left.search(value)
    } else if n.checkRight() {
        match = n.Right.search(value)
    }
    return match
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
