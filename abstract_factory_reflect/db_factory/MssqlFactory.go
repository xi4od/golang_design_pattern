package db_factory

import "abstract_factory_reflect/model"

type MssqlFactory struct {

}

func(t MssqlFactory) CreateUser() model.IUser {
	return model.MssqlUser{}
}

func(t MssqlFactory) CreateDepart() model.IDepart {
	return model.MssqlDepart{}
}
