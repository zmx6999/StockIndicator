package main

import (
	"190519/indicator"
	"fmt"
)

func main()  {
	priceList := []indicator.Price{
		{230.00, 300.00, 230.00, 299.00},
		{294.00, 311.50, 290.00, 299.01},
		{299.44, 309.50, 271.50, 301.95},
		{301.95, 328.96, 284.00, 289.41},
		{289.40, 339.87, 240.00, 329.37},
		{329.38, 376.00, 322.01, 371.51},
		{371.48, 485.00, 360.00, 455.86},
		{455.86, 510.00, 387.16, 446.82},
		{446.80, 480.00, 421.18, 423.99},
		{423.99, 489.00, 394.99, 466.44},
		{466.39, 752.00, 460.00, 680.00},
		{680.01, 852.23, 663.05, 765.38},
		{766.79, 798.99, 500.00, 742.15},
		{741.49, 769.99, 650.00, 716.74},
		{716.74, 1035.00, 716.42, 982.99},
		{982.99, 1405.00, 944.00, 1340.03},
		{1340.04, 1422.00, 1089.00, 1308.50},
		{1308.50, 1308.78, 766.95, 1123.19},
		{1124.04, 1168.54, 915.00, 1038.00},
		{1038.85, 1272.88, 992.96, 1119.07},
		{1119.69, 1163.41, 790.01, 858.62},
		{857.54, 885.25, 568.00, 844.32},
		{844.99, 907.75, 778.00, 900.00},
		{899.93, 986.88, 895.00, 954.70},
		{954.90, 962.41, 784.28, 829.98},
		{829.91, 895.50, 808.10, 865.45},
		{865.45, 879.40, 819.01, 829.50},
		{829.50, 830.88, 638.21, 698.78},
		{698.78, 737.44, 571.71, 622.30},
		{622.32, 625.00, 453.53, 575.80},
		{575.80, 578.00, 476.56, 483.30},
		{483.79, 496.13, 367.07, 402.52},
		{402.50, 418.15, 358.00, 377.88},
		{377.65, 432.01, 363.21, 406.37},
		{406.23, 529.79, 405.02, 521.15},
		{521.75, 595.94, 498.15, 585.73},
		{585.88, 712.98, 573.99, 622.31},
		{622.02, 702.00, 600.11, 683.26},
		{683.26, 833.01, 630.00, 821.50},
		{821.04, 841.59, 700.84, 762.59},
		{762.59, 765.44, 623.58, 718.37},
		{718.43, 727.01, 657.00, 714.42},
		{714.39, 723.45, 549.00, 599.37},
		{599.42, 606.63, 508.00, 549.18},
		{549.18, 629.00, 542.63, 586.92},
	}
	for _, price := range priceList{
		indicator.AppendPrice(price)
	}

	indicator.InitMACD(12, 26, 9)
	for i := 1; i <= len(priceList); i++ {
		emaX, emaY, diff, dea, macd := indicator.GetMACD(uint64(i))
		indicator.AppendEmaX(emaX)
		indicator.AppendEmaY(emaY)
		indicator.AppendDiff(diff)
		indicator.AppendDea(dea)
		indicator.AppendMacd(macd)

		fmt.Println(i, macd, diff, dea)
	}

	indicator.InitRSI(14)
	for i := 1; i <= len(priceList); i++ {
		up, down, upEma, downEma, rsi := indicator.GetRSI(uint64(i))
		indicator.AppendUp(up)
		indicator.AppendDown(down)
		indicator.AppendUpEma(upEma)
		indicator.AppendDownEma(downEma)
		indicator.AppendRsi(rsi)

		fmt.Println(i, rsi)
	}

	indicator.InitBollinger(20, 2)
	for i := 1; i <= len(priceList); i++ {
		up, mid, down := indicator.GetBollinger(uint64(i))
		indicator.AppendUpBoll(up)
		indicator.AppendMidBoll(mid)
		indicator.AppendDownBoll(down)

		fmt.Println(i, mid, up, down)
	}
}
