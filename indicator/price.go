package indicator

var (
	priceList []float64
)

func AppendPrice(price float64)  {
	priceList = append(priceList, price)
}
