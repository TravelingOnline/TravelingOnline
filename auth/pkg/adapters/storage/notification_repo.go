package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/onlineTraveling/auth/internal/notification/domain"
	notifPort "github.com/onlineTraveling/auth/internal/notification/port"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/mapper"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type notifRepo struct {
	db *gorm.DB
}

func NewNotifRepo(db *gorm.DB) notifPort.Repo {
	return &notifRepo{
		db: db,
	}
}

func (r *notifRepo) SendMessage(ctx context.Context, notif domain.Notification) error {
	o := mapper.NotifDomain2Storage(notif)
	return r.db.Table("inbox").WithContext(ctx).Create(o).Error
}

func (r *notifRepo) GetUnreadMessages(ctx context.Context, userID string) ([]*domain.Notification, error) {
	var notif []types.Notification
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	err = r.db.Table("inbox").WithContext(ctx).Where("user_id = ? AND read = ?", uid, false).Find(&notif).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	r.db.Table("inbox").WithContext(ctx).Where("user_id = ? AND read = ?", uid, false).Updates(map[string]interface{}{"read": true})
	return mapper.NotifStorage2Domain(notif), nil
}
