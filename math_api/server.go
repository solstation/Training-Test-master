package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Server is a struct that will handle everything
// like handlers, routes, middlewares, sockets, etc
type Server struct {
	router *Router
	port   string
}

// NewServer is a constructor for Server struct
func NewServer(port int) *Server {
	return &Server{
		router: NewRouter(),
		port:   ":" + strconv.Itoa(port),
	}
}

// Listen will start the Server
func (s *Server) Listen() error {
	http.Handle("/", s.router)

	return http.ListenAndServe(s.port, nil)
}

// Handle will save a new handler in a certain path and method
func (s *Server) Handle(path, method string, handler http.HandlerFunc) error {
	_, pathExists, methodExists := s.router.GetHandler(path, method)

	if !pathExists {
		s.router.routes[path] = make(map[string]http.HandlerFunc)
	}

	if methodExists {
		return fmt.Errorf("[Server] Method already exists: %s, %s", path, method)
	}

	s.router.routes[path][method] = handler

	return nil
}

// AddMiddleware allows to implement n middlewares to a certain handler
func (s *Server) AddMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}