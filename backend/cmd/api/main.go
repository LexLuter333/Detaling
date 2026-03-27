package main

import (
	"deteleng-backend/internal/config"
	"deteleng-backend/internal/database"
	"deteleng-backend/internal/handlers"
	"deteleng-backend/internal/middleware"
	"deteleng-backend/internal/repository"
	"deteleng-backend/internal/services"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	err := database.InitDB()
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Run migrations
	err = runMigrations()
	if err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	// Seed default data
	err = seedData()
	if err != nil {
		log.Fatalf("❌ Failed to seed data: %v", err)
	}

	// Start background job for deleting old bookings
	go func() {
		ticker := time.NewTicker(24 * time.Hour) // Run once per day
		defer ticker.Stop()
		for range ticker.C {
			repo := repository.NewDatabaseRepository()
			err := repo.DeleteOldCompletedBookings()
			if err != nil {
				log.Printf("⚠️  Error deleting old bookings: %v", err)
			} else {
				log.Println("✅ Old completed bookings deleted successfully")
			}
		}
	}()

	// Initialize repository
	repo := repository.NewDatabaseRepository()

	// Initialize services
	bookingService := services.NewBookingService(repo)
	authService := services.NewAuthService(repo)
	adminService := services.NewAdminService(repo)
	serviceService := services.NewServiceService(repo)
	reviewService := services.NewReviewService(repo)

	// Initialize handlers
	bookingHandler := handlers.NewBookingHandler(bookingService)
	authHandler := handlers.NewAuthHandler(authService)
	adminHandler := handlers.NewAdminHandler(adminService, authService)
	serviceHandler := handlers.NewServiceHandler(serviceService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	// Create Gin router
	r := gin.Default()

	// CORS configuration with configurable origins
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSAllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// API routes
	api := r.Group("/api")
	{
		// Public routes
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		// Services routes (public)
		services := api.Group("/services")
		{
			services.GET("/", serviceHandler.GetPublicServices)
		}

		// Reviews routes (public)
		reviews := api.Group("/reviews")
		{
			reviews.GET("/", reviewHandler.GetPublicReviews)
		}

		// Booking routes
		bookings := api.Group("/bookings")
		{
			bookings.POST("/", bookingHandler.CreateBooking)
			bookings.GET("/", bookingHandler.GetAllBookings)
			bookings.GET("/:id", bookingHandler.GetBooking)
		}

		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// Admin routes (protected)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		{
			admin.GET("/dashboard", adminHandler.Dashboard)
			admin.GET("/bookings", adminHandler.GetAllBookings)
			admin.PUT("/bookings/:id/status", adminHandler.UpdateBookingStatus)
			admin.DELETE("/bookings/:id", adminHandler.DeleteBooking)
			admin.GET("/stats", adminHandler.GetStats)

			// Admin services management
			admin.GET("/services", serviceHandler.GetAllServices)
			admin.POST("/services", serviceHandler.CreateService)
			admin.PUT("/services/:id", serviceHandler.UpdateService)
			admin.DELETE("/services/:id", serviceHandler.DeleteService)

			// Admin reviews management
			admin.GET("/reviews", reviewHandler.GetAllReviews)
			admin.POST("/reviews", reviewHandler.CreateReview)
			admin.PUT("/reviews/:id", reviewHandler.UpdateReview)
			admin.DELETE("/reviews/:id", reviewHandler.DeleteReview)
			admin.POST("/reviews/parse", reviewHandler.ParseReviews)
			admin.GET("/reviews/stats", reviewHandler.GetReviewStats)

			// Admin review sources management
			admin.GET("/review-sources", reviewHandler.GetReviewSources)
			admin.POST("/review-sources", reviewHandler.CreateReviewSource)
			admin.PUT("/review-sources/:id", reviewHandler.UpdateReviewSource)
			admin.DELETE("/review-sources/:id", reviewHandler.DeleteReviewSource)
		}
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

// runMigrations runs SQL migration files
func runMigrations() error {
	log.Println("📦 Running database migrations...")

	migrations := []string{
		"database/migrations/001_create_tables.sql",
		"database/migrations/002_seed_data.sql",
	}

	for _, migration := range migrations {
		log.Printf("📄 Running migration: %s", migration)

		sqlFile, err := os.ReadFile(migration)
		if err != nil {
			return err
		}

		_, err = database.DB.Exec(string(sqlFile))
		if err != nil {
			return err
		}

		log.Printf("✅ Migration completed: %s", migration)
	}

	log.Println("✅ All migrations completed successfully")
	return nil
}

// seedData seeds the database with default data if needed
func seedData() error {
	log.Println("🌱 Checking default data...")

	// Check if admin user exists
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", "admin@deteleng.com").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		log.Println("⚠️  Admin user not found, migrations should have created it")
	}

	// Check if services exist
	err = database.DB.QueryRow("SELECT COUNT(*) FROM services").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		log.Println("⚠️  Services not found, migrations should have created them")
	}

	log.Println("✅ Default data check completed")
	return nil
}
