package main

import "response_chain/manager"

func main() {
	manager1 := manager.Manager{Name: "刘文杰"}
	wenjie := manager.CommonManager{manager1}

	manager2 := manager.Manager{Name: "王振飞"}
	zhenfei := manager.MajorManager{manager2}

	manager3 := manager.Manager{Name: "杜磊"}
	dulei := manager.GeneralManager{manager3}

	wenjie.SetSuperior(&zhenfei)
	zhenfei.SetSuperior(&dulei)

	request := manager.Request{
		Type:    "请假",
		Num:     3,
		Content: "田国华请假3天",
	}
	wenjie.Request(request)

	request1 := manager.Request{
		Type:    "加薪",
		Num:     500,
		Content: "田国华申请加薪500",
	}
	wenjie.Request(request1)
}
