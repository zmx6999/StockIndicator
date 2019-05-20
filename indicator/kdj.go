package indicator

import (
	"math"
	)

var (
	kdjN, kdjK, kdjD uint64
	rsvList, kList, dList, jList []float64
)

func InitKDJ(n, k, d uint64)  {
	kdjN, kdjK, kdjD = n, k, d
}

func AppendRSV(rsv float64)  {
	rsvList = append(rsvList, rsv)
}

func AppendK(k float64)  {
	kList = append(kList, k)
}

func AppendD(d float64)  {
	dList = append(dList, d)
}

func AppendJ(j float64)  {
	jList = append(jList, j)
}

func GetRSV(currentIndex uint64) (rsv float64) {
	if currentIndex < kdjN {
		rsv = math.NaN()
		return
	}

	high := HighestPrice(priceList, currentIndex-kdjN, currentIndex)
	low := LowestPrice(priceList, currentIndex-kdjN, currentIndex)
	cl := priceList[currentIndex-1].Close
	rsv = ToFixed((cl-low)/(high-low)*100, 4)
	return
}

func GetK(currentIndex uint64, currentRSV float64) (k float64) {
	n := kdjN+kdjK-1
	if currentIndex < n {
		k = math.NaN()
		return
	}

	_rsvList := make([]float64, kdjK)
	copy(_rsvList, rsvList[currentIndex-kdjK:])
	_rsvList[kdjK-1] = currentRSV
	k = ToFixed(Avg(_rsvList, 0, kdjK), 4)
	return
}

func GetD(currentIndex uint64, currentK float64) (d float64) {
	n := kdjN+kdjK+kdjD-2
	if currentIndex < n {
		d = math.NaN()
		return
	}

	_kList := make([]float64, kdjD)
	copy(_kList, kList[currentIndex-kdjD:])
	_kList[kdjD-1] = currentK
	d = ToFixed(Avg(_kList, 0, kdjD), 4)
	return
}

func GetKDJ(currentIndex uint64) (rsv, k, d, j float64) {
	k = math.NaN()
	d = math.NaN()
	j = math.NaN()
	rsv = GetRSV(currentIndex)
	if rsv == math.NaN() {
		return
	}

	k = GetK(currentIndex, rsv)
	if k == math.NaN() {
		return
	}

	d = GetD(currentIndex, k)
	if d == math.NaN() {
		return
	}

	j = ToFixed(k*3-d*2, 4)
	return
}
