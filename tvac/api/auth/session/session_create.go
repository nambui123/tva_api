package session

import (
	"github.com/golang/glog"
	"tvac/o/session"
	"tvac/o/user"
	"tvac/web"
)

func New(u *user.User) (*session.Session, error) {

	var s = &session.Session{
		UserID:   u.ID,
		Username: u.UserName,
	}

	var err = s.Create()
	if err != nil {
		glog.Error(err)
		return nil, web.InternalServerError("save session failed")
	}
	return s, nil
}

func MustNew(u *user.User) *session.Session {
	s, e := New(u)
	if e != nil {
		panic(e)
	}
	return s
}
