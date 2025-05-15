package server

import (
	"github.com/go-chi/chi/v5"

	"hackathons/internal/handlers/hackathons"
	"hackathons/internal/handlers/users"
)

func (s *Server) initRouter() {
	s.router = chi.NewRouter()

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/hackathons", s.registerHackathonsRoutes)
		r.Route("/users", s.registerUsersRoutes)
	})
}

func (s *Server) registerHackathonsRoutes(r chi.Router) {
	r.Get("/", hackathons.GetHackathons(s.hackathonsService))
}

func (s *Server) registerUsersRoutes(r chi.Router) {
	r.Get("/me", users.GetMe(s.usersService))
	r.Get("/participants", users.GetParticipants(s.usersService))
}
