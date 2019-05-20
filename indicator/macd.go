package indicator

import "math"

var (
	macdX, macdY, macdZ uint64
	emaXList, emaYList, diffList, deaList, macdList []float64
)

func InitMACD(x, y, z uint64)  {
	if x > y {
		x, y = y, x
	}
	macdX, macdY, macdZ = x, y, z
}

func AppendEmaX(emaX float64)  {
	emaXList = append(emaXList, emaX)
}

func AppendEmaY(emaY float64)  {
	emaYList = append(emaYList, emaY)
}

func AppendDiff(diff float64)  {
	diffList = append(diffList, diff)
}

func AppendDea(dea float64)  {
	deaList = append(deaList, dea)
}

func AppendMacd(macd float64)  {
	macdList = append(macdList, macd)
}

func GetEMA(currentIndex uint64, emaList []float64, currentPrice float64, n uint64) (ema float64) {
	if currentIndex < n {
		ema = math.NaN()
		return
	}

	if currentIndex == n {
		ema = ClosePriceAvg(priceList, 0, n)
		return
	}

	lastEma := emaList[len(emaList)-1]
	ema = (lastEma*float64(n-1)+currentPrice*2)/float64(n+1)
	return
}

func GetDEA(currentIndex uint64, currentDiff float64) (dea float64) {
	n := macdY+macdZ-1
	if currentIndex < n {
		dea = math.NaN()
		return
	}

	if currentIndex == n {
		_diffList := make([]float64, macdZ)
		copy(_diffList, diffList[macdY-1:])
		_diffList[macdZ-1] = currentDiff
		dea = Avg(_diffList,0, macdZ)
		return
	}

	lastDea := deaList[len(deaList)-1]
	dea = (lastDea*float64(macdZ-1)+currentDiff*2)/float64(macdZ+1)
	return
}

func GetMACD(currentIndex uint64) (emaX, emaY, diff, dea, macd float64) {
	diff = math.NaN()
	dea = math.NaN()
	macd = math.NaN()

	emaX = GetEMA(currentIndex, emaXList, priceList[currentIndex-1].Close, macdX)
	emaY = GetEMA(currentIndex, emaYList, priceList[currentIndex-1].Close, macdY)
	if emaX == math.NaN() || emaY == math.NaN() {
		return
	}

	diff = emaX-emaY
	dea = GetDEA(currentIndex, diff)
	if dea == math.NaN() {
		return
	}

	macd = diff-dea
	return
}
