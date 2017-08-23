package video

import (
	"net/http"
	"tvac/web"
)

type VideoServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewVideoServer() *VideoServer {
	s := &VideoServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/update", s.HandleUpdate)
	return s

}
func (s *VideoServer) HandleCreate(w http.ResponseWriter, r *http.Request) {

}
func (s *VideoServer) HandleDelete(w http.ResponseWriter, r *http.Request) {

}
func (s *VideoServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {

}
