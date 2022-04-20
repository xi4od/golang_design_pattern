package db_factory

import "abstract_factory_reflect/model"

type IFactory interface {
	CreateUser() model.IUser
	CreateDepart() model.IDepart
}
