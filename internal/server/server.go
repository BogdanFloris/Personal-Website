package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	logger *zap.SugaredLogger
	config *Config
}

func New(app *fiber.App, logger *zap.SugaredLogger, config *Config) *Server {
	return &Server{
		app:    app,
		logger: logger,
		config: config,
	}
}

func (s *Server) Listen() {
	if !fiber.IsChild() {
		s.logger.Infof("Starting up %s\n", s.config.AppName)
	}
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	if err := s.app.Listen(address); err != nil {
		s.logger.Error(err)
	}
}

func (s *Server) ListenTLS(certFilePath string, keyFilePath string) {
	if !fiber.IsChild() {
		s.logger.Infof("Starting up %s\n", s.config.AppName)
	}
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	if err := s.app.ListenTLS(address, certFilePath, keyFilePath); err != nil {
		s.logger.Error(err)
	}
}
