package http

import (
	"agency/api/service"
	"agency/app"
	"context"
)

func agencyServiceGetter(appContainer app.App) ServiceGetter[*service.AgencyService] {
	return func(ctx context.Context) *service.AgencyService {
		return service.NewAgencyService(appContainer.AgencyService(ctx))
	}
}
