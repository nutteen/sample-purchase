package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nutteen/png-core/core/db"
	"github.com/nutteen/png-core/core/logger"
	middlewarelogger "github.com/nutteen/png-core/core/middlewares/logger"
	"github.com/nutteen/png-core/core/middlewares/usercontext"
	"github.com/nutteen/png-core/core/validator"
	"github.com/nutteen/sample-purchase/pkg/config"
	"log"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error loading config")
	}

	logger.InitializeLogger(logger.LoggerConfig{
		IsProductionMode: true,
	})

	validator.InitializeValidator()

	app := fiber.New()
	dbInstance := db.New(config.AppConfig.Database)

	// Setup middlewares
	app.Use(requestid.New())
	app.Use(usercontext.New())
	app.Use(middlewarelogger.NewLogger(logger.Log, middlewarelogger.ConfigDefault))

	// todo: Setup  service
	// todo: Registers routes

	app.Listen(config.AppConfig.Server.Port)
}