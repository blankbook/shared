package web

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

type HttpRouter struct {
    router *mux.Router
    pathPrefix string
}

func NewHttpRouter(router *mux.Router, pathPrefix string) *HttpRouter {
    return &HttpRouter{router, pathPrefix}
}

func (r *HttpRouter) HandleRoute(method int, path string,
                                handler func(http.ResponseWriter,
                                             *http.Request)) {
    var methodStr string
    switch method {
    case GET:
        methodStr = "GET"
    case POST:
        methodStr = "POST"
    case PUT:
        methodStr = "PUT"
    case DELETE:
        methodStr = "DELETE"
    }
    r.router.HandleFunc(path, handler).Methods(methodStr)
}

func (r *HttpRouter) StartListening() {
    http.Handle("/", http.StripPrefix(r.pathPrefix, r.router))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Testing!!")
}
