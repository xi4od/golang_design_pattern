package manager

import "fmt"

type CommonManager struct {
	Manager
}

func (t CommonManager) Request(request Request) {
	if request.Type == "请假" && request.Num <= 2 {
		fmt.Printf("\n%v 批准: %v,数量: %v, 内容: %v", t.Name, request.Type, request.Num, request.Content)
	} else {
		fmt.Printf("\n%v 权限不够, 转上级审批", t.Name)
		t.Superior.Request(request)
	}
}
