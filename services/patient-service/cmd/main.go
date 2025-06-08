// Entry point aplikasi
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"patient-service/internal/config"
	"patient-service/internal/database"
	"patient-service/internal/handler"
	"patient-service/internal/middleware"
	"patient-service/internal/repository"
	"patient-service/internal/service"
	"patient-service/pkg/validator"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize validator
	validate := validator.New()

	// Initialize repositories
	patientRepo := repository.NewPatientRepository(db)

	// Initialize services
	patientService := service.NewPatientService(patientRepo)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())
	app.Use(middleware.Metrics())

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "patient-service",
			"version": cfg.App.Version,
		})
	})

	// API routes
	api := app.Group("/api/v1")

	// Public routes
	api.Get("/patients/:id/public", handler.GetPatientPublicInfo(patientService))

	// Protected routes
	protected := api.Group("/", middleware.JWTAuth(cfg.JWT.Secret))

	// Patient routes
	patientHandler := handler.NewPatientHandler(patientService, validate)
	protected.Post("/patients", patientHandler.CreatePatient)
	protected.Get("/patients/:id", patientHandler.GetPatient)
	protected.Put("/patients/:id", patientHandler.UpdatePatient)
	protected.Delete("/patients/:id", patientHandler.DeletePatient)
	protected.Get("/patients", patientHandler.ListPatients)

	// Metrics endpoint (untuk Prometheus)
	app.Get("/metrics", middleware.PrometheusHandler())

	// Graceful shutdown
	go func() {
		if err := app.Listen(":" + cfg.App.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"error": fiber.Map{
			"message": message,
			"code":    code,
		},
	})
}
