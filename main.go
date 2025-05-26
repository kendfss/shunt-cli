package main

import (
	"fmt"

	"github.com/nsf/termbox-go"

	rube "github.com/kendfss/shunt-cli/game"
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
