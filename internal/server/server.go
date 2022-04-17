package server

import (
	"github.com/plutonio00/pay-api/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(conf *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + conf.Server.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
