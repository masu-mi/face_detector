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

func route(m *web.Mux) {
	m.Get("/face_detect/:name", http.StripPrefix("/face_detect/", http.FileServer(http.Dir("./results/"))))

	m.Get(toolURI, controllers.ControllPannel)
	m.Post(toolURI, controllers.RegisterFace)
}

func main() {
	route(goji.DefaultMux)
	goji.Serve()

	fmt.Println("hello")
}
