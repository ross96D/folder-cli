package cmd

import (
	"fmt"
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

var nav render.NavTree
var isOpen bool

func Nav(cmd *cobra.Command, args []string) {
	err := termbox.Init()
	isOpen = true
	if err != nil {
		fmt.Println("Error termbox init:", err)
	}
	defer termbox.Close()

	nav = render.NewNavTree()
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
	isHandle := nav.HandleEvent(e)
	if !isHandle {

	}
}
