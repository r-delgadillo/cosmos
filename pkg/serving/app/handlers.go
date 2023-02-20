package app

import (
	"net/http"

	"github.com/r-delgadillo/cosmos/pkg/examples/pipelines"
)

type HealthOk struct {
	Decription string
}

func (s *Server) handleHealth() Handler {
	return func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return &Response{
			StatusCode: http.StatusOK,
			Data: HealthOk{
				Decription: "OK",
			},
		}, nil
	}
}

func (s *Server) handlePipelinesSquaring() Handler {
	return func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		pipelines.SquarePipelines([]int{1, 2, 3})
		return &Response{
			StatusCode: http.StatusOK,
			Data: HealthOk{
				Decription: "OK",
			},
		}, nil
	}
}
