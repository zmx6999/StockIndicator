package indicator

import "math"

var (
	upList, downList, upEMAList, downEMAList, rsiList []float64
	rsin uint64
)

func InitRSI(n uint64)  {
	rsin = n
}

func AppendUp(up float64)  {
	upList = append(upList, up)
}

func AppendDown(down float64)  {
	downList = append(downList, down)
}

func AppendUpEMA(upEMA float64)  {
	upEMAList = append(upEMAList, upEMA)
}

func AppendDownEMA(downDEA float64)  {
	downEMAList = append(downEMAList, downDEA)
}

func AppendRSI(rsi float64)  {
	rsiList = append(rsiList, rsi)
}

func GetUpDownEMA(currentIndex uint64, upDownList []float64, currentUpDown float64, emaList []float64) (ema float64) {
	if currentIndex <= rsin {
		ema = math.NaN()
		return
	}

	if currentIndex == rsin+1 {
		 _upDownList := make([]float64, rsin)
		 copy(_upDownList, upDownList[currentIndex-rsin:])
		 _upDownList[rsin-1] = currentUpDown
		 ema = toFixed(avg(_upDownList, 0, rsin), 4)
		 return
	}

	last := len(emaList)-1
	lastEMA := emaList[last]
	ema = toFixed((lastEMA*float64(rsin-1)+currentUpDown)/float64(rsin), 4)
	return
}

func GetRSI(currentIndex uint64) (up, down, upEMA, downEMA, rsi float64) {
	up = math.NaN()
	down = math.NaN()
	upEMA = math.NaN()
	downEMA = math.NaN()
	rsi = math.NaN()

	if currentIndex <= 1 {
		return
	}

	if priceList[currentIndex-1] > priceList[currentIndex-2] {
		up = priceList[currentIndex-1]-priceList[currentIndex-2]
		down = 0
	} else {
		up = 0
		down = priceList[currentIndex-2]-priceList[currentIndex-1]
	}

	upEMA = GetUpDownEMA(currentIndex, upList, up, upEMAList)
	downEMA = GetUpDownEMA(currentIndex, downList, down, downEMAList)
	if upEMA == math.NaN() || downEMA == math.NaN() {
		return
	}
	rsi = toFixed(upEMA/(upEMA+downEMA)*100, 4)
	return
}
