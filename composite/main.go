package main

import (
	"composite/company"
	"fmt"
)

func main(){
	mCompany := company.NewCompany("总公司")
	mCompany.Add(company.NewHR("总公司人力资源部"))
	mCompany.Add(company.NewFinance("总公司财务部"))

	mCompanyHD := company.NewCompany("上海华东分公司")
	mCompanyHD.Add(company.NewHR("华东分公司人力资源部"))
	mCompanyHD.Add(company.NewFinance("华东分公司财务部"))
	mCompany.Add(mCompanyHD)

	mCompanyNJ := company.NewCompany("南京办事处")
	mCompanyNJ.Add(company.NewHR("南京办事处人力资源部"))
	mCompanyNJ.Add(company.NewFinance("南京办事处财务部"))
	mCompanyHD.Add(mCompanyNJ)

	mCompanyHZ := company.NewCompany("杭州办事处")
	mCompanyHZ.Add(company.NewHR("杭州办事处人力资源部"))
	mCompanyHZ.Add(company.NewFinance("杭州办事处财务部"))
	mCompanyHD.Add(mCompanyHZ)

	fmt.Println("组织结构图:")
	mCompany.Display(1)

	fmt.Println("职责:")
	mCompany.LineOfDuty()
}