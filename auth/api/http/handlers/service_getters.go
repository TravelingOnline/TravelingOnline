package handlers

import (
	"context"

	"github.com/onlineTraveling/auth/api/service"
	"github.com/onlineTraveling/auth/app"
	"github.com/onlineTraveling/auth/config"
)

// user service transient instance handler
func UserServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.UserService] {
	return func(ctx context.Context) *service.UserService {
		return service.NewUserService(appContainer.UserService(ctx),
			cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute, appContainer.CodeVerificationService(ctx))
	}
}

func NotificationServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.NotificationService] {
	return func(ctx context.Context) *service.NotificationService {
		return service.NewNotificationSerivce(appContainer.NotifService(ctx), cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	}
}
