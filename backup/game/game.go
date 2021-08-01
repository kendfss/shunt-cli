// package game // import "tildegit.org/eli2and40/rube/game"
// package game // import "rube/game"
package main

import (
	"fmt"
	"unicode/utf8"
	// "strings"
	
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
	
	// "tildegit.org/eli2and40/rube/board"
	"rube/board"
	st"rube/stringtools"
)

type Game struct {
	board *board.Board
	slide bool
	moves, unmoves []string
	moveCount, tileWidth, tileHeight int // 0, 3, 3
}
func (self Game) findMove(move string) func(*board.Board) {
	var moves = map[string]func(*board.Board) {
		"FlipVertical": (*board.Board).FlipVertical,
		"FlipHorizontal": (*board.Board).FlipHorizontal,
		"SlideUp": (*board.Board).SlideDown,
		"SlideDown": (*board.Board).SlideUp,
		"SlideLeft": (*board.Board).SlideRight,
		"SlideRight": (*board.Board).SlideLeft,
	}
	return moves[move]
}
func (self Game) negateMove(move string) func(*board.Board) {
	var negs = map[string]func(*board.Board) {
		"FlipVertical": (*board.Board).FlipVertical,
		"FlipHorizontal": (*board.Board).FlipHorizontal,
		"SlideUp": (*board.Board).SlideDown,
		"SlideDown": (*board.Board).SlideUp,
		"SlideLeft": (*board.Board).SlideRight,
		"SlideRight": (*board.Board).SlideLeft,
	}
	return negs[move]
}
func (self Game) Count() int {
	return self.moveCount + len(self.moves)
}
func (self *Game) Undo() {
	if len(self.moves) > 0 {
		
		move := self.moves[len(self.moves)-1]
		self.negateMove(move)(self.board)
		
		self.moves = self.moves[:len(self.moves)-1]
		self.unmoves = append(self.unmoves, move)
		self.moveCount += 2
	}
}
func (self *Game) Redo() {
	if len(self.unmoves) > 0 {
		move := self.unmoves[0]
		self.findMove(move)(self.board)
		self.unmoves = self.unmoves[1:]
		moves := []string{}
		moves = append(moves, move)
		moves = append(moves, self.moves...)
		self.moves = moves
	}
}
func (self *Game) ActionUp() {
	var move string
	if self.slide {
		self.board.SlideUp()
		move = "SlideUp"
	} else {
		self.board.FlipVertical()
		move = "FlipVertical"
	}
	self.moves = append(self.moves, move)
	self.unmoves = []string{}
}
func (self *Game) ActionDown() {
	var move string
	if self.slide {
		self.board.SlideDown()
		move = "SlideDown"
	} else {
		self.board.FlipVertical()
		move = "FlipVertical"
	}
	self.moves = append(self.moves, move)
	self.unmoves = []string{}
}
func (self *Game) ActionLeft() {
	var move string
	if self.slide {
		self.board.SlideLeft()
		move = "SlideLeft"
	} else {
		self.board.FlipHorizontal()
		move = "FlipHorizontal"
	}
	self.moves = append(self.moves, move)
	self.unmoves = []string{}
}
func (self *Game) ActionRight() {
	var move string
	if self.slide {
		self.board.SlideRight() 
		move = "SlideRight"
	} else {
		self.board.FlipHorizontal()
		move = "FlipHorizontal"
	}
	self.moves = append(self.moves, move)
	self.unmoves = []string{}
}
func (self *Game) MoveUp() {
	self.board.MoveUp()
}
func (self *Game) MoveDown() {
	self.board.MoveDown()
}
func (self *Game) MoveLeft() {
	self.board.MoveLeft()
}
func (self *Game) MoveRight() {
	self.board.MoveRight() 
}
func (self *Game) Reset() {
	self.moves = []string{}
	self.unmoves = []string{}
	self.moveCount = 0
	self.board.Reset()
}
func (self *Game) ToggleSlideFlip() {
	self.slide = !self.slide
}
func (self Game) Solved() bool {
	return self.board.Solved()
}
func (self Game) Dim() (int, int) {
	return self.board.Width(), self.board.Height()
}
func (self *Game) NewBoard(board board.Board) {
	self.moves = []string{}
	self.unmoves = []string{}
	self.moveCount = 0
	self.board = &board
}
func (self *Game) GoTo(x, y int) {
	self.board.MoveTo(x, y)
}
func (self Game) Solution() Game {
	sol := self.board.Solution()
	return Game{&sol, self.slide, self.moves, self.unmoves, self.moveCount, self.tileWidth, self.tileHeight}
}
func (self Game) Save() {
}
func (self *Game) New() {
	new := board.NewBoard(self.Dim())
	new.Shuffle()
	self.NewBoard(new)
}
func (self Game) Display(message string) {
	
}
func (self *Game) NextPanel() {
	
}
func (self Game) Congratulate() {
	
}
func (self *Game) Loop() {
	
}



func redraw_all(g Game) {
	// width := len(st.Lines(self.board.String())[0])
	width, height := g.Dim()
	width = len(st.Lines(g.board.String())[0])
	// width *= g.tileWidth
	height *= g.tileHeight
	
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := h / 2
	midx := (w - width) / 2

	// unicode box drawing chars around the edit box
	// if runewidth.EastAsianWidth {
	// 	termbox.SetCell(midx-1, midy, '|', coldef, coldef)
	// 	termbox.SetCell(midx+width, midy, '|', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy-1, '+', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy+1, '+', coldef, coldef)
	// 	termbox.SetCell(midx+width, midy-1, '+', coldef, coldef)
	// 	termbox.SetCell(midx+width, midy+1, '+', coldef, coldef)
	// 	fill(midx, midy-1, width, 1, termbox.Cell{Ch: '-'})
	// 	fill(midx, midy+1, width, 1, termbox.Cell{Ch: '-'})
	// } else {
	// 	termbox.SetCell(midx-1, midy, '│', coldef, coldef)
	// 	termbox.SetCell(midx+width, midy, '│', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy-1, '┌', coldef, coldef)
	// 	termbox.SetCell(midx-1, midy+1, '└', coldef, coldef)
	// 	termbox.SetCell(midx+width, midy-1, '┐', coldef, coldef)
	// 	termbox.SetCell(midx+width, midy+1, '┘', coldef, coldef)
	// 	fill(midx, midy-1, width, 1, termbox.Cell{Ch: '─'})
	// 	fill(midx, midy+1, width, 1, termbox.Cell{Ch: '─'})
	// }

	for i, line := range st.Lines(g.board.String()) {
		shift := (width - st.RuneLen(line))/2
		tbprint(midx+shift, midy+i, coldef, coldef, line)
	}
	termbox.Flush()
}


func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func rune_advance_len(r rune, pos int) int {
	if r == '\t' {
		return tabstop_length - pos%tabstop_length
	}
	return runewidth.RuneWidth(r)
}



const tabstop_length = 8
var messages = []string{"Press Ctrl+Q to quit"}




func NewGame(width, height int) Game {
	board := board.NewBoard(width, height)
	return Game{&board, true, []string{}, []string{}, 0, 3, 3}
}



func byte_slice_grow(s []byte, desired_cap int) []byte {
	if cap(s) < desired_cap {
		ns := make([]byte, len(s), desired_cap)
		copy(ns, s)
		return ns
	}
	return s
}


func byte_slice_insert(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = byte_slice_grow(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

func legis(r rune) {
	l := len(messages)
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	// messages = append(messages, byte_slice_insert(eb.text, eb.cursor_boffset, buf[:n]))
	messages = append(messages, st.Read(byte_slice_insert([]byte{}, 0, buf[:n])...))
	if len(messages) <= l {
		panic("append failed")
	}
	// messages = append(messages, st.Read(n))
	
}


func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
		for {
			fmt.Println("errorloop")
		}
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)
	
	g := NewGame(3, 3)
	g.board.Shuffle()
	
	redraw_all(g)
	loop:
		for {
			if !g.Solved() {
				switch ev := termbox.PollEvent(); ev.Type {
					case termbox.EventKey:
						switch ev.Key {
							// case termbox.KeyEsc, termbox.KeyEnter:
							case termbox.KeyEsc, termbox.KeyCtrlC:
								break loop
							case termbox.KeyArrowLeft:
								// self.ActionLeft()
								g.ActionLeft()
							case termbox.KeyArrowRight:
								// self.ActionRight()
								g.ActionRight()
							case termbox.KeyArrowUp:
								// self.ActionUp()
								g.ActionUp()
							case termbox.KeyArrowDown:
								// self.ActionDown()
								g.ActionDown()
							case termbox.KeyCtrlTilde, termbox.KeySpace:
								// self.ToggleSlideFlip()
								g.ToggleSlideFlip()
							case termbox.KeyCtrlN:
								// new := board.NewBoard(self.Dim())
								new := board.NewBoard(g.Dim())
								new.Shuffle()
								// self.NewBoard(new)
								g.NewBoard(new)
							case termbox.KeyCtrlZ:
								// self.Undo()
								g.Undo()
							case termbox.KeyCtrlY:
								// self.Redo()
								g.Redo()
							case termbox.KeyCtrlS:
								// self.Save()
								g.Save()
							case termbox.KeyCtrlR:
								// self.Reset()
								g.Reset()
							case termbox.KeyTab:
								// self.NextPanel()
								g.NextPanel()
						}
						switch ev.Ch {
							case 'w':
								g.MoveUp()
							case 'a':
								g.MoveLeft()
							case 's':
								g.MoveDown()
							case 'd':
								g.MoveRight()
							case 'n':
								// new := board.NewBoard(self.Dim())
								new := board.NewBoard(g.Dim())
								new.Shuffle()
								// self.NewBoard(new)
								g.NewBoard(new)
						}
					case termbox.EventError:
						panic(ev.Err)
				}
			} else {
				g.Congratulate()
			}
			// self.redraw_all()
			redraw_all(g)
		}
}
