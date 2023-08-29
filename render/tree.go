package render

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/exp/slices"
)

func getEntries(path string, n *Node) []fs.DirEntry {
	if n.entry == nil || (*n.entry).IsDir() {
		dirs, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("ReadDir!!:", err)
		}
		return dirs
	} else {
		return nil
	}

}

type Tree struct {
	Root Node
}

func (t *Tree) GetNode(index int) *Node {
	n := &t.Root.Childrens[0]
	for i := 1; i < index; i++ {
		n = &n.Childrens[0]
	}
	return n
}

func (t *Tree) FillParents() {
	t.Root.FillParents(nil)
}

func NewTree() (t *Tree, index int) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	homeNode := NewNodeFully(home)
	var tempNode Node
	tempPath := home
	oldPath := tempPath
	count := 0
	for {
		tempPath = filepath.Dir(home)
		if tempPath != oldPath {
			tempNode = NewNode(tempPath, nil)
			tempNode.Childrens = append(tempNode.Childrens, homeNode)
			oldPath = tempPath
			count++
		} else {
			break
		}
	}

	tree := Tree{
		Root: tempNode,
	}
	tree.FillParents()

	return &tree, count
}

type Load int

const (
	Fully Load = iota
	Partial
	Unloaded
)

type Node struct {
	Childrens []Node
	parent    *Node
	load      Load
	path      string
	entry     *fs.DirEntry
	entries   *[]fs.DirEntry
	Focus     int
	// value     fs.DirEntry
}

func (n *Node) FillParents(parent *Node) {
	n.parent = parent
	// n.FileInfo = fi
	for i := 0; i < len(n.Childrens); i++ {
		n.Childrens[i].FillParents(n)
	}
}

func NewNodeFully(path string) Node {
	n := Node{
		path:  path,
		load:  Fully,
		Focus: 0,
	}
	entrys := getEntries(path, &n)
	nodes := make([]Node, 0, len(entrys))
	for i := 0; i < len(entrys); i++ {
		nodes = append(nodes, NewNode(path+"/"+entrys[i].Name(), &entrys[i]))
	}
	n.entries = &entrys
	n.Childrens = nodes
	return n
}

func NewNode(path string, entry *fs.DirEntry) Node {
	return Node{
		Childrens: make([]Node, 0),
		path:      path,
		load:      Unloaded,
		entry:     entry,
		Focus:     0,
	}
}

func (n *Node) Entries() []fs.DirEntry {
	length := len(n.Childrens)
	entries := make([]fs.DirEntry, 0, length)
	for i := 0; i < length; i++ {
		entries = append(entries, *n.Childrens[i].entry)
	}
	return entries
}

// func (Node) IsLeaf() bool {
// 	return false
// }

func (n *Node) Parent() *Node {
	return n.parent
}

func (n *Node) FousedNode() *Node {
	if len(n.Childrens) > 0 {
		return &n.Childrens[n.Focus]
	}
	return nil
}

func (n *Node) Load() {
	if n.load == Unloaded {
		entrys := getEntries(n.path, n)
		nodes := make([]Node, 0, len(entrys))
		for i := 0; i < len(entrys); i++ {
			nodes = append(nodes, NewNode(n.path+"/"+entrys[i].Name(), &entrys[i]))
		}
		n.Childrens = nodes
		n.entries = &entrys
		n.load = Fully
	} else if n.load == Partial {
		entrys := getEntries(n.path, n)
		paths := make([]string, 0, len(n.Childrens))
		// TODO 3 Maybe this code could be cleaner and more performance
		for i := 0; i < len(n.Childrens); i++ {
			paths = append(paths, n.Childrens[i].path)
		}
		for i := 0; i < len(entrys); i++ {
			if slices.Contains(paths, n.path+"/"+entrys[i].Name()) {
				n.Childrens = append(n.Childrens, NewNode(n.path+"/"+entrys[i].Name(), &entrys[i]))
			}
		}
		n.entries = &entrys
		n.load = Fully
	}
}

func (n *Node) AsyncLoad(deep int) *sync.WaitGroup {
	w := &sync.WaitGroup{}

	w.Add(1)
	n.asyncLoad(deep, w)

	return w
}

func (n *Node) asyncLoad(deep int, w *sync.WaitGroup) *sync.WaitGroup {
	if w == nil {
		w = &sync.WaitGroup{}
	}

	defer w.Done()

	if deep < 0 {
		return w
	}

	if n.load == Partial || n.load == Unloaded {
		n.Load()
	}
	for i := 0; i < len(n.Childrens); i++ {
		w.Add(1)
		go n.Childrens[i].asyncLoad(deep-1, w)
	}
	return w
}
