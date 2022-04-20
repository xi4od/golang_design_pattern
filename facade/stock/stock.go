package stock

type Stock struct {
	Id int32
	Name string
	Num int32
}

func(t Stock) GetStockInfo(id int32) Stock {
	if id == 1 {
		return Stock{1, "ipad", 78}
	} else {
		return Stock{2, "iphone", 67}
	}
}
