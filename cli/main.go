// package game // import "tildegit.org/eli2and40/rube/game"
package main

import (
	"fmt"
	// "tildegit.org/eli2and40/rube/board"
	"rube/board"
)

type Game struct {
	board *board.Board
	slide bool
	moves []string
	unmoves []string
	moveCount int
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
	return Game{&sol, self.slide, self.moves, self.unmoves, self.moveCount}
}
func (self Game) Save() {
}
func (self Game) Display(message string) {
	
}
func (self *Game) Loop() {
	
}















func NewGame(width, height int) Game {
	board := board.NewBoard(width, height)
	return Game{&board, true, []string{}, []string{}, 0}
}












func main() {
	maxw, maxl := 10, 10
	
	for height:=1; height<maxl; height++ {
		for width:=1; width<maxw; width++ {
			g := NewGame(width, height)
			for y:=0; y<height; y++ {
				for x:=0; x<width; x++ {
					g.GoTo(x, y)
					g = g.Solution()
					
					// fmt.Println(g.board, "start")
					
					g.ActionUp()
					// fmt.Println(g.board, "up")
					g.ToggleSlideFlip()
					g.ActionDown()
					// fmt.Println(g.board, "down")
					g.ToggleSlideFlip()
					g.ActionLeft()
					// fmt.Println(g.board, "left")
					g.ToggleSlideFlip()
					g.ActionRight()
					// fmt.Println(g.board, "right")
					if g.Count() != 4 {
						panic(fmt.Sprintf("Counting error!(action)\n\thave\t%v\n\twant\t%v\n%s\n", g.Count(), 4, g.board))
					}
					
					// fmt.Println(g.board)
					g.Undo()
					// fmt.Println(g.board)
					g.Undo()
					// fmt.Println(g.board)
					g.Undo()
					// fmt.Println(g.board)
					g.Undo()
					// fmt.Println(g.board)
					if g.Count() != 8 {
						panic(fmt.Sprintf("Counting error!(undoing)\n\thave\t%v\n\twant\t%v\n%s\n", g.Count(), 8, g.board))
					}
					if len(g.unmoves) != 4 {
						panic(fmt.Sprintf("Counting error!(unmoves)\n\thave\t%v\n\twant\t%v\n%s\n", len(g.unmoves), 4, g.board))
					}
					if !g.Solved() {
						panic(fmt.Sprintf("Undo error!\n%s\n", g.board))
					}
					g.Redo()
					g.Redo()
					g.Redo()
					g.Redo()
					if g.Count() != 12 {
						panic(fmt.Sprintf("Counting error!(redoing)\n\thave\t%v\n\twant\t%v\n%s\n", g.Count(), 12, g.board))
					}
					if len(g.unmoves) != 0 {
						panic(fmt.Sprintf("Counting error!(unmoves)\n\thave\t%v\n\twant\t%v\n%s\n", len(g.unmoves), 0, g.board))
					}
					
					g.Reset()
					g.GoTo(0, 0)
					// fmt.Println()
					// fmt.Println()
					// fmt.Println()
					// fmt.Println()
				}
			}
		}
	}
	// g := NewGame(3, 1)
	// g.GoTo(1, 0)
	// fmt.Println(g.board, g.Solved(), g.Count(), g.moves, "\n")
	
	// fmt.Println("Doing")
	// g.ActionUp()
	// fmt.Println(g.board, "up")
	// g.ToggleSlideFlip()
	// g.ActionDown()
	// fmt.Println(g.board, "down")
	// g.ToggleSlideFlip()
	// g.ActionLeft()
	// fmt.Println(g.board, "left")
	// g.ToggleSlideFlip()
	// g.ActionRight()
	// fmt.Println(g.board, "right")
	
	// fmt.Println("Undoing")
	// fmt.Println(g.board)
	// g.Undo()
	// fmt.Println(g.board)
	// g.Undo()
	// fmt.Println(g.board)
	// g.Undo()
	// fmt.Println(g.board)
	// g.Undo()
	// fmt.Println(g.board)
	
	// fmt.Println("Over-undoing")
	// g.Undo()
	// fmt.Println(g.board, g.Solved(), g.Count(), g.moves, "\n")
	
	// fmt.Println("Redoing")
	// g.Redo()
	// g.Redo()
	// g.Redo()
	// g.Redo()
	
	// fmt.Println("Over-redoing")
	// g.Redo()
	// fmt.Println(g.board, g.Solved(), g.Count(), g.moves, "\n")
}
