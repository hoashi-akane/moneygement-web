package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "localhost:8080/?msg=sweets", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK { t.Fatal("unexcepted status code")}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexcepted error")
	}
	const excepted = "<html><body>sweets</body></html>"
	if s := string(b); s!= excepted { t.Fatalf("unexcepted response: %s", s)}
}