package indicator

import (
	"strconv"
	"fmt"
)

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

func getHighPriceList(x []Price, start uint64, end uint64) (y []float64) {
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
		y = append(y, x[i].High)
	}
	return
}

func getLowPriceList(x []Price, start uint64, end uint64) (y []float64) {
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
		y = append(y, x[i].Low)
	}
	return
}

func ClosePriceSum(x []Price, start uint64, end uint64) (z float64) {
	y := getClosePriceList(x, start, end)
	z = Sum(y, 0, uint64(len(y)))
	return
}

func ClosePriceAvg(x []Price, start uint64, end uint64) (z float64) {
	y := getClosePriceList(x, start, end)
	z = Avg(y, 0, uint64(len(y)))
	return
}

func ClosePriceVariance(x []Price, start uint64, end uint64) (z float64) {
	y := getClosePriceList(x, start, end)
	z = Variance(y, 0, uint64(len(y)))
	return
}

func ClosePriceStandardDeviation(x []Price, start uint64, end uint64) (z float64) {
	y := getClosePriceList(x, start, end)
	z = StandardDeviation(y, 0, uint64(len(y)))
	return
}

func HighPrice(x []Price, start uint64, end uint64) (z float64) {
	y := getHighPriceList(x, start, end)
	z = Max(y, 0, uint64(len(y)))
	return
}

func LowPrice(x []Price, start uint64, end uint64) (z float64) {
	y := getLowPriceList(x, start, end)
	z = Min(y, 0, uint64(len(y)))
	return
}

func ToFix(x float64, decimal uint64) (y float64) {
	y, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(decimal))+"f", x), 64)
	return
}
