package stringtools // import "tildegit.org/eli2and40/rube/stringtools"

import (
	"strings"
	"unicode/utf8"
	"fmt"
	
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	
	nt"rube/numtools"
)

const tabstop_length = 8

func Str(value interface{}) string {
	return fmt.Sprint(value)
}

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

func ResolveTabsRight(str string) string {
	out := ""
	words := strings.Split(str, "\t")
	for i:=len(words)-1; 0<=i; i-- {
		word := words[i]
		out = word + out
		// if i != len(words) - 1 {
		if i != 0 {
			for RuneLen(out) % tabstop_length != 0 {
				out = " " + out
			}
		}
	}
	return out
}
func ResolveTabsLeft(str string) string {
	out := ""
	// words := strings.Split(str, "\t")
	// for i, word := range words {
		// out += word
		// if i != len(words) - 1 {
		// 	for RuneLen(out) % tabstop_length != 0 {
		// 		out += " "
		// 	}
		// }
	for i, c := range str {
		if c == rune('\t') {
			dx := (RuneLen(out) + i) % tabstop_length
			out += BlankString(dx)
			continue
		}
		out += string(c)
	}
	return out
}












func Fill(x, y, w, h int, cell termbox.Cell) {
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

func DeRune(r rune) string {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	return Read(byte_slice_insert([]byte{}, 0, buf[:n])...)
}



func CenterWrap(prompt string, width int) string {
	out := ""
	prompt = ResolveTabsLeft(prompt)
	lines := []string{}
	for len(prompt)>0 {
		for i:=0; i<width && len(prompt)>0; i++ {
			out += prompt[:1]
			prompt = prompt[1:]
		}
		// out += "\n"
		if len(out) > 0 {
			lines = append(lines, out)
			out = ""
		}
	}
	for i := range lines {
		out += Center(lines[i], width) + "\n"
	}
	return out
}

func Options(labels []string, index, width int) string {
	out := clone(labels)
	if delta:=(len(out)-index); delta < 1 {
		panic(fmt.Sprintf("Not enough labels(%v) for index(%v). Add %v labels or subtract %v from the index", len(out), index, delta, delta))
	}
	for i := range out {
		if i==index {
			out[i] = fmt.Sprintf("{%s}", out[i])
			break
		}
	}
	result := strings.Join(out, "    ")
	return CenterWrap(result, width-len(result))
}

func clone(slice []string) []string {
	out := make([]string, len(slice))
	for i := range out {
		out[i] = slice[i]
	}
	return out
}
