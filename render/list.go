package render

import (
	"github.com/nsf/termbox-go"
)

type Item struct {
	Text     string
	Callback func()
}

type ListItem struct {
	items []*Item
}

func NewList() ListItem {
	return ListItem{
		items: make([]*Item, 0),
	}
}

func (l *ListItem) Draw() {
	for i := 0; i < len(l.items); i++ {
		err := RString(l.items[i].Text, i)
		if err != nil {
			// fmt.Println("Error:", err)
		}
	}
	termbox.Flush()
}

func (l *ListItem) AddItem(i *Item) {
	l.items = append(l.items, i)
}
