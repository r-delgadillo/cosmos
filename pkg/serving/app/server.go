package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/r-delgadillo/cosmos/lib/cosmoscontext"
	"github.com/r-delgadillo/cosmos/lib/jsonutil"
)

const (
	defaultAppID = "SERVICE"
	component    = "Server"

	// StandardTimeout is the default timeout for a request
	StandardTimeout = 45 * time.Second
	// ExtendedTimeout is predefined timeout for a request
	ExtendedTimeout = 5 * time.Minute
)

type (
	// Server representation
	Server struct {
		Router *mux.Router
		port   int
	}

	Handler func(w http.ResponseWriter, r *http.Request) (*Response, error)

	// Response encapsulates information about the Response to send back to the
	// client.
	Response struct {
		StatusCode int
		Data       interface{}
	}

	// Option specifies optional data for a request
	Option func(o *options)

	options struct {
		middleware []func() error
		timeout    time.Duration
	}
)

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		port:   8080,
	}
	s.routes()
	return s
}

func (s *Server) handle(
	handler Handler,
	rel string,
	opts ...Option,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o := &options{
			timeout: StandardTimeout,
		}
		for _, opt := range opts {
			opt(o)
		}

		// logger which exists just for this ISR. We create a new logger
		// variable to avoid sharing among different calls to the same API.
		_, cancel := createContextFromRequest(r, rel, o.timeout)
		defer cancel()

		// ISR method steps CV
		statusCode := new(int)

		// recover from panic
		defer func() {
			if err := recover(); err != nil {
				s.handleError(w, errors.New("Error Handle with HTTP request handlers"))
			}
		}()

		for _, m := range o.middleware {
			if err := m(); err != nil {
				s.handleError(w, err)
				return
			}
		}

		response, err := handler(w, r)
		if err != nil {
			s.handleError(w, err)
			return
		}

		if response.StatusCode != 0 {
			*statusCode = response.StatusCode
		} else {
			*statusCode = http.StatusOK
		}

		if response.Data != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(*statusCode)
			if err := jsonutil.NewSafeEncoder(w).Encode(response.Data); err != nil {
				fmt.Println("Error with data")
			}
		} else {
			w.WriteHeader(*statusCode)
			fmt.Fprintf(w, "%s", http.StatusText(*statusCode))
		}
	}
}

func createContextFromRequest(
	r *http.Request,
	operationName string,
	timeout time.Duration,
) (*cosmoscontext.CosmosContext, context.CancelFunc) {
	timeoutCtx, cancel := context.WithTimeout(r.Context(), timeout)

	ctx := cosmoscontext.New(timeoutCtx)

	return ctx, cancel
}

type PlaceHolderError struct{}

func (s *Server) handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(599)

	if encodeErr := jsonutil.NewSafeEncoder(w).Encode(PlaceHolderError{}); encodeErr != nil {
		panic(encodeErr)
	}
}

// Run starts listening on the configured server port.
func (s *Server) Run() error {
	fmt.Println("Starting server...")
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.Router,
	}

	// Watch for the context to cancel and shut down the server. This is done in
	// a separate goroutine because the parent one will block on the call to
	// ListenAndServe().
	// go func() {

	// 	// TODO: configurable timeout?
	// 	closeCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()

	// 	// Shut down the server. This will cause the blocking call of
	// 	// ListenAndServe() below to return with the error http.ErrServerClosed.
	// 	if err := server.Shutdown(closeCtx); err != nil {
	// 		// There's really nothing to be done here because the server is
	// 		// already closing, but we log it in case there's something
	// 		// interesting.
	// 		panic(err)
	// 	}
	// }()

	// Run the server. This will block until the server fails or is shut down.
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
		// return err
	}

	return nil
}
