package board

import (
	"fmt"
	
	it"tildegit.org/eli2and40/rube/cli/itertools"
	nt"tildegit.org/eli2and40/rube/cli/numtools"
	st"tildegit.org/eli2and40/rube/cli/stringtools"
)

type Board struct {
	Tiles []int
	stride, current int
	original_state []int
}
func (self *Board) newTiles(tiles []int) {
	self.Tiles = tiles
}
func (self Board) Opening() []int {
	return it.Clone(self.original_state)
}
func (self *Board) Reset() {
	// (*self).Tiles = it.Clone(self.original_state)
	// self.Tiles = it.Clone(self.original_state)
	self.Tiles = self.Opening()
	self.MoveTo(0, 0)
}
func (self Board) Height() int {
	return len(self.Tiles) / self.Width()
}
func (self Board) Width() int {
	return self.stride
}
func (self Board) String() string {
	repr := ""
	for i, elem := range self.Tiles {
		var str string
		if i==self.current {
			str = "|" + st.Center(fmt.Sprintf("{%v}", elem), 7)
		} else {
			str = "|" + st.Center(fmt.Sprintf("%v", elem), 7)
		}
		
		repr += str
		
		if (i+1) % self.Width() == 0 {
			repr += "|"
			if i < len(self.Tiles) - 1 {
				repr += "\n"
			}
		}
	}
	return repr
}
func (self Board) Solution() Board {
	new := Board{it.Clone(self.Tiles), self.stride, self.current, it.Clone(self.original_state)}
	
	for !new.Solved() {
		for i, elem := range new.Tiles[1:] {
			i++
			if elem < new.Tiles[i-1] {
				new.Swap(i, i-1)
			}
		}
	}
	return new
}
func (self *Board) Swap(m, n int)  {
	self.Tiles[m], self.Tiles[n] = self.Tiles[n], self.Tiles[m]
}
func (self Board) Solved() bool {
	for i := 0; i < len(self.Tiles)-1; i++ {
		elem := self.Tiles[i]
		if (elem > self.Tiles[i+1]) {
			return false
		}
	}
	return true
}
func (self *Board) Shuffle() {
	indices := nt.Range(0, len(self.Tiles), 1)
	for len(indices) > 1 {
		pair := nt.IndexPair(len(indices))
		indices = append(indices[:pair[0]], indices[pair[0]+1:]...)
		shift := -1
		if !(pair[1] > 0) {
			shift = 0
		}
		indices = append(indices[:pair[1]+shift], indices[pair[1]+shift:]...)
		self.Swap(pair[0], pair[1])
	}
}
func (self Board) Grid() [][]int {
 	return it.CastWide(self.Tiles, self.Width())
}
func (self Board) Equals(other Board) bool {
	return self.Width()==other.Width() && it.Equal(self.Tiles, other.Tiles)
}
func (self *Board) Current() (int, int) {
	x := self.current % self.Width()
	y := self.current / self.Width()
	return x, y
}
func (self *Board) Coords(index int) (int, int) {
	// if index > len(self.Tiles) {
	// 	// panic(fmt.Sprintf("Index(%v) is greater than the number of tiles(%v)", index, len(self.Tiles)))
	// 	index %= len(self.Tiles)
	// }
	for index < 0 {
		index += len(self.Tiles)
	}
	
	index %= len(self.Tiles)
	x := index % self.Width()
	y := index / self.Width()
	return x, y
}
func (self Board) Index(x, y int) int {
	var i int
	
	sx := nt.Sign(x)
	sy := nt.Sign(y)
	
	x += self.Width()
	y += self.Height()
	x %= self.Width()
	y %= self.Height()
	
	for ;nt.AbsInt(y) > 0; y -= sy {
		i += self.Width()
	}
	for ;nt.AbsInt(x) > 0; x -= sx {
		i++
	}
	return i
}
func (self *Board) Move(x, y int) {
	for x < 0 {
		x += self.Width()
		// x += self.Height()
	}
	for y < 0 {
		y += self.Height()
	}
	x %= self.Width()
	// x %= self.Height()
	y %= self.Height()
	sx := nt.Sign(x)
	sy := nt.Sign(y)
	dx := 0
	// for ;nt.AbsInt(y) > 0; y -= sy {
	for ;nt.AbsInt(y) > 0; y -= sy {
		self.current += sy * self.Width()
		// dy += sy * self.Width()
	}
	for ;nt.AbsInt(x) > 0; x -= sx {
		// self.current += sx 
		dx += sx + self.Height()
		dx %= self.Width()
		// if dx == self.Width() {
		// 	dx = 0
		// }
		// dx += sx * len(self.Tiles)
	}
	self.current %= len(self.Tiles)
	dx %= self.Width()
	self.current += dx
	self.current %= len(self.Tiles)
}
func (self *Board) MoveTo(x, y int) {
	self.current = self.Index(x, y)
}
func (self *Board) MoveUp() {	
	x, y := self.Current()
	y -= 1
	y += self.Height()
	self.MoveTo(x, y)
}
func (self *Board) MoveDown() {	
	x, y := self.Current()
	y += 1
	y %= self.Height()
	self.MoveTo(x, y)
}
func (self *Board) MoveLeft() {
	x, y := self.Current()
	x -= 1
	x += self.Width()
	self.MoveTo(x, y)
}
func (self *Board) MoveRight() {
	x, y := self.Current()
	x += 1
	x %= self.Width()
	self.MoveTo(x, y)
}
func (self *Board) flip(vertical bool) {
	X, Y := self.Current()
	if vertical {
		// for y:=0; y<self.Height()/2; y++ {
		// 	for x:=0; x<self.Height()/2; x++ {
				
		// 	}
		// }                                      
		for y:=0; y<self.Height()/2; y++ {
			i0 := self.Index(X, y)
			i1 := self.Index(X, self.Height()-y-1)
			self.Swap(i0, i1)
		}
	} else {
		for x:=0; x<self.Width()/2; x++ {
			i0 := self.Index(x, Y)
			i1 := self.Index(self.Width()-x-1, Y)
			self.Swap(i0, i1)
		}
	}
}
func (self *Board) FlipVertical() {
	self.flip(true)
}
func (self *Board) FlipHorizontal() {
	self.flip(false)
}
func (self *Board) slide(vertical bool, direction int) {
	if !(direction==1) && !(direction==-1) {
		panic(fmt.Sprintf("Cannot interpret direction(%v). Choose 1(down/right) or -1(up/left)"))
	}
	X, Y := self.Current()
	tiles := it.Clone(self.Tiles)
	if vertical {
		for y:=0; y<self.Height(); y++ {
			i0 := self.Index(X, y)
			i1 := self.Index(X, (y+1)%self.Height())
			if direction==-1 {
				i0, i1 = i1, i0
			}
			self.Tiles[i1] = tiles[i0]
		}
	} else {
		for x:=0; x<self.Width(); x++ {
			i0 := self.Index(x, Y)
			i1 := self.Index((x+1)%self.Width(), Y)
			if direction==-1 {
				i0, i1 = i1, i0
			}
			self.Tiles[i1] = tiles[i0]
		}
	}
}
func (self *Board) SlideUp() {
	self.slide(true, -1)
}
func (self *Board) SlideDown() {
	self.slide(true, 1)
}
func (self *Board) SlideLeft() {
	self.slide(false, -1)
}
func (self *Board) SlideRight() {
	self.slide(false, 1)
}













func NewBoard(width, height int) Board {
	length := width * height
	tiles := make([]int, length)
	for i := 0; i < length; i++ {
		tiles[i] = 1 + i
	}
	// return Board{tiles, width, 0, tiles}
	tiles = it.Shuffle(tiles)
	return Board{tiles, width, 0, tiles}
}
func NewUniformBoard(width, height, min, max int) Board {
	length := width * height
	if max - min < length {
		panic(fmt.Sprintf("Cannot have max-min(%v) less than length(%v)", max-min, length))
	}
	tiles := make([]int, length)
	tiles[0] = nt.RandInt(min, max)
	for i := 1; i < length; i++ {
		for len(it.Freqs(tiles[:i+1])) < i+1 {
			tiles[i] = nt.RandInt(min, max)
		}
	}
	return Board{tiles, width, 0, tiles}
}
func NewRandomBoard(width, height, min, max int) Board {
	length := width * height
	tiles := make([]int, length)
	
	for i := 0; i < length; i++ {
		tiles[i] = nt.RandInt(min, max)
	}
	return Board{tiles, width, 0, tiles}
}




func TestMoveCoordsIndexUniformBoard() {
	for height:=2; height<10; height++ {
		for width:=2; width<10; width++ {
			b := NewUniformBoard(width, height, 0, height*width)
			for y:=0; y<(b.Height()*b.Height()); y++ {
				for x:=0; x<(b.Width()*b.Width()); x++ {
					cx0, cy0 := b.Current()
					i0 := b.current
					
					
					b.Move(x, y)
					b.Move(x, y)
					cx1, cy1 := b.Current()
					i1 := b.current
					
					if b.Index(cx1, cy1) != i1 {
						panic(fmt.Sprintf("Board.Move(%v, %v):\n\tMiddle index(%v) != Actual index(%v)", x, y, b.Index(cx1, cy1), i1))
					}
					
					b.Move(-x, -y)
					b.Move(-x, -y)
					cx2, cy2 := b.Current()
					i2 := b.current
					
					if cx0 != cx2 {
						panic(fmt.Sprintf("Board.Move(%v, %v):\n\tFinal x-coordinate(%v) != Starting x-coordinate(%v)", x, y, cx0, cx2))
					}
					if cy0 != cy2 {
						panic(fmt.Sprintf("Board.Move(%v, %v):\n\tFinal y-coordinate(%v) != Starting y-coordinate(%v)", x, y, cy0, cy2))
					}
					if i0 != i2 {
						panic(fmt.Sprintf("Board.Move(%v, %v):\n\tFinal index(%v) != Starting index(%v)", x, y, i0, i2))
					}
				}
			}
		}
	}
	fmt.Println("made it")
}






func main() {
	TestMoveCoordsIndexUniformBoard()
}

