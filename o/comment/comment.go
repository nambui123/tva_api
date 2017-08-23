package comment

import (
	"tva_api/o/content"
)

type Comment struct {
	ID      string          `bson:"_id" json:"id"`
	UserID  string          `bson:"user_id" json"user_id"`
	PostID  string          `bson:"post_id" json:"post_id"`
	Content content.Content `bson:"content" json:"content"`
}
