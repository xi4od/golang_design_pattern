package product

type Product struct {
	Id int32
	Name string
	Price float64
}

func(t Product) GetProductInfo(id int32) Product {
	if id == 1 {
		return Product{1, "ipad", 2000}
	} else {
		return Product{2, "iphone", 6000}
	}
}
