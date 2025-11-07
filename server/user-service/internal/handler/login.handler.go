package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/dto"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/utils"
)

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()

	tokens, err := h.Service.LoginUser(req.EmailOrUsername, req.Password, ip, userAgent)
	if err != nil {
		c.Error(err) // ✅ langsung lempar ke middleware
		return
	}

	utils.JSONResponse(c, http.StatusOK, "success", "Login successful.", tokens)
}
