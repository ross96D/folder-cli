package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"ross/fastfolder/render"
	"strings"

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

const items = 300
const ShowInvisible = false

var list render.ListItem

func Nav(cmd *cobra.Command, args []string) {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Error termbox init:", err)
	}
	defer termbox.Close()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}
	dirs := getDirs(home)
	list = render.NewList()
	for i := 0; i < len(dirs); i++ {
		dir := dirs[i]
		list.AddItem(&render.Item{
			Text: dir.Name(),
		})
	}
	list.Items[0].Focus = true
	// time.Sleep(300 * time.Millisecond)
	list.Draw()

	handleMain()
}

func handleMain() {
	e := termbox.PollEvent()
	if e.Type != termbox.EventKey {
		handleMain()
		return
	}
	if e.Ch == 'q' {
		termbox.Close()
		return
	}
	if e.Key == termbox.KeyArrowDown {
		list.Focus(1)
	} else if e.Key == termbox.KeyArrowUp {
		list.Focus(-1)
	}
	handleMain()
}

func getDirs(home string) []fs.DirEntry {
	dirs, err := os.ReadDir(home)
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
		if len(*dirs) < items {
			s = *dirs
		} else {
			s = (*dirs)[0:items]
		}
	} else {
		s = make([]fs.DirEntry, 0, items)
		count := 0
		for i := 0; count < items && i < len(*dirs); i++ {
			dir := (*dirs)[i]
			if !strings.HasPrefix(dir.Name(), ".") {
				count++
				s = append(s, dir)
			}
		}
	}
	return s
}
