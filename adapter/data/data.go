package data

import (
	"adapter/cache"
	"adapter/db"
)

type DataAdapter struct {

}

func(t DataAdapter)SaveData(){
	var mCache cache.ICache
	mCache = cache.Redis{}
	mCache.Save2Cache()

	var mIDb db.IDb
	mIDb = db.Mysql{}
	mIDb.SaveData()
}