package main

import (
	"190518/indicator"
	"fmt"
)

func main()  {
	priceList := []float64{299.00, 299.01, 301.95, 289.41, 329.37, 371.51, 455.86, 446.82, 423.99, 466.44, 680.00, 765.38, 742.15, 716.74, 982.99, 1340.03, 1308.50, 1123.19, 1038.00, 1119.07, 858.62, 844.32, 900.00, 954.70, 829.98, 865.45, 829.50, 698.78, 622.30, 575.80, 483.30, 402.52, 377.88, 406.37, 521.15, 585.73, 622.31, 683.26, 821.50, 762.59, 718.37, 714.42, 599.37, 549.18, 586.92, 602.59, 464.22}
	for _, price := range priceList{
		indicator.AppendPrice(price)
	}

	indicator.InitMACD(12, 26, 9)
	for i := 1; i <= len(priceList); i++ {
		emax, emay, diff, dea, macd := indicator.GetMACD(uint64(i))
		indicator.AppendEmax(emax)
		indicator.AppendEmay(emay)
		indicator.AppendDiff(diff)
		indicator.AppendDea(dea)
		indicator.AppendMacd(macd)

		fmt.Println(i, macd, diff, dea)
	}

	indicator.InitRSI(14)
	for i := 1; i <= len(priceList); i++ {
		up, down, upEMA, downEMA, rsi := indicator.GetRSI(uint64(i))
		indicator.AppendUp(up)
		indicator.AppendDown(down)
		indicator.AppendUpEMA(upEMA)
		indicator.AppendDownEMA(downEMA)
		indicator.AppendRSI(rsi)

		fmt.Println(i, rsi)
	}

	indicator.InitBoll(20, 2)
	for i := 1; i <= len(priceList); i++ {
		upBoll, midBoll, downBoll := indicator.GetBoll(uint64(i))
		indicator.AppendUpBoll(upBoll)
		indicator.AppendMidBoll(midBoll)
		indicator.AppendDownBoll(downBoll)

		fmt.Println(i, midBoll, upBoll, downBoll)
	}
}
