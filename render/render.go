package render

import (
	"errors"

	"github.com/nsf/termbox-go"
)

func RString(s string, line int) error {
	w, h := termbox.Size()
	if h < line {
		return errors.New("Height Overflow")
	}
	if w < len(s) {
		return errors.New("Width Overflow")
	}
	for i := 0; i < len(s); i++ {
		termbox.SetChar(i, line, rune(s[i]))
	}
	return nil
}

func Clear() {
	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			termbox.SetChar(x, y, 0)
		}
	}
}
