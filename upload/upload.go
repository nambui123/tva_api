package upload

import (
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
	"path"
)

type UploadFileServer struct {
	folder        string
	maxFileSize   int64
	fileServer    http.Handler
	GetHandler    http.Handler
	PostHandler   http.Handler
	DeleteHandler http.Handler
}

func NewUploadFileServer(folder string, maxFileSize int64) *UploadFileServer {
	var s = &UploadFileServer{
		folder:      folder,
		fileServer:  http.FileServer(http.Dir(folder)),
		maxFileSize: maxFileSize,
	}

	s.GetHandler = http.HandlerFunc(s.getFile)
	s.PostHandler = http.HandlerFunc(s.postFile)
	s.DeleteHandler = http.HandlerFunc(s.deleteFile)
	return s
}

func (s *UploadFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		s.PostHandler.ServeHTTP(w, r)
	case http.MethodDelete:
		s.DeleteHandler.ServeHTTP(w, r)
	case http.MethodPut:
	case http.MethodGet:
		s.GetHandler.ServeHTTP(w, r)
	default:
		http.Error(w, "METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}
}

func (s *UploadFileServer) filename(r *http.Request) string {
	return path.Join(s.folder, path.Clean(r.URL.Path), r.FormValue("name"))
}

func (s *UploadFileServer) postFile(w http.ResponseWriter, r *http.Request) {
	filename := s.filename(r)
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte("fail to read file form " + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	outstream, err := os.Create(filename)
	if err != nil {
		glog.Error("create ", filename, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer outstream.Close()
	_, err = io.Copy(outstream, file)
	if err != nil {
		glog.Error("save", filename, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *UploadFileServer) deleteFile(w http.ResponseWriter, r *http.Request) {
	filename := s.filename(r)
	err := os.Remove(filename)
	if err != nil {
		glog.Error("remove", filename, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *UploadFileServer) getFile(w http.ResponseWriter, r *http.Request) {
	s.fileServer.ServeHTTP(w, r)
}
