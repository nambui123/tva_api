package chat

import (
	"tvac/o/content"
)

type Chat struct {
	ID         string          `bson:"_id" json:"id"`
	SenderId   string          `bson:"sender_id" json"sender_id"`
	ReceiverId string          `bson:"receiver_id" json:"receiver_id"`
	Content    content.Content `bson:"content" json:"content"`
}
