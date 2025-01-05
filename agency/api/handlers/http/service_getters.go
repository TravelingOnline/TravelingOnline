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

func tourServiceGetter(appContainer app.App) ServiceGetter[*service.TourService] {
	return func(ctx context.Context) *service.TourService {
		return service.NewTourService(appContainer.TourService(ctx))
	}
}
