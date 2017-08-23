package tag

import (
	"tva_api/db/mgo"
)

type Tag struct {
	mgo.BaseModel `bson:",inline"`
	Name          []string `bson:"name" json"name"`
}
