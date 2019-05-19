package indicator

import "math"

var (
	macdX, macdY, macdZ uint64
	emaXList, emaYList, diffList, deaList, macdList []float64
)

func InitMACD(x, y, z uint64)  {
	if x > y {
		x, y = y, z
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

func GetEMA(currentIndex uint64, emaList []float64, n uint64) (ema float64) {
	if currentIndex < n {
		ema = math.NaN()
		return
	}

	if currentIndex == n {
		ema = ToFixed(ClosePriceAvg(priceList, 0, n), 4)
		return
	}

	lastEma := emaList[len(emaList)-1]
	ema = ToFixed((lastEma*float64(n-1)+priceList[currentIndex-1].Close*2)/float64(n+1), 4)
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
		dea = ToFixed(Avg(_diffList, 0, macdZ), 4)
		return
	}

	lastDea := deaList[len(deaList)-1]
	dea = ToFixed((lastDea*float64(macdZ-1)+currentDiff*2)/float64(macdZ+1), 4)
	return
}

func GetMACD(currentIndex uint64) (emaX, emaY, diff, dea, macd float64) {
	diff = math.NaN()
	dea = math.NaN()
	macd = math.NaN()

	emaX = GetEMA(currentIndex, emaXList, macdX)
	emaY = GetEMA(currentIndex, emaYList, macdY)
	if emaX == math.NaN() || emaY == math.NaN() {
		return
	}

	diff = ToFixed(emaX-emaY, 4)
	dea = GetDEA(currentIndex, diff)
	if dea == math.NaN() {
		return
	}

	macd = ToFixed(diff-dea, 4)
	return
}
