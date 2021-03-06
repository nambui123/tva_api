package category

import (
	"tva_api/db/mgo"
)

type Category struct {
	mgo.BaseModel `bson:",inline"`
	Name          string `bson:"name" json:"name"`
	Detail        string `bson:"detail" json:"detail"`
}
