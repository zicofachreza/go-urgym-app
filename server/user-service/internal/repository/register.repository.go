package repository

import "github.com/zicofachreza/go-urgym-app/user-service/internal/model"

func (r *UserRepository) RegisterUser(user *model.User) error {
	return r.DB.Create(user).Error
}
