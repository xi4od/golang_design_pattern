package db_factory

import "abstract_factory/model"

type IFactory interface {
	CreateUser() model.IUser
	CreateDepart() model.IDepart
}
