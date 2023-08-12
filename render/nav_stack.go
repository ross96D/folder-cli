package render

import (
	"os"
	"sync"
)

type NavStack struct {
	pathList []ListItem
	paths    []string
	index    int
}

func NewNavStack() NavStack {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	nav := NavStack{
		pathList: make([]ListItem, 50, 50),
		paths:    make([]string, 50, 50),
		index:    -1,
	}
	lis := NewList()
	lis.Repopulate(home, false, 0)
	nav.pathList[0] = lis
	return nav
}

func (n *NavStack) Push(s string, focus int) {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()

	n.index++
	n.paths.paths[n.index] = s
}

// this function panics when stack is empty
func (n *NavStack) Pop() string {
	m := sync.Mutex{}
	m.Lock()

	if n.index == 0 {
		return n.paths.paths[n.index]
	}
	result := n.paths.paths[n.index]
	n.index -= 1
	return result
}

func (n *NavStack) Get() string {
	return n.paths.paths[n.index]
}

func (n *NavStack) GetFocus() int {
	return n.paths.focus
}
