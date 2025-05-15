package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"hackathons/config"
	postgresAdapter "hackathons/internal/adapter/postgres"
	"hackathons/internal/adapter/repository/hackathons"
	"hackathons/internal/infrastructure/database"
	"hackathons/internal/infrastructure/database/postgres"
	"hackathons/internal/ports/repository"

	"hackathons/internal/services"
	hackathonsServices "hackathons/internal/services/hackathons"
)

type Server struct {
	cfg *config.Config

	hackathonsDB *postgres.Postgres

	// repositories
	usersRepository      repository.UsersRepository
	hackathonsRepository repository.HackathonsRepository

	// services
	usersService      services.Users
	hackathonsService services.Hackathons

	router *chi.Mux
	server *http.Server
}

func New(cfg *config.Config) (*Server, error) {
	s := &Server{
		cfg: cfg,
	}

	if err := s.init(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) init() error {
	if err := s.initDB(); err != nil {
		return fmt.Errorf("init db: %v", err)
	}
	if err := database.MigrateHackathonsDB(s.hackathonsDB); err != nil {
		return fmt.Errorf("migrate static db: %v", err)
	}

	s.initRepositories()
	s.initUseCases()
	s.initRouter()
	s.initHTTPServer()

	return nil
}

func (s *Server) initDB() error {
	var err error

	s.hackathonsDB, err = postgresAdapter.Connect(s.cfg.Server.StaticData.Connection)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) initRepositories() {
	s.usersRepository = hackathons.NewUsersRepository(s.hackathonsDB)
	s.hackathonsRepository = hackathons.NewHackathonsRepository(s.hackathonsDB)
}

func (s *Server) initUseCases() {
	s.usersService = hackathonsServices.NewUsersService(s.usersRepository)
	s.hackathonsService = hackathonsServices.NewHackathonsService(s.hackathonsRepository)
}

func (s *Server) initHTTPServer() {
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.cfg.Server.Addr, s.cfg.Server.Port),
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (s *Server) Run() {
	log.Println("Server started")

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
