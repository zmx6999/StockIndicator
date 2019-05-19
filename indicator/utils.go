package indicator

import (
	"math"
	"strconv"
	"fmt"
)

func Sum(x []float64, start uint64, end uint64) (s float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		s = math.NaN()
		return
	}

	if start == end {
		s = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	for i := start; i < end; i++ {
		s += x[i]
	}
	return
}

func Avg(x []float64, start uint64, end uint64) (a float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		a = math.NaN()
		return
	}

	if start == end {
		a = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	a = Sum(x, start, end)/float64(end-start)
	return
}

func Variance(x []float64, start uint64, end uint64) (v float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		v = math.NaN()
		return
	}

	if start == end {
		v = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	a := Avg(x, start, end)
	s := 0.0
	for i := start; i < end; i++ {
		s += math.Pow(x[i]-a, 2)
	}
	v = s/float64(end-start)
	return
}

func StandardDeviation(x []float64, start uint64, end uint64) (s float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		s = math.NaN()
		return
	}

	if start == end {
		s = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	s = math.Pow(Variance(x, start, end), 0.5)
	return
}

func ToFixed(x float64, decimal uint64) (y float64) {
	y, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(decimal))+"f", x), 64)
	return
}
