package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/masu-mi/face_detector/controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var toolURI = regexp.MustCompile(`/face_detect/?`)

func renameID(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = fmt.Sprintf("%s.jpg", r.URL.Path)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func route(m *web.Mux) {

	resultMux := web.New()
	resultMux.Get("/face_detect/:name", http.StripPrefix("/face_detect/", http.FileServer(http.Dir("./results/"))))
	resultMux.Use(renameID)

	m.Handle("/face_detect/:name", resultMux)
	m.Get(toolURI, controllers.ControllPannel)
	m.Post(toolURI, controllers.RegisterFace)
}

func main() {
	route(goji.DefaultMux)
	goji.Serve()

	fmt.Println("hello")
}
