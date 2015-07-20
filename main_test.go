package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func Test_Upload(t *testing.T) {
	m := web.New()
	route(m)
	ts := httptest.NewServer(m)
	res, err := http.Get(ts.URL + "/face_detect")
	if err != nil {
		t.Error("unexpected")
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Error("status code not mutch; expected: %s,actual: %s", http.StatusOK, res.StatusCode)
	}
}
