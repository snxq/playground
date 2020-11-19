package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Server multipart/form server
type Server struct {
	w http.ResponseWriter
}

// ServeHTTP impl http.Handler
// simple upload server. return file content.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.w = w
	fmt.Println(r.URL, r.RemoteAddr)
	// max memory size 64M
	if err := r.ParseMultipartForm(64 << 20); err != nil {
		s.w.Write([]byte(fmt.Sprintf("parse multipart/form data failed: %+v", err)))
		return
	}
	defer r.MultipartForm.RemoveAll()

	for _, fs := range r.MultipartForm.File {
		for _, f := range fs {
			file, err := f.Open()
			if err != nil {
				s.w.Write([]byte(fmt.Sprintf("open file '%s' failed: %+v", f.Filename, err)))
				return
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				s.w.Write([]byte(fmt.Sprintf("read file '%s' failed: %+v", f.Filename, err)))
				return
			}
			s.w.Write([]byte(f.Filename + ":"))
			s.w.Write(b)
		}
	}
	return
}

func (s *Server) Write(p []byte) (n int, err error) {
	n, err = s.w.Write(p)
	if err != nil {
		fmt.Printf("Write response failed: %+v", err)
	}
	return n, err
}
