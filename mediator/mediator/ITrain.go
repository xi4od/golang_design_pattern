package mediator

// 列车接口
type train interface {
	RequestArrival() // 请求到达
	Departure()      // 离开
	PermitArrival()  // 批准进入
}
