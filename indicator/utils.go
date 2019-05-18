package indicator

import (
	"math"
	"fmt"
	"strconv"
)

func sum(x []float64, start uint64, end uint64) (s float64) {
	if start > end {
		start, end = end, start
	}
	for i := start; i < end && i < uint64(len(x)); i++ {
		s += x[i]
	}
	return
}

func avg(x []float64, start uint64, end uint64) (a float64) {
	if start > end {
		start, end = end, start
	}
	a = sum(x, start, end)/float64(end-start)
	return
}

func variance(x []float64, start uint64, end uint64) (v float64) {
	if start > end {
		start, end = end, start
	}
	a := avg(x, start, end)
	s := 0.0
	for i := start; i < end && i < uint64(len(x)); i++ {
		s += math.Pow(x[i]-a, 2)
	}
	v = s/float64(end-start)
	return
}

func standardDeviation(x []float64, start uint64, end uint64) (s float64) {
	s = math.Pow(variance(x, start, end), 0.5)
	return
}

func toFixed(x float64, decimal uint64) (y float64) {
	_x := fmt.Sprintf("%."+strconv.Itoa(int(decimal))+"f", x)
	y, _ = strconv.ParseFloat(_x, 64)
	return
}
