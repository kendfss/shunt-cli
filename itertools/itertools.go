package itertools

import (
	"fmt"
	"math"

	nt "github.com/kendfss/shunt-cli/numtools"
)

func Clone(slice []int) []int {
	clone := make([]int, len(slice))
	for i, elem := range slice {
		clone[i] = elem
	}
	return clone
}

func Equal(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func Cast(slice []int, width, height int) [][]int {
	length := width * height
	if len(slice) != length {
		panic(fmt.Sprintf("Length(%v) does not agree with dimensions(%v, %v)", length, width, height))
	}
	grid := make([][]int, height)
	var i int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[y] = append(grid[y], slice[i])
			i += 1
		}
	}
	return grid
}

func CastWide(slice []int, width int) [][]int {
	height := len(slice) / width
	if len(slice) != width*height {
		panic(fmt.Sprintf("Length(%v) is not a multiple of width(%v)", len(slice), width))
	}
	return Cast(slice, width, height)
}

func CastLong(slice []int, height int) [][]int {
	width := len(slice) / height
	if len(slice) != width*height {
		panic(fmt.Sprintf("Length(%v) is not a multiple of height(%v)", len(slice), height))
	}
	return Cast(slice, width, height)
}

func Max(slice []int) int {
	var max int
	for _, elem := range slice {
		if elem > max {
			max = elem
		}
	}
	return max
}

func Freq(elem int, slice []int) int {
	var ctr int
	for i := range slice {
		if elem == slice[i] {
			ctr++
		}
	}
	return ctr
}

func Freqs(slice []int) map[int]int {
	table := map[int]int{}
	for _, elem := range slice {
		table[elem] = Freq(elem, slice)
	}
	return table
}

func Expectation(slice []int) float64 {
	var sum float64
	for _, elem := range slice {
		sum += float64(elem) / float64(Freq(elem, slice))
	}
	return sum
}

func Sum(slice []float64) float64 {
	var sum float64
	for _, elem := range slice {
		sum += elem
	}
	return sum
}

func Floats(slice []int) []float64 {
	out := make([]float64, len(slice))
	for i, elem := range slice {
		out[i] = float64(elem)
	}
	return out
}

func Mean(slice []float64) float64 {
	return Sum(slice) / float64(len(slice))
}

func StDev(slice []float64) float64 {
	devs := make([]float64, len(slice))
	mean := Mean(slice)
	for i, elem := range slice {
		devs[i] = elem - mean
	}
	return Mean(devs)
}

func Variance(slice []float64) float64 {
	return math.Pow(StDev(slice), 2)
}

func Shuffle(slice []int) []int {
	indices := nt.Range(0, len(slice), 1)
	for len(indices) > 1 {
		pair := nt.IndexPair(len(indices))
		indices = append(indices[:pair[0]], indices[pair[0]+1:]...)
		shift := -1
		if !(pair[1] > 0) {
			shift = 0
		}
		indices = append(indices[:pair[1]+shift], indices[pair[1]+shift:]...)
		slice = Swap(slice, pair[0], pair[1])
	}
	return slice
}

func Swap(slice []int, m, n int) []int {
	slice[m], slice[n] = slice[n], slice[m]
	return slice
}

func HasInt(rack []int, i int) bool {
	for _, e := range rack {
		if e == i {
			return true
		}
	}
	return false
}

func FlatSample(rack []int, size int) []int {
	if size > len(rack) {
		panic("Sample size exceeds slice length")
	}
	sample := []int{}
	sampled := []int{}
	for len(sample) < size {
		i := nt.Randex(len(rack))
		if !HasInt(sampled, i) {
			sampled = append(sampled, i)
			sample = append(sample, rack[i])
		}
	}
	return sample
}

func Sample(rack []int, size int) []int {
	sample := []int{}
	for len(sample) < size {
		i := nt.Randex(len(rack))
		sample = append(sample, rack[i])
	}
	return sample
}

func ReverseList(a []int) []int {
	rack := Clone(a)
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		rack[l], rack[r] = rack[r], rack[l]
	}
	return rack
}
