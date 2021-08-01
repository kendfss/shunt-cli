// package clocktools // import "rube/clocktools"
package main

import (
	"time"
	"fmt"
	"strconv"
	
	nt"rube/numtools"
	it"rube/itertools"
)

const (
	Second time.Duration = time.Second
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = Hour * 24
	Week                 = Day * 7
	Feb                  = 4 * Week
	Year                 = 52 * Week
)

func TimeDict(duration time.Duration) map[string]time.Duration {
    periods := map[string]time.Duration {
        "Years": duration / Year,
        "Febs": (duration % Year) / Feb,
        "Weeks": (duration % Feb) / Week,
        "Days": (duration % Week) / Day,
        "Hours": (duration % Day) / Hour,
        "Minutes": (duration % Hour) / Minute,
        "Seconds": (duration % Minute) / Second,
    }
    return periods
}


func band(level, base, shift, root int) (int, int) {
	rule := func(x int) int { return root + x + shift }
	return rule((level + 1) * base),  rule(level * base)
	
}
func format(selection []int) int {
	out := ""
	for _, e := range selection {
		// out += string(e)
		out += strconv.FormatUint(uint64(e), 10)
	}
	i, err := strconv.ParseInt(out, 0, 0)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse %q:\n\t%s", out, err))
	}
	return int(i)
}
func sizes(base []int, x, y int) []int {
	value := it.Sample(base, y)
	rack := []int{}
	for i:=x; i<y; i++ {
		// rack = append(rack, format(value[:i][:len(value)-1:-1]))
		rack = append(rack, format(it.ReverseList(value[:i])))
	}
	return rack
}
// func watch
func main() {
	digits := nt.Range(0, 10, 1)
	
	fmt.Println(digits)
    fmt.Println(format(digits[1:]))
	l := 0
    s := 0
    b := 30
    r := 1
    x, y := band(l, b, s, r)
    // fmt.Println(x,y)
    // val := it.Sample(digits, y)
    // sizes := [format(val[:i][:len(val)-1:-1]) for i in nt.Range(x, y, 1)]
    // szs := sizes(value, x, y)
    // for _, e := range szs {
    // 	fmt.Println(e)
    // }
    // show(sizes)
    // sizes = [44259260028,315436]
    
    // for size in sizes[::-1]:
    //     string := f"""{(len(lasso(max(sizes)))-len(lasso(size)))*' '+lasso(size)}
    //         {size = }
    //         {len(str(size)) = }
    //         {nice_size(size) = }
    //         {magnitude(size) = }
        
            
    //     """.splitlines()
    //     fmt.Println(*((x.strip(), x)[i<1] for i, x in enumerate(string)), sep='\n\t', end='\n\n')
    // fmt.Println(lasso(sizes[-1]))
    fmt.Println(band(l,s,b,r))
    fmt.Println(x, y)
    // fmt.Println(f'{format(digits):,}')
    fmt.Println(format(digits[1:]))
}
