package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/config"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/handler"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/middleware"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/model"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/repository"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/router"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/service"
)

func main() {
	env := os.Getenv("NODE_ENV")
	if env == "" {
		env = "development"
	}

	port := getPort(env)
	log.Printf("🚀 Starting user-service in %s mode on port %s\n", env, port)

	db := config.ConnectDB()

	if err := db.AutoMigrate(&model.User{}, &model.Session{}); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	repo := &repository.UserRepository{DB: db}
	sessionRepo := &repository.SessionRepository{DB: db}
	svc := &service.UserService{Repo: repo, SessionRepo: sessionRepo}
	h := &handler.UserHandler{Service: svc}

	// === Setup Gin ===
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler)

	// === Setup routes dari package router ===
	router.SetupRoutes(r, h)

	log.Printf("✅ User service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

func getPort(env string) string {
	switch env {
	case "production":
		return "80"
	case "test":
		return "4001"
	default:
		return "3001"
	}
}
