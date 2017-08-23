package session

import (
	"encoding/json"
	"tvac/db/mgo"
)

type Session struct {
	mgo.BaseModel `bson:",inline"`
	Username      string `bson:"username" json:"username"`
	UserID        string `bson:"userid" json:"userid"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
