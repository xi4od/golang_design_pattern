package company

import "fmt"

type HR struct {
	Company
}

func NewHR(name string) *HR {
	return &HR{Company: Company{name: name}}
}

func (c HR) Display(depth int) {
	nOfLine := ""
	for n:=0; n<depth;n++{
		nOfLine += "-"
	}
	fmt.Printf("%v%v\n", nOfLine, c.name)
}

func (c HR) LineOfDuty() {
	fmt.Printf("%v 员工招聘培训管理\n", c.name)
}


