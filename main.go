package main

import (
	"fmt"
	"net/http"

	"github.com/masu-mi/face_detector/contorollers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func route(m *web.Mux) {
	m.Get("/face_detect/:name", http.StripPrefix("/face_detect/", http.FileServer(http.Dir("./results/"))))

	m.Get("/face_detect", controllers.ControllPannel)
	m.Post("/face_detect", controllers.RegisterFace)
}

func main() {
	route(goji.DefaultMux)
	goji.Serve()

	fmt.Println("hello")
}
