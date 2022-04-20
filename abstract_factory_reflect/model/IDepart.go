package model

type IDepart interface {
	Insert(depart Depart)
	GetDepart(id int)
}
