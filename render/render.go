package render

import (
	"errors"
	"fmt"

	"github.com/nsf/termbox-go"
)

const ErrorHeightOverflow = "Height Overflow"
const ErrorWidthOverflow = "Width Overflow"

func RString(s string, line int, bg termbox.Attribute) error {
	w, h := termbox.Size()
	if h < line {
		return errors.New(ErrorHeightOverflow)
	}
	if w < len(s) {
		return errors.New(ErrorWidthOverflow)
	}
	for i := 0; i < len(s); i++ {
		termbox.SetBg(i, line, bg)
		termbox.SetChar(i, line, rune(s[i]))
	}
	return nil
}

func Clear() {
	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			termbox.SetChar(x, y, 0)
			termbox.SetBg(x, y, termbox.ColorDefault)
		}
	}
}

func CleanScreen() {
	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := termbox.GetCell(x, y)
			c.Bg = termbox.ColorDefault
			c.Ch = 0
			c.Fg = termbox.ColorDefault
		}
	}
}

func Debug(line int, a ...any) {
	// return
	w, _ := termbox.Size()
	s := fmt.Sprint(a...)
	length := len(s)
	for i := 0; i < len(s); i++ {
		termbox.SetChar(w-length+i, line, rune(s[i]))
	}
	termbox.Flush()
	// time.Sleep(1 * time.Millisecond)
}
