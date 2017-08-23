package upload

import (
	"net/http"
	"tvac/api/upload/image"
	"tvac/api/upload/video"
	"tvac/web"
)

type UploadServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewUploadServer() *UploadServer {

	var s = &UploadServer{
		ServeMux: http.NewServeMux(),
	}
	s.Handle("/video/", http.StripPrefix("/video", video.NewVideoServer()))
	s.Handle("/image/", http.StripPrefix("/image", image.NewImageServer()))
	return s
}
