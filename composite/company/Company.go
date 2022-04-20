package company

import (
	"container/list"
	"fmt"
	"reflect"
)

type Company struct {
	Companys *list.List
	name string
}

func NewCompany(name string) *Company {
	return &Company{name: name, Companys: list.New()}
}

func (c *Company) Add(company ICompany) {
	c.Companys.PushBack(company)
}

func (c *Company) Remove(company ICompany) {
	for i := c.Companys.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, company) {
			c.Companys.Remove(i)
		}
	}
}

func (c Company) Display(depth int) {
	nOfLine := ""
	for n:=0; n<depth;n++{
		nOfLine += "-"
	}
	fmt.Printf("%v%v\n", nOfLine, c.name)
	for i := c.Companys.Front(); i != nil; i = i.Next() {
		i.Value.(ICompany).Display(depth + 2)
	}
}

func (c Company) LineOfDuty() {
	for i := c.Companys.Front(); i != nil; i = i.Next() {
		i.Value.(ICompany).LineOfDuty()
	}
}





