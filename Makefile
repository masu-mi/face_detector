source := $(shell find ./ -name *.go)

all:: face_detector

face_detector:: $(source)
	go build
