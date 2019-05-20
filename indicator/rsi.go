package indicator

import "math"

var (
	rsiN uint64
	upList, downList, upEmaList, downEmaList, rsiList []float64
)

func InitRSI(n uint64)  {
	rsiN = n
}

func AppendUp(up float64)  {
	upList = append(upList, up)
}

func AppendDown(down float64)  {
	downList = append(downList, down)
}

func AppendUpEma(upEma float64)  {
	upEmaList = append(upEmaList, upEma)
}

func AppendDownEma(downEma float64)  {
	downEmaList = append(downEmaList, downEma)
}

func AppendRSI(rsi float64)  {
	rsiList = append(rsiList, rsi)
}

func GetUpDownEMA(currentIndex uint64, upDownList []float64, upDownEmaList []float64, currentUpDown float64) (ema float64) {
	if currentIndex <= rsiN {
		ema = math.NaN()
		return
	}

	if currentIndex == rsiN+1 {
		_upDownList := make([]float64, rsiN)
		copy(_upDownList, upDownList[1:])
		_upDownList[rsiN-1] = currentUpDown
		ema = Avg(_upDownList, 0, rsiN)
		return
	}

	lastEma := upDownEmaList[len(upDownEmaList)-1]
	ema = (lastEma*float64(rsiN-1)+currentUpDown)/float64(rsiN)
	return
}

func GetRSI(currentIndex uint64) (up, down, upEma, downEma, rsi float64) {
	up = math.NaN()
	down = math.NaN()
	upEma = math.NaN()
	downEma = math.NaN()
	rsi = math.NaN()

	if currentIndex <= 1 {
		return
	}

	if priceList[currentIndex-1].Close > priceList[currentIndex-2].Close {
		up = priceList[currentIndex-1].Close-priceList[currentIndex-2].Close
		down = 0
	} else {
		up = 0
		down = priceList[currentIndex-2].Close-priceList[currentIndex-1].Close
	}

	upEma = GetUpDownEMA(currentIndex, upList, upEmaList, up)
	downEma = GetUpDownEMA(currentIndex, downList, downEmaList, down)
	if upEma == math.NaN() || downEma == math.NaN() {
		return
	}

	rsi =upEma/(upEma+downEma)*100
	return
}
