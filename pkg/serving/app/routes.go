package app

import (
	"net/http"
)

// routes defines and attaches the full set of routes in the application to the
// server's router. This should only be called once on server initialization
func (s *Server) routes() {
	// Discovery
	s.Route(
		"health",
		"/health",
		http.MethodGet,
		s.handleHealth,
	)

	// s.Router.NotFoundHandler = s.internal.Router.NewRoute().HandlerFunc(http.NotFound).GetHandler()
}

// Route is helper function to register routes
func (s *Server) Route(
	operationName, path, method string,
	f func() Handler,
	opts ...Option,
) {
	s.Router.HandleFunc(path, s.handle(f(), operationName, opts...)).Methods(method)
}
