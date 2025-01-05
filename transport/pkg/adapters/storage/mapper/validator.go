package mapper

import (
	"fmt"

	"github.com/onlineTraveling/transport/internal/tour/domain"
	"github.com/onlineTraveling/transport/pkg/adapters/storage/helpers"
)

func TourValidator(tour domain.Tour) error {
	if tour.Id == "" {
		return fmt.Errorf("TOUR ID CANNOT BE EMPTY")
	}

	if !helpers.ValidDate(tour.StartDate) {
		return fmt.Errorf("WRONG STARTDATE FORMAT")
	}

	if !helpers.ValidDate(tour.EndDate) {
		return fmt.Errorf("WRONG ENDDATE FORMAT")
	}

	if tour.Source == "" {
		return fmt.Errorf("SOURCE LOCATION CANNOT BE EMPTY")
	}

	if tour.Destination == "" {
		return fmt.Errorf("DESTINATION LOCATION CANNOT BE EMPTY")
	}

	if tour.Type == "" {
		return fmt.Errorf("TOUR TYPE CANNOT BE EMPTY")
	}

	if tour.Price <= 0 {
		return fmt.Errorf("PRICE MUST BE GREATER THAN ZERO")
	}

	if tour.Capacity <= 0 {
		return fmt.Errorf("CAPACITY MUST BE GREATER THAN ZERO")
	}

	if tour.CompanyID == "" {
		return fmt.Errorf("COMPANY ID CANNOT BE EMPTY")
	}

	if len(tour.TechnicalTeam) == 0 {
		return fmt.Errorf("AT LEAST ONE TECHNICAL TEAM MEMBER MUST BE ASSIGNED")
	}

	// if !tour.AdminApprove {
	// 	return fmt.Errorf("TOUR MUST BE APPROVED BY AN ADMIN")
	// }

	return nil
}
