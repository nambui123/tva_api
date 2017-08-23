package content

import (
	"tvac/db/mgo"
)

type Content struct {
	mgo.BaseModel `bson:",inline"`
	Text          string `bson:"text" json:"text"`
	UrlVideo      string `bson:"url_video" json:"url_video"`
	UrlIamge      string `bson:"url_image" json:"url_iamge"`
}
