package indicator

import "math"

func Sum(x []float64, start uint64, end uint64) (y float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		y = math.NaN()
		return
	}

	if start == end {
		y = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	for i := start; i < end; i++ {
		y += x[i]
	}
	return
}

func Avg(x []float64, start uint64, end uint64) (y float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		y = math.NaN()
		return
	}

	if start == end {
		y = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	y = Sum(x, start, end)/float64(end-start)
	return
}

func Variance(x []float64, start uint64, end uint64) (y float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		y = math.NaN()
		return
	}

	if start == end {
		y = math.NaN()
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
	y = s/float64(end-start)
	return
}

func StandardDeviation(x []float64, start uint64, end uint64) (y float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		y = math.NaN()
		return
	}

	if start == end {
		y = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	y = math.Pow(Variance(x, start, end), 0.5)
	return
}

func Max(x []float64, start uint64, end uint64) (y float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		y = math.NaN()
		return
	}

	if start == end {
		y = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	y = x[start]
	for i := start+1; i < end; i++ {
		if x[i] > y {
			y = x[i]
		}
	}
	return
}

func Min(x []float64, start uint64, end uint64) (y float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		y = math.NaN()
		return
	}

	if start == end {
		y = math.NaN()
		return
	}

	if start > end {
		start, end = end, start
	}

	y = x[start]
	for i := start+1; i < end; i++ {
		if x[i] < y {
			y = x[i]
		}
	}
	return
}
