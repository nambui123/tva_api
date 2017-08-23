package tag

import (
	"tvac/db/mgo"
)

type Tag struct {
	mgo.BaseModel `bson:",inline"`
	Name          []string `bson:"name" json"name"`
}
