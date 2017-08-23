package post

import (
	"tvac/db/mgo"
)

type Post struct {
	mgo.BaseModel `bson:",inline"`
	Title         string `bson:"title" json"title"`
	UserID        string `bson:"user_id" json:"user_id"`
	Detail        string `bson:"detail" json:"detail"`
	CategoryId    string `bson:"category_id" json:"category_id"`
	ContentId     string `bson:"content_id" json:"content_id"`
	TagId         string `bson:"tag_id" json:"tag_id"`
}
