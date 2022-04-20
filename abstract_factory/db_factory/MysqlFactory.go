package db_factory

import "abstract_factory/model"

type MysqlFactory struct {

}

func(t MysqlFactory) CreateUser() model.IUser {
	return model.MysqlUser{}
}

func(t MysqlFactory) CreateDepart() model.IDepart {
	return model.MysqlDepart{}
}

