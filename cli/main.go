package main

import (
	"fmt"
	
	"github.com/nsf/termbox-go"
	
	rube"tildegit.org/eli2and40/rube/cli/game"
)

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
	
	g := rube.NewGame(3, 3)
	g.Loop()
	
}
