package transport

import (
	"context"
	"log"
	"net/http"
)

type Server struct{
	srv *http.Server
}

func NewHttpServer(addr string,handler http.Handler) *Server {
	return &Server{srv: &http.Server{
		Addr:              addr,
		Handler:           handler,
	}}
}

func (s *Server) Starrt() error{
	log.Printf("[HTTP] Listening on: %s\n", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Printf("[HTTP] Shutdown on: %s\n", s.srv.Addr)
	return s.srv.Shutdown(ctx)
}
