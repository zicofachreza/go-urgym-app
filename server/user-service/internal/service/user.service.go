package service

import "github.com/zicofachreza/go-urgym-app/user-service/internal/repository"

type UserService struct {
	Repo        *repository.UserRepository
	SessionRepo *repository.SessionRepository
}
