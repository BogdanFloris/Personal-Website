package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	config *Config
}

func New(app *fiber.App, config *Config) *Server {
	return &Server{
		app:    app,
		config: config,
	}
}

func (s *Server) Listen() {
	if !fiber.IsChild() {
		fmt.Printf("Starting up %s\n", s.config.AppName)
	}
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	if err := s.app.Listen(address); err != nil {
		panic(err)
	}
}

func (s *Server) ListenTLS(certFilePath string, keyFilePath string) {
	if !fiber.IsChild() {
		fmt.Printf("Starting up %s\n", s.config.AppName)
	}
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	if err := s.app.ListenTLS(address, certFilePath, keyFilePath); err != nil {
		panic(err)
	}
}
