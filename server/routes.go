package server

import (
	"github.com/RedeployAB/gpip/middleware"
)

// routes initializes routes for Server struct.
func (s *Server) routes() {
	s.router.Handle("/", middleware.Logger(s.getIP()))
}
