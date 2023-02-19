package app

import (
	"net/http"
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
