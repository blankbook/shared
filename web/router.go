package web

import (
    "net/http"
)

const (
    GET = iota
    POST = iota
    PUT = iota
    DELETE = iota
)

type Router interface {
    HandleRoute(method int, path string, handler func(http.ResponseWriter, *http.Request))
}
