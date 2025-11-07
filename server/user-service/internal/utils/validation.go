package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(c *gin.Context, err error) {
	var fieldErrMsg string

	if errs, ok := err.(validator.ValidationErrors); ok && len(errs) > 0 {
		fieldErr := errs[0] // ambil error pertama saja

		switch fieldErr.Field() {
		case "EmailOrUsername":
			fieldErrMsg = "Email/Username is required."

		case "Email":
			switch fieldErr.Tag() {
			case "required":
				fieldErrMsg = "Email is required."
			case "email":
				fieldErrMsg = "Invalid email format."
			default:
				fieldErrMsg = "Invalid email field."
			}

		case "Username":
			switch fieldErr.Tag() {
			case "required":
				fieldErrMsg = "Username is required."
			case "min":
				fieldErrMsg = "Username must be at least 5 characters long."
			default:
				fieldErrMsg = "Invalid username field."
			}

		case "Password":
			switch fieldErr.Tag() {
			case "required":
				fieldErrMsg = "Password is required."
			case "min":
				fieldErrMsg = "Password must be at least 5 characters long."
			default:
				fieldErrMsg = "Invalid password field."
			}

		default:
			fieldErrMsg = "Invalid request body."
		}
	} else {
		fieldErrMsg = "Invalid request body."
	}

	c.Error(&CustomError{
		Name:    "ValidationError",
		Message: fieldErrMsg,
	})
}
