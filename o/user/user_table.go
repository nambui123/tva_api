package user

import (
	"github.com/golang/glog"
	"tva_api/o/model"
)

var TableUser = model.NewTable("users")

func (b *User) Create() error {
	var p = password(b.Password)
	// replace
	if err := p.HashTo(&b.Password); err != nil {
		return err
	}
	return TableUser.Create(b)
}

func MarkDelete(id string) error {
	return TableUser.MarkDelete(id)
}

func (v *User) Update(newValue *User) error {
	return TableUser.UnsafeUpdateByID(v.ID, newValue)
}

func (v *User) UpdatePass(newValue string) error {

	return TableUser.UnsafeUpdateByID(v.ID, newValue)
}

func init() {
	if err := TableUser.EnsureIndex("branch_id,"); err != nil {
		glog.Error("user index error", err)
	}
}
