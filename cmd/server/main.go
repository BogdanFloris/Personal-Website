package main

import (
	"bogdanfloris-com/internal/common/config"
	"bogdanfloris-com/internal/common/logger"
	"bogdanfloris-com/internal/server"
	"flag"
	"github.com/gofiber/fiber/v2"
	"log"
	"path/filepath"
)

var (
	configPath   string
	prefork      bool
	https        bool
	certFilePath string
	keyFilePath  string
)

func init() {
	flag.StringVar(&configPath, "config", "./config/.env", "The application .env file.")
	flag.BoolVar(&prefork, "prefork", false, "Runs the app the prefork mode.")
	flag.BoolVar(&https, "https", false, "Runs the app in https mode. (needs cert file and key file).")
	flag.StringVar(&certFilePath, "certFilePath", "./config/server.crt", "Public key file path.")
	flag.StringVar(&keyFilePath, "keyFilePath", "./config/server.key", "Private key file path.")
}

func main() {
	flag.Parse()

	configPath, err := filepath.Abs(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	sugaredLogger := logger.NewLogger(
		config.GetBool("DEBUG", false),
		config.GetString("BF_LOG_PATH", ""))

	app := fiber.New(fiber.Config{Prefork: prefork})
	fiberServer := server.New(app, sugaredLogger, &server.Config{
		AppName: config.GetString("BF_APP_NAME", "bogdanfloris-com"),
		Host:    config.GetString("BF_APP_HOST", "127.0.0.1"),
		Port:    config.GetString("BF_APP_PORT", "5000"),
	})

	if https {
		fiberServer.ListenTLS(certFilePath, keyFilePath)
	} else {
		fiberServer.Listen()
	}
}
