package numtools

import (
	"crypto/rand"
	"math"

	et "github.com/kendfss/shunt-cli/errortools"
)

func Sign(i int) int {
	if i >= 0 {
		return 1
	}
	return -1
}

func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}

func PowInt(base, height int) int {
	return int(math.Pow(float64(base), float64(height)))
}

func Randex(length int) int {
	bites := make([]byte, 1)
	_, err := rand.Read(bites)
	et.Check(err, "Couldn't generate random index")
	return int(bites[0]) % length
}

func OddInts(cutoff int) []int {
	var out []int
	for _, elem := range Range(0, cutoff, 1) {
		if elem%2 != 0 {
			out = append(out, elem)
		}
	}
	return out
}

func EvenInts(cutoff int) []int {
	var out []int
	for _, elem := range Range(0, cutoff, 1) {
		if elem%2 == 0 {
			out = append(out, elem)
		}
	}
	return out
}

func RandInt(min, max int) int {
	return Randex(max-min+1) + min
}

func IndexPair(max int) [2]int {
	pair := [2]int{}
	for pair[0] == pair[1] {
		pair[0] = Randex(max)
		pair[1] = Randex(max)
	}
	return pair
}

func Conform(value, cutoff int) int {
	value %= cutoff
	if value < 1 {
		value += cutoff
	}
	return value
}

func BitFlip(n int) int {
	var out int
	if n == 0 || n == 1 {
		out = n ^ 1
	} else {
		panic("Cannot Flip non-Bit")
	}
	return out
}

func Range(start, stop, step int) []int {
	var rack []int
	for ; start < stop; start += step {
		rack = append(rack, start)
	}
	return rack
}
