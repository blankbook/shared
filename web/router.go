package web

import (
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// Router is an interface used for all incoming and outgoing network requests
// in the services
type Router interface {
	HandleRoute(methods []string, path string,
		handler func(w http.ResponseWriter,
			queryParams map[string][]string,
			body string))
}
