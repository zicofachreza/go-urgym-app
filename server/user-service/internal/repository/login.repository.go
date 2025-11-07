package repository

import "github.com/zicofachreza/go-urgym-app/user-service/internal/model"

func (r *UserRepository) FindByEmailOrUsername(identifier string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SessionRepository) CreateSession(session *model.Session) error {
	return r.DB.Create(session).Error
}

// LimitDeviceSessions: hapus session lama jika user punya >5 device
func (r *SessionRepository) LimitDeviceSessions(userID uint) error {
	var sessions []model.Session
	r.DB.Where("user_id = ?", userID).Order("last_used_at desc").Find(&sessions)

	if len(sessions) > 5 {
		oldSessions := sessions[5:]
		var ids []uint
		for _, s := range oldSessions {
			ids = append(ids, s.ID)
		}
		return r.DB.Where("id IN ?", ids).Delete(&model.Session{}).Error
	}
	return nil
}
