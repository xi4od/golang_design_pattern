package product_detail

import (
	"facade/product"
	"facade/promotion"
	"facade/stock"
)

type ProductDetail struct {
	Name string
	Price float64
	Debate float64
	Num int32
}

func (t ProductDetail) GetProductDetail(id int32) IProductDetail {
	productInfo := new(product.Product).GetProductInfo(id)
	promotionInfo := new(promotion.Promotion).GetPromotionInfo(id)
	stockInfo := new(stock.Stock).GetStockInfo(id)
	t.Name = productInfo.Name
	t.Price = productInfo.Price
	t.Debate = promotionInfo.Debate
	t.Num = stockInfo.Num
	return t
}
