package render

import (
	"github.com/nsf/termbox-go"
)

type StateList uint8

const (
	initial StateList = iota
	onRender
)

type Item struct {
	Text     string
	Focus    bool
	Callback func()
}

type ListItem struct {
	Items      []*Item
	ritemStart int
	ritemEnd   int
	iFocus     int
	state      StateList
}

func NewList() ListItem {
	return ListItem{
		Items: make([]*Item, 0),
	}
}

func (l *ListItem) Draw() {
	// if l.state == initial {
	// 	l.ritemStart = 0
	// 	l.state = onRender
	// }

	var err error
	Debug(0, "start", l.ritemStart)
	for e := l.ritemStart; e < len(l.Items); e++ {
		i := e - l.ritemStart
		if l.Items[e].Focus {
			Debug(1, "ifocus", l.iFocus)
			Debug(4, l.Items[e].Text, " ", e)
			err = RString(l.Items[e].Text, i, termbox.ColorGreen)
		} else {
			Debug(8, "ZZ ", e, " ", l.iFocus)
			err = RString(l.Items[e].Text, i, termbox.ColorDefault)
		}

		if err != nil {
			if err.Error() == ErrorHeightOverflow {
				Debug(5, "error", e)
				l.ritemEnd = e
			}
			break
		}
	}

	if err == nil {
		l.ritemEnd = len(l.Items)
	}
	termbox.Flush()
}

func (l *ListItem) AddItem(i *Item) {
	l.Items = append(l.Items, i)
}

// dir 1 is down, -1 is up. The list start from the top to bottom
func (l *ListItem) Focus(dir int) {
	if l.iFocus == 0 && l.Items[0].Focus == false {
		l.Items[0].Focus = true
	} else {
		if dir == 1 {
			if len(l.Items)-1 == l.iFocus {
				return
			}
			l.Items[l.iFocus].Focus = false
			Debug(6, l.iFocus)
			l.iFocus = l.iFocus + 1
			l.Items[l.iFocus].Focus = true
			Debug(7, l.iFocus)

		} else if dir == -1 {
			if l.iFocus == 0 {
				return
			}
			l.Items[l.iFocus].Focus = false
			l.iFocus = l.iFocus - 1
			l.Items[l.iFocus].Focus = true
		}
	}
	Debug(2, "end", l.ritemEnd)
	if l.iFocus < l.ritemStart {
		if l.ritemStart > 0 {
			l.ritemStart -= 1
			l.ritemEnd -= 1
		}
	} else if l.iFocus+5 > l.ritemEnd {
		l.ritemStart += 1
		l.ritemEnd += 1
	}
	Clear()
	l.Draw()
}
