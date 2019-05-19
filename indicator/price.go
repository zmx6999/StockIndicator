package indicator

type Price struct {
	Open float64
	High float64
	Low float64
	Close float64
}

var (
	priceList []Price
)

func AppendPrice(price Price)  {
	priceList = append(priceList, price)
}

func getClosePriceList(x []Price, start uint64, end uint64) (y []float64) {
	if start > uint64(len(x)) || end > uint64(len(x)) {
		return
	}

	if start == end {
		return
	}

	if start > end {
		start, end = end, start
	}

	for i := start; i < end; i++ {
		y = append(y, x[i].Close)
	}

	return
}

func ClosePriceSum(x []Price, start uint64, end uint64) (s float64) {
	y := getClosePriceList(x, start, end)
	s = Sum(y, 0, uint64(len(y)))
	return
}

func ClosePriceAvg(x []Price, start uint64, end uint64) (a float64) {
	y := getClosePriceList(x, start, end)
	a = Avg(y, 0, uint64(len(y)))
	return
}

func ClosePriceVariance(x []Price, start uint64, end uint64) (v float64) {
	y := getClosePriceList(x, start, end)
	v = Variance(y, 0, uint64(len(y)))
	return
}

func ClosePriceStandardDeviation(x []Price, start uint64, end uint64) (s float64) {
	y := getClosePriceList(x, start, end)
	s = StandardDeviation(y, 0, uint64(len(y)))
	return
}
