package api

import (
	"net/http"
	"tva_api/api/auth"
	"tva_api/api/upload"
	"tva_api/web"
)

type ApiServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewApiServer() *ApiServer {

	var s = &ApiServer{
		ServeMux: http.NewServeMux(),
	}
	s.Handle("/auth/", http.StripPrefix("/auth", auth.NewAuthServer()))
	s.Handle("/upload/", http.StripPrefix("/upload", upload.NewUploadServer()))
	return s
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.Recover(w)
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
	header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	header.Add(
		"Access-Control-Max-Age",
		"2520000", // 30 days
	)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}
