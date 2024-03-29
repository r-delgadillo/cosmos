package app

import (
	"net/http"
)

// routes defines and attaches the full set of routes in the application to the
// server's router. This should only be called once on server initialization
func (s *Server) routes() {

	s.Route(
		"health",
		"/health",
		http.MethodGet,
		s.handleHealth,
	)

	s.Route(
		"squaringpipelines",
		"/examples/pipelines/squaring",
		http.MethodGet,
		s.handlePipelinesSquaring,
	)

	s.Route(
		"squaringpipelines",
		"/examples/pipelines/producer",
		http.MethodGet,
		s.handleProducer,
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
