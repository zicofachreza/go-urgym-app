package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/dto"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/model"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/utils"
)

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.Service.RegisterUser(&user); err != nil {
		c.Error(err) // lempar ke middleware
		return
	}

	userResponse := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	utils.JSONResponse(c, http.StatusCreated, "success", "Register successful.", userResponse)
}
