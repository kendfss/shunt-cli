package board_test

import (
    "testing"
    
    
    // nt"tildegit.org/eli2and40/rube/numtools"
    nt"rube/numtools"
    // ."tildegit.org/eli2and40/rube/board"
    ."rube/board"
)


func TestNewBoard(t *testing.T) {
	width := nt.RandInt(2, 20)
	height := nt.RandInt(2, 20)
    b := NewBoard(width, height)
    if !(b.Height() == height) {
    	t.Errorf("NewBoard.Height():\n\texpected\t%v\n\tbut got\t%v", height, b.Height())
    }
}


func TestMoveCoordsIndexUniformBoard(t *testing.T) {
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
						t.Errorf("Board.Move(%v, %v):\n\tMiddle index(%v) != Actual index(%v)", x, y, b.Index(cx1, cy1), i1)
					}
					
					b.Move(-x, -y)
					b.Move(-x, -y)
					cx2, cy2 := b.Current()
					i2 := b.current
					
					if cx0 != cx2 {
						t.Errorf("Board.Move(%v, %v):\n\tFinal x-coordinate(%v) != Starting x-coordinate(%v)", x, y, cx0, cx2)
					}
					if cy0 != cy2 {
						t.Errorf("Board.Move(%v, %v):\n\tFinal y-coordinate(%v) != Starting y-coordinate(%v)", x, y, cy0, cy2)
					}
					if i0 != i2 {
						t.Errorf("Board.Move(%v, %v):\n\tFinal index(%v) != Starting index(%v)", x, y, i0, i2)
					}
				}
			}
		}
	}
}

func TestFlipSlide(t *testing.T) {
	for height:=1; height<20; height++ {
		for width:=1; width<20; width++ {
			b := NewBoard(width, height)
			for y:=0; y<height; y++ {
				for x:=0; x<width; x++ {
					b.Move(x, y)
					
					var i int
					for ; i < b.Width(); i++ {
						b.SlideLeft()
					}
					if !b.Solved() {
						t.Errorf("\n%s.SlideLeft(%v) does not return to solution", b, i)
					}
					
					i = 0
					for ; i < b.Width(); i++ {
						b.SlideRight()
					}
					if !b.Solved() {
						t.Errorf("\n%s.SlideRight(%v) does not return to solution", b, i)
					}
					
					i = 0
					for ; i < b.Height(); i++ {
						b.SlideUp()
					}
					if !b.Solved() {
						t.Errorf("\n%s.SlideUp(%v) does not return to solution", b, i)
					}
					
					i = 0
					for ; i < b.Height(); i++ {
						b.SlideDown()
					}
					if !b.Solved() {
						t.Errorf("\n%s.SlideDown(%v) does not return to solution", b, i)
					}
					
					b.SlideDown()
					b.SlideUp()
					if !b.Solved() {
						c := b.Solution()
						c.SlideDown()
						t.Errorf("\n%s    ...    %s    ...    %s\n\tSlideUp does not undo SlideDown", b.Solution(), c, b)
					}
					
					b.SlideUp()
					b.SlideDown()
					if !b.Solved() {
						c := b.Solution()
						c.SlideUp()
						t.Errorf("\n%s    ...    %s    ...    %s\n\tSlideUp does not undo SlideDown", b.Solution(), c, b)
					}
					
					b.SlideLeft()
					b.SlideRight()
					if !b.Solved() {
						c := b.Solution()
						c.SlideLeft()
						t.Errorf("\n%s    ...    %s    ...    %s\n\tSlideRight does not undo SlideLeft", b.Solution(), c, b)
					}
					
					b.SlideRight()
					b.SlideLeft()
					if !b.Solved() {
						c := b.Solution()
						c.SlideRight()
						t.Errorf("\n%s    ...    %s    ...    %s\n\tSlideLeft does not undo SlideRight", b.Solution(), c, b)
					}
					
					
					i = 0
					for ; i < 2; i++ {
						b.FlipHorizontal()
					}
					if !b.Solved() {
						t.Errorf("\n%s.FlipHorizontal(%v) does not return to solution", b, i)
					}
					
					i = 0
					for ; i < 2; i++ {
						b.FlipVertical()
					}
					if !b.Solved() {
						t.Errorf("\n%s.FlipVertical(%v) does not return to solution", b, i)
					}
					b.Move(-x, -y)
				}
			}
		}
	}
}
