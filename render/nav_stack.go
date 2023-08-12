package render

import (
	"os"
	"sync"

	"github.com/nsf/termbox-go"
)

type NavStack struct {
	paths    []ListItem
	index    int
	currpath string
}

func NewNavStack() NavStack {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	nav := NavStack{
		paths:    make([]ListItem, 50, 50),
		index:    0,
		currpath: home,
	}
	lis := NewList()
	lis.Repopulate(home, 0)
	nav.paths[0] = lis
	return nav
}

func (n *NavStack) Push(s string) {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()

	n.currpath += "/" + s
	n.index++
	if n.paths[n.index].FolderName != s {
		n.paths[n.index] = NewList()
		n.paths[n.index].Repopulate(n.currpath, 0)
	}
}

// this function panics when stack is empty
func (n *NavStack) Pop() *ListItem {
	m := sync.Mutex{}
	m.Lock()

	if n.index == 0 {
		return &n.paths[n.index]
	}
	n.paths[n.index].clear()
	n.index -= 1
	result := n.paths[n.index]
	result.Draw()
	return &result
}

func (n *NavStack) Get() *ListItem {
	return &n.paths[n.index]
}

func (n *NavStack) GetFocus() int {
	return n.paths[n.index].iFocus
}

func (n *NavStack) HandleEvent(e termbox.Event) bool {
	handled := true
	if e.Key == termbox.KeyArrowUp {
		n.Get().Focus(-1)
	} else if e.Key == termbox.KeyArrowDown {
		n.Get().Focus(1)
	} else if e.Key == termbox.KeyArrowRight {
		if n.Get().FocusItem().IsDir {
			newpath := n.Get().FolderName + "/" + n.Get().FocusItem().Name
			n.Push(newpath)
		}
	} else if e.Key == termbox.KeyArrowLeft {
		n.Pop()
	} else {
		handled = false
	}
	return handled
}
