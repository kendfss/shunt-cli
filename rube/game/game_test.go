package game_test

import (
	"testing"
)

func MechanicsTest() {
	maxw, maxl := 10, 10
	
	for height:=1; height<maxl; height++ {
		for width:=1; width<maxw; width++ {
			g := NewGame(width, height)
			for y:=0; y<height; y++ {
				for x:=0; x<width; x++ {
					g.GoTo(x, y)
					g = g.Solution()
					
					
					g.ActionUp()
					g.ToggleSlideFlip()
					g.ActionDown()
					g.ToggleSlideFlip()
					g.ActionLeft()
					g.ToggleSlideFlip()
					g.ActionRight()
					if g.Count() != 4 {
						panic(t.Errorf("Counting error!(action)\n\thave\t%v\n\twant\t%v\n%s\n", g.Count(), 4, g.board))
					}
					g.Undo()
					g.Undo()
					g.Undo()
					g.Undo()
					if g.Count() != 8 {
						panic(t.Errorf("Counting error!(undoing)\n\thave\t%v\n\twant\t%v\n%s\n", g.Count(), 8, g.board))
					}
					if len(g.unmoves) != 4 {
						panic(t.Errorf("Counting error!(unmoves)\n\thave\t%v\n\twant\t%v\n%s\n", len(g.unmoves), 4, g.board))
					}
					if !g.Solved() {
						panic(t.Errorf("Undo error!\n%s\n", g.board))
					}
					g.Redo()
					g.Redo()
					g.Redo()
					g.Redo()
					if g.Count() != 12 {
						panic(t.Errorf("Counting error!(redoing)\n\thave\t%v\n\twant\t%v\n%s\n", g.Count(), 12, g.board))
					}
					if len(g.unmoves) != 0 {
						panic(t.Errorf("Counting error!(unmoves)\n\thave\t%v\n\twant\t%v\n%s\n", len(g.unmoves), 0, g.board))
					}
					g.Reset()
				}
			}
		}
	}
}
