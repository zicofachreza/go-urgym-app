package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/utils"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // jalankan handler dulu

	// Ambil semua error dari context Gin
	errs := c.Errors
	if len(errs) == 0 {
		return
	}

	err := errs.Last().Err

	// Tampilkan log kalau bukan production
	if os.Getenv("NODE_ENV") != "production" {
		log.Println("Error:", err)
	}

	switch e := err.(type) {
	case *utils.CustomError:
		switch e.Name {
		case "ValidationError":
			utils.JSONError(c, http.StatusBadRequest, "error", e.Message)
			return

		case "InvalidCredentials":
			utils.JSONError(c, http.StatusUnauthorized, "error", "Email/Username or Password is incorrect.")
			return

		case "InvalidAccessToken", "InvalidRefreshToken", "JsonWebTokenError":
			utils.JSONError(c, http.StatusUnauthorized, "error", "Invalid or expired token. Please sign in again.")
			return

		case "AccessTokenExpired", "RefreshTokenExpired":
			utils.JSONError(c, http.StatusUnauthorized, "error", "Your session has expired. Please sign in again.")
			return

		case "AccessTokenError":
			utils.JSONError(c, http.StatusUnauthorized, "error", "Failed to generate access token.")
			return

		case "RefreshTokenError":
			utils.JSONError(c, http.StatusUnauthorized, "error", "Failed to generate refresh token.")
			return

		case "NotFound":
			utils.JSONError(c, http.StatusNotFound, "error", "User not found.")
			return
		}

	default:
		utils.JSONError(c, http.StatusInternalServerError, "error", "Internal Server Error.")
	}
}
