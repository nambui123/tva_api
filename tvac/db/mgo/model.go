package mgo

import (
	"github.com/satori/go.uuid"
	"time"
)

type IModel interface {
	GetID() string
	SetID(id string)
	IsDeleted() bool
	BeforeCreate()
	BeforeUpdate()
	BeforeDelete()
}

type BaseModel struct {
	ID    string `bson:"_id" json:"id"`      //
	CTime int64  `bson:"ctime" json:"ctime"` //Modify Time
	DTime int64  `bson:"dtime" json:"dtime"` //Delete Time
}

func (b *BaseModel) GetID() string {
	return b.ID
}

func (b *BaseModel) SetID(s string) {
	b.ID = s
}

func (b *BaseModel) IsDeleted() bool {
	return b.DTime > 0
}

func (b *BaseModel) BeforeCreate() {
	b.ID = uuid.NewV4().String()
	b.CTime = time.Now().Unix()
	b.DTime = 0
}

func (b *BaseModel) BeforeUpdate() {
	b.CTime = time.Now().Unix()
}

func (b *BaseModel) BeforeDelete() {

}
