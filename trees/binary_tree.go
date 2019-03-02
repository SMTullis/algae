package algae

import (
    "errors"
    "fmt"
)

type Node struct {
    content Container
    parent *Node
    left *Node
    right *Node
}

func (n *Node) setLeft(left *Node) {
    n.left = left
    n.left.setParent(n)
}

func (n *Node) setRight(right *Node) {
    n.right = right
    n.right.setParent(n)
}

func (n *Node) setParent(parent *Node) {
    n.parent = parent
}

func (n Node) GetContent() *Container {
    return &n.content
}

func (n Node) setContent(content Container) {
    n.content = content
}

func (n Node) Insert(other *Node) {
    if n.content.IsGreater(*other.GetContent()) {
        if n.checkLeft() {
            n.left.Insert(other)
        } else {
            n.setLeft(other)
        }
    } else if n.content.IsLesser(*other.GetContent()) {
        if n.checkRight() {
            n.right.Insert(other)
        } else {
            n.setRight(other)
        }
    }
}

func (n Node) checkLeft() bool {
    return n.left != nil
}

func (n Node) checkRight() bool {
    return n.right != nil
}

func (n Node) checkParent() bool {
    return n.parent != nil
}

func (n Node) clearLeft() {
    n.left = nil
}

func (n Node) clearRight() {
    n.right = nil
}

func (n Node) clearParent() {
    n.parent = nil
}

func (n Node) TraverseInorder() {
    if n.checkLeft() {n.left.TraverseInorder()}
    fmt.Println(n.content.GetValue())
    if n.checkRight() {n.right.TraverseInorder()}
}

func (n Node) TraversePostorder() {
    if n.checkLeft() {n.left.TraversePostorder()}
    if n.checkRight() {n.right.TraversePostorder()}
    fmt.Println(n.content.GetValue())
}

func (n Node) TraversePreorder() {
    fmt.Println(n.content.GetValue())
    if n.checkLeft() {n.left.TraversePreorder()}
    if n.checkRight() {n.right.TraversePreorder()}
}

func (n Node) shiftLeft() {
    if n.checkRight() {
        n.right.shiftRight()
    }
    if n.checkParent() {
        n.right.setParent(n.parent)
    } else {
        n.right.clearParent()
    }
    n.right.setLeft(&n)
    n.clearRight()
}

func (n Node) shiftRight() {
    if n.checkLeft() {
        n.left.shiftLeft()
    }
    if n.checkParent() {
        n.left.setParent(n.parent)
    } else {
        n.left.clearParent()
    }
    n.left.setRight(&n)
    n.left.clearLeft()
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
        match = n.left.search(value)
    } else if n.checkRight() {
        match = n.right.search(value)
    }
    return match
}

func (n Node) Delete(value string) error {
    match, err := n.Search(value)
    if err != nil {
        match.delete()
    }
    return err
}

func (n Node) delete() {
    if n.checkParent() {
        if n.parent.GetContent().IsGreater(*n.GetContent()) {
            if n.checkRight() {
                n.right.shiftRight()
                n.right.setLeft(n.left)
                n.right.setParent(n.parent)
            }
            n.parent.setLeft(n.right)
        } else {
            if n.checkLeft() {
                n.left.shiftLeft()
                n.left.setRight(n.right)
                n.left.setParent(n.parent)
            }
            n.parent.setRight(n.left)
        }
    } else {
        if n.checkRight() {
            n.right.shiftRight()
            n.right.setLeft(n.left)
            n.right.setParent(n.parent)
        } else if n.checkLeft() {
            n.left.shiftLeft()
            n.left.setRight(n.right)
            n.left.setParent(n.parent)
        }
    }
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
