package indicator

import "math"

var (
	bollN, bollK uint64
	upBollList, midBollList, downBollList []float64
)

func InitBollinger(n, k uint64)  {
	bollN, bollK = n, k
}

func AppendUpBoll(upBoll float64)  {
	upBollList = append(upBollList, upBoll)
}

func AppendMidBoll(midBoll float64)  {
	midBollList = append(midBollList, midBoll)
}

func AppendDownBoll(downBoll float64)  {
	downBollList = append(downBollList, downBoll)
}

func GetBollinger(currentIndex uint64) (up, mid, down float64) {
	up = math.NaN()
	mid = math.NaN()
	down = math.NaN()

	if currentIndex < bollN {
		return
	}

	mid = ToFixed(ClosePriceAvg(priceList, currentIndex-bollN, currentIndex), 4)
	sd := ClosePriceStandardDeviation(priceList, currentIndex-bollN, currentIndex)
	up = ToFixed(mid+sd*float64(bollK), 4)
	down = ToFixed(mid-sd*float64(bollK), 4)
	return
}
