package indicator

import "math"

var (
	bollingerN, bollingerK uint64
	bollingerUpList, bollingerMidList, bollingerDownList []float64
)

func InitBollinger(n, k uint64)  {
	bollingerN, bollingerK = n, k
}

func AppendBollingerUp(up float64)  {
	bollingerUpList = append(bollingerUpList, up)
}

func AppendBollingerMid(mid float64)  {
	bollingerMidList = append(bollingerMidList, mid)
}

func AppendBollingerDown(down float64)  {
	bollingerDownList = append(bollingerDownList, down)
}

func GetBollinger(currentIndex uint64) (up, mid, down float64) {
	up = math.NaN()
	mid = math.NaN()
	down = math.NaN()

	if currentIndex < bollingerN {
		return
	}

	mid = ClosePriceAvg(priceList, currentIndex-bollingerN, currentIndex)
	sd := ClosePriceStandardDeviation(priceList, currentIndex-bollingerN, currentIndex)
	up = mid+sd*float64(bollingerK)
	down = mid-sd*float64(bollingerK)
	return
}
