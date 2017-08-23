package studio

import (
	"tva_api/db/mgo"
)

type Studio struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string `bson:"user_id" json:"user_id"`
	Name          string `bson:"name" json"name"`
	Detail        string `bson:"detail" json:"detail"`
	PhoneNumber   string `bson:"phone_number" json:"phone_number"`
	Address       string `bson:"address" json:"address"`
}
