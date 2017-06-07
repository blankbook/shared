package web

import (
    "github.com/gorilla/mux"
)

type HttpRouter struct {
    router mux.Router
}

func NewHttpRouter(router mux.Router) {
   // gotta do stuf here, use mux to set path prefixes
}

func (r HttpRouter) HandleRoute(method int, path string, handler func(http.ResponseWriter, *http.Request)) {
   r.router.... 
}
