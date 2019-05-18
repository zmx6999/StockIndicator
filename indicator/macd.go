package indicator

import "math"

var (
	emaxList, emayList, diffList, deaList, macdList []float64
	macdx, macdy, macdz uint64
)

func InitMACD(x, y, z uint64)  {
	if x > y {
		x, y = y, x
	}
	macdx, macdy, macdz = x, y, z
}

func AppendEmax(emax float64)  {
	emaxList = append(emaxList, emax)
}

func AppendEmay(emay float64)  {
	emayList = append(emayList, emay)
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
		ema = toFixed(avg(priceList, 0, n), 4)
		return
	}

	last := len(emaList)-1
	lastEMA := emaList[last]
	ema = toFixed((lastEMA*float64(n-1)+priceList[currentIndex-1]*2)/float64(n+1), 4)
	return
}

func GetDEA(currentIndex uint64, deaList []float64, currentDiff float64) (dea float64) {
	n := macdy+macdz-1
	if currentIndex < n {
		dea = math.NaN()
		return
	}

	if currentIndex == n {
		_diffList := make([]float64, macdz)
		copy(_diffList, diffList[currentIndex-macdz:])
		_diffList[macdz-1] = currentDiff
		dea = toFixed(avg(_diffList, 0, macdz), 4)
		return
	}

	last := len(deaList)-1
	lastDEA := deaList[last]
	dea = toFixed((lastDEA*float64(macdz-1)+currentDiff*2)/float64(macdz+1), 4)
	return
}

func GetMACD(currentIndex uint64) (emax, emay, diff, dea, macd float64) {
	diff = math.NaN()
	dea = math.NaN()
	macd = math.NaN()

	emax = GetEMA(currentIndex, emaxList, macdx)
	emay = GetEMA(currentIndex, emayList, macdy)
	if emax == math.NaN() || emay == math.NaN() {
		return
	}

	diff = toFixed(emax-emay, 4)
	dea = GetDEA(currentIndex, deaList, diff)
	if dea == math.NaN() {
		return
	}

	macd = toFixed(diff-dea, 4)
	return
}
