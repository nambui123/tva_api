package content

import (
	"tvac/db/mgo"
)

type Follow struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string `bson:"user_id" json:"user_id"`
	IFollow       string `bson:"i_follow" json:"i_follow"`
	FollowMe      string `bson:"follow_me" json:"follow_me"`
}
