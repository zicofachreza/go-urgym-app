package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/handler"
)

func SetupRoutes(r *gin.Engine, h *handler.UserHandler) {
	// Auth routes
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	// Contoh jika nanti ada route grup:
	// api := r.Group("/api")
	// {
	// 	api.GET("/profile", h.GetProfile)
	// 	api.POST("/logout", h.Logout)
	// }
}
