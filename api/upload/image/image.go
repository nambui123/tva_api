package image

import (
	"net/http"
	"tva_api/web"
)

type ImageServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewImageServer() *ImageServer {
	s := &ImageServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/update", s.HandleUpdate)
	return s

}
func (s *ImageServer) HandleCreate(w http.ResponseWriter, r *http.Request) {

}
func (s *ImageServer) HandleDelete(w http.ResponseWriter, r *http.Request) {

}
func (s *ImageServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {

}
