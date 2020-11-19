package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Server multipart/form server
type Server struct {
}

// ServeHTTP impl http.Handler
// simple upload server. return file content.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, r.RemoteAddr)
	// max memory size 64M
	if err := r.ParseMultipartForm(64 << 20); err != nil {
		w.Write([]byte(fmt.Sprintf("parse multipart/form data failed: %+v", err)))
		return
	}
	defer r.MultipartForm.RemoveAll()

	for _, fs := range r.MultipartForm.File {
		for _, f := range fs {
			file, err := f.Open()
			if err != nil {
				w.Write([]byte(fmt.Sprintf("open file '%s' failed: %+v", f.Filename, err)))
				return
			}
			b, err := ioutil.ReadAll(file)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("read file '%s' failed: %+v", f.Filename, err)))
			}
			w.Write([]byte(f.Filename + ":"))
			w.Write(b)
		}
	}
	return
}
