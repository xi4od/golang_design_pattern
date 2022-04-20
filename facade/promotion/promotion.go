package promotion

type Promotion struct {
	Id int32
	Name string
	Debate float64
}

func(t Promotion) GetPromotionInfo(id int32) Promotion {
	if id == 1 {
		return Promotion{1, "ipad", 0.8}
	} else {
		return Promotion{2, "iphone", 0.9}
	}
}
