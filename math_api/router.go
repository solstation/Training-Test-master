package main

import "net/http"

// Router will storage all the routes with its respective
// handler by its route and its http method, example:
// /api/2.0/login - GET - handler
type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

// NewRouter is constructor for Router struct
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

// ServeHTTP will serve the handlers using the handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, pathExists, methodExists := r.GetHandler(req.URL.Path, req.Method)

	if !pathExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, req)
}

// GetHandler will find a handler by its route and method
func (r *Router) GetHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, pathExists := r.routes[path]
	handler, methodExists := r.routes[path][method]

	return handler, pathExists, methodExists
}