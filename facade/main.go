package main

import (
	"facade/product_detail"
	"fmt"
)

func main(){
	var productDetail  product_detail.IProductDetail
	productDetail = product_detail.ProductDetail{}
	fmt.Println(productDetail.GetProductDetail(2))
}
