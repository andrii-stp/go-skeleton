package internal

import (
	"context"

	"github.com/andrii-stp/go-skeleton/config"
)

type Server struct {
	cfg *config.Server
}

func NewServer(cfg *config.Server) *Server {
	return &Server{cfg}
}

func (s *Server) Start(ctx context.Context) {

}

func (s *Server) Stop(gracefully <-chan bool) {

}
