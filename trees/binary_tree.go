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

func (n *Node) SetLeft(left *Node) {
    n.Left = left
    n.Left.SetParent(n)
}

func (n *Node) SetRight(right *Node) {
    n.Right = right
    n.Right.SetParent(n)
}

func (n *Node) SetParent(parent *Node) {
    n.Parent = parent
}

func (n Node) GetContent() *Container {
    return &n.content
}

func (n Node) SetContent(content Container) {
    n.content = content
}

func (n Node) InsertNode(other *Node) {
    if n.content.IsGreater(*other.GetContent()) {
        if n.checkLeft() {
            n.Left.InsertNode(other)
        } else {
            n.SetLeft(other)
        }
    } else if n.content.IsLesser(*other.GetContent()) {
        if n.checkRight() {
            n.Right.InsertNode(other)
        } else {
            n.SetRight(other)
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

func (n Node) ClearLeft() {
    n.Left = nil
}

func (n Node) ClearRight() {
    n.Right = nil
}

func (n Node) ClearParent() {
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

func (n Node) ShiftLeft() {
    if n.checkRight() {
        n.Right.ShiftRight()
    }
    if n.checkParent() {
        n.Right.SetParent(n.Parent)
    } else {
        n.Right.ClearParent()
    }
    n.Right.SetLeft(&n)
    n.ClearRight()
}

func (n Node) ShiftRight() {
    if n.checkLeft() {
        n.Left.ShiftLeft()
    }
    if n.checkParent() {
        n.Left.SetParent(n.Parent)
    } else {
        n.Left.ClearParent()
    }
    n.Left.SetRight(&n)
    n.Left.ClearLeft()
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
