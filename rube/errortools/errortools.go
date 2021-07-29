// package errortools // import "tildegit.org/eli2and40/rube/errortools"
package errortools // rube/errortools


import (
	"log"
	"fmt"
)


func Assertf(err error, fstr string) {
    if err != nil {
        fstr += "\n\t%v"
        str := fmt.Sprintf(fstr, err) + "\n"
        panic(str)
    }
}
func Checkf(err error, fstr string) {
    if err != nil {
        str := fmt.Sprintf(fstr, err) + "\n"
        log.Fatalf(str)
    }
}

func Assert(err error, str string) {
    if err != nil {
        str += "\n\t%v"
        str := fmt.Sprintf(str, err) + "\n"
        panic(str)
    }
}
func Check(err error, str string) {
    if err != nil {
        str += "\n\t%v"
        str := fmt.Sprintf(str, err) + "\n"
        log.Fatalf(str)
    }
}

func Bool(err error) bool {
	if err != nil {
		return true
	}
	return false
}
