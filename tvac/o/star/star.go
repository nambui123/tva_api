package star

import (
	"tvac/db/mgo"
)

type Star struct {
	mgo.BaseModel  `bson:",inline"`
	PostID         string `bson:"post_id" json:"post_id"`
	NumberLike     int    `bson:"number_like" json:"number_like"`
	NumberWatch    int    `bson:"number_watch" json:"number_watch"`
	NumberFeedback int    `bson:"number_feedback" json:"number_feedback"`
	TotalFeedback  int    `bson:"total_feedback" json:"total_feedback"`
}
