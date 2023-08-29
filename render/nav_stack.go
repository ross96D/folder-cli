package render

import (
	"os"

	"github.com/nsf/termbox-go"
)

type NavTree struct {
	tree  *Tree
	node  *Node
	index int
	// currpath string
}

func NewNavTree() NavTree {
	_, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	tree, index := NewTree()
	node := tree.GetNode(index)

	nav := NavTree{
		tree:  tree,
		index: index,
		node:  node,
	}
	// lis := NewList()
	// lis.Repopulate(home, 0)
	// nav.paths[0] = lis
	return nav
}

// If dir is positive will change to the rigth.
// If dir is negative will change to the left.
//
// If dir is 0 will panic
func (n *NavTree) ChangeIndex(dir int) {
	if dir == 0 {
		panic("Dir cannot be 0")
	}
	if dir > 0 {
		newNode := n.node.FousedNode()
		if newNode != nil {
			n.node = newNode
		}
	} else {
		newNode := n.node.parent
		if newNode != nil {
			n.node = newNode
		}
	}
}

func (n *NavTree) HandleEvent(e termbox.Event) bool {
	handled := true
	if e.Key == termbox.KeyArrowUp {
		n.node.Focus -= 1
	} else if e.Key == termbox.KeyArrowDown {
		n.node.Focus += 1
	} else if e.Key == termbox.KeyArrowRight {
		n.ChangeIndex(1)
	} else if e.Key == termbox.KeyArrowLeft {
		n.ChangeIndex(-1)
	} else {
		handled = false
	}
	return handled
}
