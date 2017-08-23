package session

import (
	"github.com/golang/glog"
	"net/http"
	"tva_api/o/session"
)

var accessToken = "token"

func MustGet(r *http.Request) *session.Session {
	var sessionID = r.URL.Query().Get(accessToken)
	var s, e = Get(sessionID)
	if e != nil {
		panic(e)
	}
	return s
}

func MustClear(r *http.Request) {
	var sessionID = r.URL.Query().Get(accessToken)
	var e = session.MarkDelete(sessionID)
	if e != nil {
		glog.Error(e, "remove session")
	}
}
