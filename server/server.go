package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RedeployAB/gpip/config"
)

// Server represents a http server with router (ServeMux)
// and configuration.
type Server struct {
	httpServer *http.Server
	router     *http.ServeMux
}

// New returns a configured Server object.
func New(conf config.Configuration, r *http.ServeMux) *Server {
	srv := &http.Server{
		Addr:         conf.Host + ":" + conf.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	return &Server{httpServer: srv, router: r}
}

// Start runs ListenAndServe in a go routine and
// runs function shutdown which blocks.
func (s *Server) Start() {
	s.routes()
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server err: %v/n", err)
		}
	}()
	log.Printf("server listening on: %s", s.httpServer.Addr)
	s.shutdown()
}

// shutdown will await signal interrupt and kill and
// block until either of those are received.
func (s *Server) shutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, os.Kill)
	sig := <-stop
	log.Printf("shutting down server. reason: %s\n", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Turn of SetKeepAlive when awaiting shutdown.
	s.httpServer.SetKeepAlivesEnabled(false)
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown server gracefully: %v", err)
	}
}
