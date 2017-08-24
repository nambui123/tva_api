package auth

import (
	"net/http"
	"strings"
	"tva_api/api/auth/session"
	"tva_api/o/user"
	"tva_api/web"
)

type AuthServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewAuthServer() *AuthServer {
	var s = &AuthServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/login", s.HandleLogin)
	s.HandleFunc("/register", s.HandleRegister)
	s.HandleFunc("/me", s.handleMe)
	s.HandleFunc("/logout", s.HandleLogout)
	s.HandleFunc("/change_pass", s.handleChangePass)
	return s
}

func (s *AuthServer) handleMe(w http.ResponseWriter, r *http.Request) {
	s.SendJson(w, session.MustGet(r))
}

func (s *AuthServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByUsername(strings.ToLower(body.Username))
	if user.TableUser.IsErrNotFound(err) {
		s.SendError(w, errUserNotFound)
		return
	}
	if err := u.ComparePassword(body.Password); err != nil {
		s.SendError(w, err)
	}

	web.AssertNil(err)

	var ses = session.MustNew(u)
	s.SendData(w, map[string]interface{}{
		"user":    u,
		"session": ses,
	})
}
func (s *AuthServer) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}

	s.MustDecodeBody(r, &body)

	var _, err = user.GetByEmail(strings.ToLower(body.Email))
	if err == nil {
		s.SendError(w, errEmailExites)
		return
	}

	us := &user.User{
		UserName: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}
	err = us.Create()
	if err != nil {
		s.SendError(w, err)
		return
	}

	web.AssertNil(err)
	s.SendData(w, registerSuccess)
}

func (s *AuthServer) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session.MustClear(r)
	s.SendData(w, nil)
}

func (s *AuthServer) handleChangePass(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		OldPass   string `json:"old_pass"`
		NewPass   string `json:"new_pass"`
		ReNewPass string `json:"re_new_pass"`
		Username  string `json:"username"`
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByUsername(strings.ToLower(body.Username))
	if user.TableUser.IsErrNotFound(err) {
		s.SendError(w, errUserNotFound)
		return
	}

	if err := u.ComparePassword(body.OldPass); err != nil {
		s.SendError(w, err)
		return
	}
	u.UpdatePass(body.NewPass)

	s.SendData(w, map[string]interface{}{
		"message": "Update password success",
	})
}
