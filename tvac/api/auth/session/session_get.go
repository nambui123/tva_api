package session

import (
	"github.com/golang/glog"
	"tvac/o/session"
	"tvac/web"
)

const (
	errReadSessonFailed   = web.InternalServerError("read session failed")
	errSessionNotFound    = web.Unauthorized("session not found")
	errUnauthorizedAccess = web.Unauthorized("unauthorized access")
)

func Get(sessionID string) (*session.Session, error) {
	var s, err = session.GetByID(sessionID)
	if err != nil {
		if session.TableSession.IsErrNotFound(err) {
			return nil, errSessionNotFound
		}
		glog.Error(err)
		return nil, errReadSessonFailed
	}

	return s, nil
}
