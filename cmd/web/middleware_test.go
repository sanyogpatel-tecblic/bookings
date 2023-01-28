package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is not handle,but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler

	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is not HTTP handler but is %T", v))
	}
}
