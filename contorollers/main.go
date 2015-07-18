package controllers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zenazn/goji/web"
)

type application struct {
	Template *template.Template
}

var app application

func init() {
	// template 一括読み込み
	var templates []string
	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}
	err := filepath.Walk("./views", fn)
	if err != nil {
		panic(err)
	}

	app = application{
		Template: template.Must(template.ParseFiles(templates...)),
	}
	// 各種オプションでdomain, port, を指定できる様にする
}

// ControllPannel : 画像アップロード用コンパネ
func ControllPannel(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	app.Template.ExecuteTemplate(w, "ControllPannel", nil)
}
