source := $(shell find ./ -name '*.go')

all:: face_detector

face_detector:: $(source) ./config/haarcascade_frontalface_default.xml
	go build
