package main

import (
	"net/http"
	"os"
	"testing"
)

func Testmain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
