package utils

import "sync"

type _paths struct {
	paths []string
	focus int
}

type NavStack struct {
	paths _paths
	index int
}

func NewNavStack() NavStack {
	return NavStack{
		paths: _paths{
			paths: make([]string, 50, 50),
			focus: -1,
		},
		index: -1,
	}
}

func (n *NavStack) Push(s string, focus int) {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()

	n.paths.focus = focus
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
