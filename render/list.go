package render

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/nsf/termbox-go"
)

const Items = 300
const ShowInvisible = false

type StateList uint8

const (
	fullRender StateList = iota
	noFullRender
)

type Item struct {
	Name     string
	Focus    bool
	IsDir    bool
	Callback func()
}

type ListItem struct {
	Items         []*Item
	ritemStart    int
	ritemEnd      int
	IFocus        int
	state         StateList
	canFullRender bool
}

func NewList() ListItem {
	return ListItem{
		Items: make([]*Item, 0),
	}
}

func (l *ListItem) Draw() {
	_, h := termbox.Size()
	if h > len(l.Items) {
		l.canFullRender = true
	} else {
		l.canFullRender = false
	}
	if !l.canFullRender && len(l.Items)-h < l.ritemStart {
		l.ritemStart = len(l.Items) - h
	}
	var err error
	Debug(0, "start ", l.ritemStart, " ", len(l.Items), " ", l.ritemEnd)
	for e := l.ritemStart; e < len(l.Items); e++ {
		i := e - l.ritemStart
		if l.Items[e].Focus {
			Debug(1, "ifocus", l.IFocus)
			Debug(4, l.Items[e].Name, " ", e)
			err = RString(l.Items[e].Name, i, termbox.ColorGreen)
		} else {
			Debug(8, "ZZ ", e, " ", l.IFocus)
			err = RString(l.Items[e].Name, i, termbox.ColorDefault)
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
	if l.IFocus == 0 && l.Items[0].Focus == false {
		l.Items[0].Focus = true
	} else {
		if dir == 1 {
			if len(l.Items)-1 == l.IFocus {
				return
			}
			l.Items[l.IFocus].Focus = false
			Debug(6, l.IFocus)
			l.IFocus = l.IFocus + 1
			l.Items[l.IFocus].Focus = true
			Debug(7, l.IFocus)

		} else if dir == -1 {
			if l.IFocus == 0 {
				return
			}
			l.Items[l.IFocus].Focus = false
			l.IFocus = l.IFocus - 1
			l.Items[l.IFocus].Focus = true
		}
	}
	if l.IFocus < l.ritemStart {
		if l.ritemStart > 0 {
			l.ritemStart -= 1
			l.ritemEnd -= 1
		}
	} else if !l.canFullRender && l.IFocus+2 > l.ritemEnd {
		l.ritemStart += 1
		l.ritemEnd += 1
	}
	Clear()
	l.Draw()
}

func (l *ListItem) clear() {
	l.Items = l.Items[:0]
	l.IFocus = 0
	Clear()
	l.ritemStart = 0
}

func (l *ListItem) Repopulate(path string, deb bool, focus int) {
	l.clear()
	entrys := getDirEntry(path)
	for i := 0; i < len(entrys); i++ {
		entry := (entrys)[i]
		l.AddItem(&Item{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}
	if deb && focus == 0 {
		l.Items[0].Focus = true
		for i := 0; i < len(l.Items); i++ {
			Debug(3, *l.Items[i])
		}
	}
	l.IFocus = focus
	l.Items[l.IFocus].Focus = true
	l.Draw()
}

func getDirEntry(path string) []fs.DirEntry {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("ReadDir:", err)
	}
	// cmd_ := exec.Command("gnome-terminal", "--", "bash", "-c", "cd "+home+" && exec $SHELL")
	// err = cmd_.Run()

	return selectFolders2Show(&dirs)

}

func selectFolders2Show(dirs *[]fs.DirEntry) []fs.DirEntry {
	var s []fs.DirEntry
	if ShowInvisible {
		if len(*dirs) < Items {
			s = *dirs
		} else {
			s = (*dirs)[0:Items]
		}
	} else {
		s = make([]fs.DirEntry, 0, Items)
		count := 0
		for i := 0; count < Items && i < len(*dirs); i++ {
			dir := (*dirs)[i]
			if !strings.HasPrefix(dir.Name(), ".") {
				count++
				s = append(s, dir)
			}
		}
	}
	return s
}
