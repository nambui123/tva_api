package model

import (
	"os"
	"tva_api/config"
	"tva_api/db/mgo"
)

func NewTable(name string) *mgo.Table {
	var db = GetDB()
	return mgo.NewTable(db, name)
}

func GetDB() *mgo.Database {
	return mgo.GetDB(os.Getenv(config.KeyDB))
}
