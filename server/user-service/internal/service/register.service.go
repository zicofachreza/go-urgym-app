package service

import (
	"strings"

	"github.com/zicofachreza/go-urgym-app/user-service/internal/model"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/utils"
)

func (s *UserService) RegisterUser(user *model.User) error {
	err := s.Repo.RegisterUser(user)
	if err != nil {
		errMsg := strings.ToLower(err.Error())

		if strings.Contains(errMsg, "duplicate") && strings.Contains(errMsg, "username") {
			return utils.NewError("ValidationError", "Username is already used.")
		}

		if strings.Contains(errMsg, "duplicate") && strings.Contains(errMsg, "email") {
			return utils.NewError("ValidationError", "Email is already used.")
		}

		return err
	}
	return nil
}
