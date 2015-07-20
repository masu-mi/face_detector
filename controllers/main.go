package controllers

//#cgo pkg-config: opencv
//#include <cv.h>
//#include <highgui.h>
import "C"

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"code.google.com/p/go-uuid/uuid"

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

// RegisterFace : アップロード画像の処理を行う
func RegisterFace(c web.C, w http.ResponseWriter, r *http.Request) {
	uuid := uuid.New()
	file, _, err := r.FormFile("body")
	if err != nil {
		fmt.Print(err.Error())
		w.WriteHeader(500)
	}
	defer file.Close()
	sourcePath := fmt.Sprintf("./tmp/%s", uuid)
	resultPath := fmt.Sprintf("./results/%s.jpg", uuid)
	fmt.Print(w, sourcePath)
	out, err := os.Create(fmt.Sprintf("./tmp/%s", uuid))
	if err != nil {
		fmt.Print(w, "fail to create")
		fmt.Print(err.Error())
		w.WriteHeader(500)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "faile to copy")
		fmt.Fprintln(w, err)
		return
	}
	detect(sourcePath, resultPath)
	http.Redirect(w, r, fmt.Sprintf("/face_detect/%s.jpg", uuid), http.StatusFound)
}

func detect(source, savePath string) {

	tarImg := C.cvLoadImage(
		C.CString(source),
		C.CV_LOAD_IMAGE_ANYDEPTH|C.CV_LOAD_IMAGE_ANYCOLOR,
	)
	defer C.cvReleaseImage(&tarImg)

	cvHCC := (*C.CvHaarClassifierCascade)(
		C.cvLoad(C.CString("./config/haarcascade_frontalface_default.xml"),
			(*C.CvMemStorage)(nil),
			(*C.char)(nil),
			(**C.char)(nil)))
	cvMStr := C.cvCreateMemStorage(0)

	face := C.cvHaarDetectObjects(
		unsafe.Pointer(tarImg),
		cvHCC,
		cvMStr,
		1.11,
		3,
		0,
		C.cvSize(0, 0),
		C.cvSize(0, 0),
	)
	for i := C.int(0); i < face.total; i++ {
		faceRect := (*C.CvRect)(unsafe.Pointer(C.cvGetSeqElem(face, i)))
		C.cvRectangle(
			unsafe.Pointer(tarImg),
			C.cvPoint(faceRect.x, faceRect.y),
			C.cvPoint(faceRect.x+faceRect.width, faceRect.y+faceRect.height),
			C.cvScalar(0, 0, 255, 0),
			3,
			C.CV_AA,
			0,
		)
	}
	C.cvSaveImage(C.CString(savePath), unsafe.Pointer(tarImg), (*C.int)(nil))
}
