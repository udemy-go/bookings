package main

import (
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var mh myHandler

	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but it is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler

	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but it is %T", v)
	}
}