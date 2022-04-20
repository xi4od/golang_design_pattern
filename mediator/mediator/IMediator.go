package mediator

type IMediator interface {
	CanLand(train) bool
	IsLanded(train) bool
	NotifyFree()
}
