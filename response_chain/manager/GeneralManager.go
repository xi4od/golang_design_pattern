package manager

import "fmt"

type GeneralManager struct {
	Manager
}

func (t GeneralManager) Request(request Request) {
	if request.Type == "请假" {
		fmt.Printf("\n%v 批准: %v,数量: %v, 内容: %v", t.Name, request.Type, request.Num, request.Content)
	} else if request.Type == "加薪" && request.Num <= 500 {
		fmt.Printf("\n%v 批准: %v,数量: %v, 内容: %v", t.Name, request.Type, request.Num, request.Content)
	} else if request.Type == "加薪" && request.Num > 500 {
		fmt.Printf("\n%v 拒绝: %v,数量: %v, 内容: %v", t.Name, request.Type, request.Num, request.Content)
	} else {
		fmt.Printf("\n%v 拒绝: %v,数量: %v, 内容: %v", t.Name, request.Type, request.Num, request.Content)
	}
}
