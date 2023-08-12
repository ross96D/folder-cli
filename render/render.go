package render

import (
	"errors"
	"fmt"

	"github.com/nsf/termbox-go"
)

const ErrorHeightOverflow = "Height Overflow"
const ErrorWidthOverflow = "Width Overflow"

func RString(s string, line int, start int, bg termbox.Attribute) error {
	w, h := termbox.Size()
	if h <= line {
		return errors.New(ErrorHeightOverflow)
	}
	if w <= len(s) {
		return errors.New(ErrorWidthOverflow)
	}
	for i := start; i < (len(s) + start); i++ {
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
			CleanCell(x, y)
		}
	}
}

func CleanBlock(lineStart int, lineEnd int, charStart int, charEnd int) error {
	w, h := termbox.Size()
	if h <= lineEnd {
		return errors.New(ErrorHeightOverflow)
	} else if w <= charEnd {
		return errors.New(ErrorWidthOverflow)
	}
	for y := lineStart; y <= lineEnd; y++ {
		for x := charStart; x < charEnd; x++ {
			CleanCell(x, y)
		}
	}
	return nil
}

func CleanCell(x, y int) {
	c := termbox.GetCell(x, y)
	c.Bg = termbox.ColorDefault
	c.Ch = 0
	c.Fg = termbox.ColorDefault
}

func UpdateScreen() {
	termbox.Flush()
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
	// time.Sleep(3 * time.Second)
}
