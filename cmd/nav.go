package cmd

import (
	"fmt"
	"os"
	"ross/fastfolder/render"

	"github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

var NavCmd = &cobra.Command{
	Use:   "nav",
	Short: "To navigate",
	Long:  "I dont care",
	Run: func(cmd *cobra.Command, args []string) {
		Nav(cmd, args)
	},
}

var list render.ListItem
var nav render.NavStack
var isOpen bool

func Nav(cmd *cobra.Command, args []string) {
	err := termbox.Init()
	isOpen = true
	if err != nil {
		fmt.Println("Error termbox init:", err)
	}
	defer termbox.Close()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}
	nav = render.NewNavStack()
	nav.Push(home, 0)

	list = render.NewList()
	list.Repopulate(nav.Get(), false, 0)
	// time.Sleep(300 * time.Millisecond)
	for isOpen {
		handleMain()
	}
}

func handleMain() {
	e := termbox.PollEvent()
	if e.Type != termbox.EventKey {
		return
	}
	if e.Ch == 'q' || e.Key == termbox.KeyCtrlC {
		termbox.Close()
		isOpen = false
		return
	}
	if e.Key == termbox.KeyArrowDown {
		list.Focus(1)
	} else if e.Key == termbox.KeyArrowUp {
		list.Focus(-1)
	} else if e.Key == termbox.KeyArrowRight {
		if list.Items[list.IFocus].IsDir {
			currpath := nav.Get() + "/" + list.Items[list.IFocus].Name
			nav.Push(currpath, list.IFocus)
			list.Repopulate(nav.Get(), true, 0)
		}
	} else if e.Key == termbox.KeyArrowLeft {
		nav.Pop()
		list.Repopulate(nav.Get(), true, nav.GetFocus())
	}
}
