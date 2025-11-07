package service

import (
	"errors"
	"time"

	"github.com/zicofachreza/go-urgym-app/user-service/internal/model"
	"github.com/zicofachreza/go-urgym-app/user-service/internal/utils"
	"gorm.io/gorm"
)

func (s *UserService) LoginUser(identifier, password, ip, userAgent string) (map[string]string, error) {
	user, err := s.Repo.FindByEmailOrUsername(identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCodeError("InvalidCredentials")
		}
		return nil, err
	}

	if !utils.ComparePassword(password, user.Password) {
		return nil, utils.NewCodeError("InvalidCredentials")
	}

	// --- Generate Tokens ---
	payload := map[string]any{
		"user_id": user.ID,
		"email":   user.Email,
	}

	accessToken, err := utils.SignToken(payload)
	if err != nil {
		return nil, utils.NewCodeError("AccessTokenError")
	}

	refreshToken, err := utils.SignRefreshToken(payload)
	if err != nil {
		return nil, utils.NewCodeError("RefreshTokenError")
	}

	// --- Save Session ---
	hashed := utils.HashSHA256(refreshToken)

	// Limit max 5 device sessions per user
	if err := s.SessionRepo.LimitDeviceSessions(user.ID); err != nil {
		return nil, err
	}

	session := model.Session{
		UserID:      user.ID,
		HashedToken: hashed,
		DeviceInfo:  userAgent,
		IpAddress:   ip,
		ExpiresAt:   time.Now().Add(7 * 24 * time.Hour), // 7 hari
		LastUsedAt:  time.Now(),
	}

	if err := s.SessionRepo.CreateSession(&session); err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}
