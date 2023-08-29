package render

import (
	"sync"
	"testing"

	"github.com/nsf/termbox-go"
)

func TestTree(t *testing.T) {
	tree, index := NewTree()

	base := tree.Root.Childrens[0]
	for i := 1; i < index; i++ {
		base = base.Childrens[0]
	}
	w := base.AsyncLoad(1)
	w.Wait()
	for i := 1; i < len(base.Childrens); i++ {
		println(base.Childrens[i].path)
		printNode(base.Childrens[i], 1)
	}
}

func printNode(n Node, level int) {
	tab := ""
	for i := 0; i < level; i++ {
		tab += "  "
	}
	for i := 0; i < len(n.Childrens); i++ {
		if n.entries == nil {
			panic("R")
		}
		printNode(n.Childrens[i], level+1)
	}
}

func BenchmarkTree(b *testing.B) {
	tree, index := NewTree()

	base := tree.Root.Childrens[0]
	for i := 1; i < index; i++ {
		base = base.Childrens[0]
	}
	w := base.AsyncLoad(22)
	w.Wait()
	for i := 1; i < len(base.Childrens); i++ {
		// println(base.Childrens[i].path)
		// printNode(base.Childrens[i], 1)
	}
}

func BenchmarkRenderString(b *testing.B) {
	termbox.Init()
	strs := []string{
		"Comida",
		"Gonzalo",
		"MasidasDULCE",
		"SENSACIONES",
		"miamor",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
		"Caballero",
		"Te quiro",
		"SADASADASD",
		"MIZZAasp",
		"ZGYUAYUASs",
	}
	for j := 0; j < 5; j++ {
		uwg := sync.WaitGroup{}
		for i := 0; i < len(strs)-2; {
			wg := sync.WaitGroup{}
			wg.Add(3)
			WriteString(strs[i], i, 0, termbox.ColorDefault, &wg)
			WriteString(strs[i], i, 13, termbox.ColorDefault, &wg)
			WriteString(strs[i], i, 26, termbox.ColorDefault, &wg)
			i += 3
			wg.Wait()
			uwg.Wait()
			uwg.Add(1)
			go UpdateScreen(&uwg)
		}
	}
}
