package user

import (
	"tva_api/db/mgo"
)

type User struct {
	mgo.BaseModel `bson:",inline"`
	UserName      string `bson:"username" json"username"`
	Email         string `bson:"email" json:"email"`
	Password      string `bson:"password" json:"password"`
	PhoneNumber   string `bson:"phone_number" json:"phone_number"`
	AvataUrl      string `bson:"avata_url" json:"avata_url"`
	Address       string `bson:"address" json:"address"`
	Role          string `bson:"role" json:"role"`
}
