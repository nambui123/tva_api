package config

import (
	"fmt"
	"os"
	"tvac/db/mgo"
)

func init() {
	var _, err = mgo.NewDB(DBHost, DBName)
	if err != nil {
		fmt.Println("Cannot connect to db at ", DBHost)
	}
	os.Setenv(KeyDB, DBName)
}
