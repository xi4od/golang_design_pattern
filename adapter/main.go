package main

import (
	"adapter/data"
	"adapter/db"
)

func main() {
	//before
	var mDb db.IDb
	mDb = db.Mysql{}
	mDb.SaveData()

	//after
	mData := data.DataAdapter{}
	mData.SaveData()
}
