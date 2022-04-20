package db_factory

import "abstract_factory/model"

type MssqlFactory struct {

}

func(t MssqlFactory) CreateUser() model.IUser {
	return model.MssqlUser{}
}

func(t MssqlFactory) CreateDepart() model.IDepart {
	return model.MssqlDepart{}
}
