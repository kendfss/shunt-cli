package game

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"

	"github.com/kendfss/shunt-cli/board"
	nt "github.com/kendfss/shunt-cli/numtools"
	st "github.com/kendfss/shunt-cli/stringtools"
)

type Game struct {
	board                            *board.Board
	slide                            bool
	moves, unmoves                   []string
	moveCount, tileWidth, tileHeight int // 0, 3, 3
	lastTime                         time.Time
	elapsedTime                      time.Duration
	clockStop                        bool
}

func (self Game) findMove(move string) func(*board.Board) {
	moves := map[string]func(*board.Board){
		"FlipVertical":   (*board.Board).FlipVertical,
		"FlipHorizontal": (*board.Board).FlipHorizontal,
		"SlideUp":        (*board.Board).SlideDown,
		"SlideDown":      (*board.Board).SlideUp,
		"SlideLeft":      (*board.Board).SlideRight,
		"SlideRight":     (*board.Board).SlideLeft,
	}
	return moves[move]
}

func (self Game) negateMove(move string) func(*board.Board) {
	negs := map[string]func(*board.Board){
		"FlipVertical":   (*board.Board).FlipVertical,
		"FlipHorizontal": (*board.Board).FlipHorizontal,
		"SlideUp":        (*board.Board).SlideDown,
		"SlideDown":      (*board.Board).SlideUp,
		"SlideLeft":      (*board.Board).SlideRight,
		"SlideRight":     (*board.Board).SlideLeft,
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
	return Game{&sol, self.slide, self.moves, self.unmoves, self.moveCount, self.tileWidth, self.tileHeight, time.Now(), time.Second * 0, false}
}

//	func (self Game) timeDelta() time.Duration {
//		return self.lastTime.Sub(time.Now())
//	}
func (self *Game) updateClock() {
	t := time.Now()
	// self.elapsedTime += self.lastTime.Sub(t)
	self.elapsedTime += t.Sub(self.lastTime)
	self.lastTime = t
	self.Draw()
}

func (self Game) clock() {
	for {
		// select
		if !self.clockStop {
			self.updateClock()
			self.Draw()
			time.Sleep(time.Second)
		} else {
			return
		}
	}
}

func (self Game) Save() {
}

func (self *Game) Load() {
	self.lastTime = time.Now()
}

func (self *Game) New() {
	self.clockStop = true
	*self = NewGame(self.Dim())
}

func (self Game) ScreenSize() (int, int) {
	return termbox.Size()
}

func (self Game) ScreenCenter() (int, int) {
	w, h := termbox.Size()
	return w / 2, h / 2
}

func (self Game) Display(message string) {
	width, _ := self.ScreenSize()
	fg := termbox.RGBToAttribute(22, 22, 22)
	bg := termbox.RGBToAttribute(255, 255, 255)
	lines := st.Lines(message)
	for i := range lines {
		line := st.Center(lines[i], width)
		x, y := self.ScreenCenter()
		x -= st.RuneLen(line) / 2
		y += len(st.Lines(self.repr())) / 2
		y -= (len(lines) - 1)
		tbprint(x, y+i, fg, bg, line)
	}
	termbox.Flush()
}

func (self *Game) NextPanel() {
}

func (self Game) Congratulate() {
	self.Display(celebrations[nt.Randex(len(celebrations))])
	ctx, exit := context.WithTimeout(context.TODO(), 5*time.Second)
	defer exit()
	go func() {
		for {
			select {
			case <-ctx.Done():
				termbox.Interrupt()
				return
			default:
				continue
			}
		}
	}()
	termbox.PollEvent()
}

func (self Game) DrawPrompt(question string, labels []string, index int) {
	const coldef = termbox.ColorDefault

	box_width := 30
	w, h := self.ScreenSize()
	midy := h / 2
	midx := (w - box_width) / 2

	question = st.CenterWrap(question, box_width-4)
	lines := st.Lines(question)

	text := ""

	for i := range lines {
		text += lines[i] + "\n"
	}

	text += st.Options(labels, index, box_width-2) + "\n"

	for i, line := range st.Lines(text) {
		xshift := (box_width - st.RuneLen(line)) / 2
		yshift := i
		tbprint(midx+xshift, midy+yshift, coldef, coldef, line)
	}
	termbox.Flush()
}

func (self Game) YesNo(prompt string) bool {
	results := [2]bool{false, true}
	labels := []string{"no", "yes"}
	index := 0
	self.DrawPrompt(prompt, labels, index%2)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				index = nt.AbsInt(index - 1)
			case termbox.KeyArrowRight:
				index = nt.AbsInt(index + 1)
			case termbox.KeyEnter, termbox.KeySpace:
				break loop
			case termbox.KeyCtrlC:
				self.Clear()
				os.Exit(0)
			}
			switch ev.Ch {
			case 'a':
				index = nt.AbsInt(index - 1)
			case 'd':
				index = nt.AbsInt(index + 1)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		self.DrawPrompt(prompt, labels, index%2)
	}

	self.Clear()
	self.Draw()
	return results[index%2]
}

func (self Game) Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func (self Game) Exit() {
	self.clockStop = true
	self.Clear()
	os.Exit(0)
}

func (self *Game) Loop() {
	self.Draw()
	// go self.clock()
	for {
		if !self.Solved() {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyEsc, termbox.KeyCtrlC:
					self.Save()
					// self.Clear()
					self.Exit()
					return
				case termbox.KeyArrowLeft:
					self.ActionLeft()
				case termbox.KeyArrowRight:
					self.ActionRight()
				case termbox.KeyArrowUp:
					self.ActionUp()
				case termbox.KeyArrowDown:
					self.ActionDown()
				case termbox.KeySpace:
					self.ToggleSlideFlip()
				case termbox.KeyCtrlZ:
					self.Undo()
				case termbox.KeyCtrlY:
					self.Redo()
				case termbox.KeyCtrlS:
					self.Save()
				case termbox.KeyCtrlR:
					self.Reset()
				case termbox.KeyTab:
					self.NextPanel()
				}
				switch ev.Ch {
				case 'w':
					self.MoveUp()
				case 'a':
					self.MoveLeft()
				case 's':
					self.MoveDown()
				case 'd':
					self.MoveRight()
				case 'n':
					self.clockStop = true
					self.New()
				}
				if self.Solved() {
					self.clockStop = true
					self.Draw()
					self.Congratulate()
					self.Clear()
					if self.YesNo("Retry?") {
						self.Reset()
						// self.board.Tiles = it.Swap(it.Swap(nt.Range(1, 10, 1), 0, 1), 0, 2)
					} else {
						self.Clear()
						if self.YesNo("Quit?") {
							self.Exit()
							return
						} else {
							self.New()
						}
					}
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		} else {
			self.clockStop = true
		}
		self.Draw()
	}
}

func (self Game) Draw() {
	width, height := self.Dim()
	width *= self.tileWidth
	height *= self.tileHeight

	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	w, h := termbox.Size()

	midy := h / 2
	midx := (w - width) / 2

	mode := "mode:\t%s"
	if self.slide {
		mode = fmt.Sprintf(mode, "slide")
	} else {
		mode = fmt.Sprintf(mode, "flip")
	}
	mode = st.ResolveTabsRight(mode)
	tbprint(w-len(mode), h-1, coldef, coldef, mode)

	moves := fmt.Sprintf("moves:\t%v", self.Count())
	moves = st.ResolveTabsRight(moves)
	tbprint(w-len(moves), 0, coldef, coldef, moves)

	// dt :=
	// tbprint(0, 0, coldef, coldef, self.elapsedTime.String())

	for i, line := range st.Lines(self.repr()) {
		xshift := (width - st.RuneLen(line)) / 2
		yshift := i
		tbprint(midx+xshift, midy+yshift, coldef, coldef, line)
	}
	termbox.Flush()
}

func (self Game) repr() string {
	charset := self.borderChars()
	var str string
	str += self.top(charset)
	str += self.lines(charset)
	str += self.bottom(charset)
	return str
}

func (self Game) lines(charset map[string]rune) string {
	var repr string
	xpos, ypos := self.board.Current()

	for y, row := range self.board.Grid() {
		line := string(charset["pipe"])
		for x := range row {
			var str string
			if y == ypos && x == xpos {
				str = fmt.Sprintf("{%v}", row[x])
			} else {
				str = fmt.Sprintf("%v", row[x])
			}
			line += st.Center(str, self.tileWidth) + string(charset["pipe"])
		}
		repr += line + "\n"
		if y < self.board.Height()-1 {
			repr += self.midline(charset)
		}
	}
	return repr
}

func (self Game) midline(charset map[string]rune) string {
	var line string

	line += string(charset["middle_left"])
	for i := 0; i < self.board.Width(); i++ {
		for j := 0; j <= self.tileWidth; j++ {
			line += string(charset["dash"])
		}
		if i < self.board.Width()-1 {
			line += string(charset["middle_corner"])
		} else {
			line += string(charset["middle_right"])
		}
	}
	line += "\n"
	return line
}

func (self Game) top(charset map[string]rune) string {
	var line string

	line += string(charset["top_left_corner"])
	for i := 0; i < self.board.Width(); i++ {
		for j := 0; j <= self.tileWidth; j++ {
			line += string(charset["dash"])
		}
		if i < self.board.Width()-1 {
			line += string(charset["middle_top"])
		} else {
			line += string(charset["top_right_corner"])
		}
	}
	line += "\n"
	return line
}

func (self Game) bottom(charset map[string]rune) string {
	var line string

	line += string(charset["bottom_left_corner"])
	for i := 0; i < self.board.Width(); i++ {
		for j := 0; j <= self.tileWidth; j++ {
			line += string(charset["dash"])
		}
		if i < self.board.Width()-1 {
			line += string(charset["middle_bottom"])
		} else {
			line += string(charset["bottom_right_corner"])
		}
	}
	return line
}

func (self Game) borderChars() map[string]rune {
	charset := map[string]rune{}
	// unicode delimiters for board and tile
	if runewidth.EastAsianWidth {
		charset["pipe"] = '|'
		charset["dash"] = '-'
		charset["top_left_corner"] = '+'
		charset["top_right_corner"] = '+'
		charset["bottom_left_corner"] = '+'
		charset["bottom_right_corner"] = '+'
		charset["middle_corner"] = '+'
		charset["middle_top"] = '+'
		charset["middle_bottom"] = '+'
		charset["middle_left"] = '+'
		charset["middle_right"] = '+'
	} else {
		charset["pipe"] = '║'
		charset["dash"] = '═'
		charset["top_left_corner"] = '╔'
		charset["top_right_corner"] = '╗'
		charset["bottom_left_corner"] = '╚'
		charset["bottom_right_corner"] = '╝'
		charset["middle_corner"] = '╬'
		charset["middle_top"] = '╦'
		charset["middle_bottom"] = '╩'
		charset["middle_left"] = '╠'
		charset["middle_right"] = '╣'
	}
	return charset
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

var celebrations = []string{
	"Congratulations! You have officialty didded it!",
	"O, u think u nice?",
	"Nawwwwsty!",
	"Get that fothermucker!",
	"So you think you're decent?",
	"Gangsta over here...",
	"Okay, but can you clutch this?",
	"I may or may not have seen better",
	"You really 'Bout that life, huh?",
	"Who died and made you president!?",
	"You went and put our whole operation under water...",
	"Alright, see me with the hands!",
	"This means WAR!",
	"Sticks and stones can't break my bones but you just hurt me :[",
	"Vou te partir a cara!",
	"That was smooth!",
	"Good looks!",
	"Lood Gooks!",
	"Dutty!",
	"ZOMG!!! We have a champion!",
	"Give yourself a pat on the back, that was piff!",
	"L33T!!",
	"Y33T!!",
	"\"WOO HOO!!\"\nDamon Albarn",
}

func NewGame(width, height int) Game {
	board := board.NewBoard(width, height)
	// return Game{&board, true, []string{}, []string{}, 0, 6, 3}
	g := Game{&board, true, []string{}, []string{}, 0, 6, 3, time.Now(), time.Second * 0, false}
	g.Reset()
	return g
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
	// g.board.Tiles = it.Swap(it.Swap(nt.Range(1, 10, 1), 0, 1), 0, 2)
	g.Loop()
}
