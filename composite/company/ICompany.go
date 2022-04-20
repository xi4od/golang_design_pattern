package company

type ICompany interface {
	Add(ICompany)
	Remove(ICompany)
	Display(int)
	LineOfDuty()
}
