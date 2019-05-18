package indicator

import "math"

var (
	upBollList, midBollList, downBollList []float64
	bolln, bollk uint64
)

func InitBoll(n, k uint64)  {
	bolln, bollk = n, k
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

func GetBoll(currentIndex uint64) (upBoll, midBoll, downBoll float64) {
	upBoll = math.NaN()
	midBoll = math.NaN()
	downBoll = math.NaN()
	if currentIndex < bolln {
		return
	}

	midBoll = toFixed(avg(priceList, currentIndex-bolln, currentIndex), 4)
	v := standardDeviation(priceList, currentIndex-bolln, currentIndex)
	upBoll = toFixed(midBoll+v*float64(bollk), 4)
	downBoll = toFixed(midBoll-v*float64(bollk), 4)
	return
}
