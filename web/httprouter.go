package web

import (
    "github.com/gorilla/mux"
    "net/http"
)

type HttpRouter struct {
    router *mux.Router
}

func NewHttpRouter(router *mux.Router) *HttpRouter {
    return &HttpRouter{router}
}

func (r *HttpRouter) HandleRoute(method int, path string,
                                handler func(http.ResponseWriter,
                                             *http.Request)) {

   r.router.HandleFunc(path, handler).Methods(path).Schemes("http")
}
