package main

import (
	"net/http"
)

type (
	path        string
	handlerfunc func(http.ResponseWriter, *http.Request)
	route       func(path, handlerfunc) *routeHandler
)

type routeHandler struct {
	path
	handlerfunc
	mux *http.ServeMux
	pr  string
}

func (s *serverhandler) newRouter(parentRoute string) route {

	return func(p path, hf handlerfunc) *routeHandler {

		return &routeHandler{
			path:        p,
			handlerfunc: hf,
			mux:         s.mux,
			pr:          parentRoute,
		}

	}
}

func (rh *routeHandler) GET() {
	rh.mux.HandleFunc(string(path(rh.pr)+rh.path), rh.handlerfunc)
}

func (rh *routeHandler) POST() {
	rh.mux.HandleFunc(string(path(rh.pr)+rh.path), rh.handlerfunc)
}

func (rh *routeHandler) TEST() {

	rh.mux.HandleFunc(string(path(rh.pr)+rh.path), rh.handlerfunc)

}
