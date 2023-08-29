package render

import (
	"errors"
	"sync"

	"github.com/nsf/termbox-go"
)

// ! Considerar otra capa de abstraccion.
//
// ! La idea seria algo entre el render y el Node que se encarge de la cache,
// ! de donde a donde renderizar y de saber si cambio la referencia
//
// ! Solo existiria (por ahora) una sola instancia q se alimenta del nodo actual del NavTree
// ! Con esto ademas podemos implementar el recibo de eventos de termobox de forma asyncrona
// ! Pero hay q verificar los problemas de concurrencia q esto podria tener
//
// ! Performance: Hacer el flush en una goroutine.. tal vez esto no cause problemas de concurrencia, aunque seria raro...
func RenderNode(n *Node) {
	for i := 0; i < len(*n.entries); i++ {
		WriteString((*n.entries)[i].Name(), i, 0, termbox.ColorDefault, nil)
	}
}

const ErrorHeightOverflow = "height Overflow"
const ErrorWidthOverflow = "width Overflow"

func WriteString(s string, line int, start int, bg termbox.Attribute, wg *sync.WaitGroup) error {
	defer wg.Done()
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
	termbox.Flush()
	return nil
}

func CleanCell(x, y int) {
	termbox.SetChar(x, y, 0)
	termbox.SetBg(x, y, termbox.ColorDefault)
	termbox.SetFg(x, y, termbox.ColorDefault)
}

func UpdateScreen(wg *sync.WaitGroup) {
	defer wg.Done()
	termbox.Flush()
}

// func Debug(line int, a ...any) {
// 	// return
// 	w, _ := termbox.Size()
// 	s := fmt.Sprint(a...)
// 	length := len(s)
// 	for i := 0; i < len(s); i++ {
// 		termbox.SetChar(w-length+i, line, rune(s[i]))
// 	}
// 	termbox.Flush()
// }
