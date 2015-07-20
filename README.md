# face-detector

サンプルコード: making face detection Web API using Goji(golang light weight web application frame work).

Cascade Classifier for Object Detection用の学習データファイル(haarcascade_frontalface_default.xml)が必要。

以下の様な内容から始まるhaarcascade_frontalface_default.xml を探してくる。
公式github から curl したxml だとver.の都合で動かなかった(自分のOpenCVは`v2.11.1`)上、どこで見つけてきたか忘れたうえ、
OpenCVは真面目に調べていない。

```xml
<opencv_storage>
<haarcascade_frontalface_default type_id="opencv-haar-classifier">
  <size>24 24</size>
  <stages>
```

## reference(I studied following sites.)

- [goji](https://goji.io/)
- [gojiのMiddlewareの使い方](http://qiita.com/reiki4040/items/a038f1b99e0caee97d3e)
- [goji sample code](https://github.com/haruyama/golang-goji-sample)
- [net/httptestでgojiのテストをする](http://qiita.com/r_rudi/items/727fb85030e91101801d)
- [html/template](http://golang.org/pkg/html/template/)
- [cgo](https://github.com/golang/go/wiki/cgo)
- [OpenCVで顔検出](http://www.aianet.ne.jp/~asada/prog_doc/opencv/opencv_obj_det_img.htm)
- [objdetect_cascade_classification](http://opencv.jp/opencv-2.2/c/objdetect_cascade_classification.html)
- [reading_and_writing_images_and_video](http://opencv.jp/opencv-2.1/cpp/reading_and_writing_images_and_video.html)
