package web

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPRouter is a structure used for all incoming and outgoing HTTP
// requests/calls in the services
type HTTPRouter struct {
	router     *mux.Router
	pathPrefix string
}

// NewHTTPRouter constructs a new HTTPRouter structure
func NewHTTPRouter(router *mux.Router, pathPrefix string) *HTTPRouter {
	return &HTTPRouter{router, pathPrefix}
}

// HandleRoute adds a new route to the HTTPRouter
func (r *HTTPRouter) HandleRoute(methods []string, path string,
	handler func(w http.ResponseWriter,
		queryParams map[string][]string,
		body string)) {

	handlerWrapper := func(w http.ResponseWriter, r *http.Request) {
		var b []byte
		var err error
		if r.Body != nil {
			b, err = ioutil.ReadAll(r.Body)
		}
		queryParams := r.URL.Query()
		body := string(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handler(w, queryParams, body)
	}

	r.router.HandleFunc(path, handlerWrapper).Methods(methods...)
}

// StartListening starts redirecting all requests through the HTTPRouter
func (r *HTTPRouter) StartListening() {
	http.Handle("/", http.StripPrefix(r.pathPrefix, r.router))
}
