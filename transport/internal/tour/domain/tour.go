package domain

type TourID string

type Tour struct {
	Id           string
	Source       string
	Destination  string
	StartDate    string
	EndDate      string
	Type         string
	Price        int32
	Vehicle      Vehicle
	Capacity     int32
	AdminApprove bool
	Ended        bool
	CompanyID    string
	// Company       *company_domain.Company
	TechnicalTeam []*TechnicalTeam
}

type TechnicalTeam struct {
	Id        string
	FirstName string
	LastName  string
	Age       int32
	Expertise string
}

type Vehicle struct {
	Id              string
	Unicode         string
	RequiredExperts int32
	Speed           int32
	RentPrice       int32
	Type            string
	Passenger       int32
	Model           int32
}
