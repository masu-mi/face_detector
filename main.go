package main

import (
	"fmt"

	"github.com/masu-mi/face_detector/contorollers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func route(m *web.Mux) {
	m.Get("/face_detect", controllers.ControllPannel)
}

func main() {
	route(goji.DefaultMux)
	goji.Serve()

	fmt.Println("hello")
}
