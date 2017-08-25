package web

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "database/sql"

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
                                 reqParams []string, optParams []string,
                                 handler func(w http.ResponseWriter,
                                              queryParams map[string][]string,
                                              body string, db *sql.DB),
                                 db *sql.DB) {
    handlerWrapper := func(w http.ResponseWriter, r *http.Request) {
        var b []byte
        if r.Body != nil {
            var err error
            b, err = ioutil.ReadAll(r.Body)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
        }
        filteredParams := make(map[string][]string)
        queryParams := r.URL.Query()
        for _, param := range reqParams {
            if val, ok := queryParams[param]; ok && len(val) > 0 {
                filteredParams[param] = val
            } else {
                http.Error(w, fmt.Sprintf(MissingParamErr, param), http.StatusBadRequest)
                return
            }
        }
        for _, param := range optParams {
            if val, ok := queryParams[param]; ok && len(val) > 0 {
                filteredParams[param] = val
            }
        }
        handler(w, filteredParams, string(b), db)
    }
    r.router.HandleFunc(path, handlerWrapper).Methods(methods...)
}

// StartListening starts redirecting all requests through the HTTPRouter
func (r *HTTPRouter) StartListening() {
    http.Handle("/", http.StripPrefix(r.pathPrefix, r.router))
}
