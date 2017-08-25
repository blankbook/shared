package web

import (
    "net/http"
    "database/sql"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

const MissingParamErr = "Missing param %v"

// Router is an interface used for all incoming and outgoing network requests
// in the services
type Router interface {
    HandleRoute(methods []string, path string, reqParams []string,
                optParams []string,
                handler func(w http.ResponseWriter,
                             queryParams map[string][]string,
                             body string, db *sql.DB),
                db *sql.DB)
}
