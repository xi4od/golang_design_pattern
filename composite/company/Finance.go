package company

import "fmt"

type Finance struct {
	Company
}

func NewFinance(name string) *Finance {
	return &Finance{Company: Company{name: name}}
}

func (c Finance) Display(depth int) {
	nOfLine := ""
	for n:=0; n<depth;n++{
		nOfLine += "-"
	}
	fmt.Printf("%v%v\n", nOfLine, c.name)
}

func (c Finance) LineOfDuty() {
	fmt.Printf("%v 公司财务管理\n", c.name)
}


