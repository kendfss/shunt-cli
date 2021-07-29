package stringtools // import "tildegit.org/eli2and40/rube/stringtools"

import (
	"strings"
	
	"github.com/mattn/go-runewidth"
	
	nt"rube/numtools"
)

func BlankString(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += " "
	}
	return str
}
func Center(str string, width int) string {
	dx := nt.AbsInt(len(str) - width)
	var blank string
	if (dx % 2) == 0 {
		blank = BlankString(dx / 2)
	} else {
		blank = BlankString((dx + 1) / 2)
	}
	return blank + str + blank
}

func RuneLen(msg string) int {
	val := 0
	for i := range msg {
		val += runewidth.RuneWidth(rune(msg[i]))
	}
	return val
}
func Read(msg...byte) string {
	out := ""
	for i := range msg {
		out += string(msg[i])
	}
	return out
}
func Lines(str string) []string {
	return strings.Split(str, "\n")
}


