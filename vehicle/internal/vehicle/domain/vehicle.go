package domain

type VehicleID string

type Vehicle struct {
	Id              string
	Unicode         string
	RequiredExperts int32
	Speed           int32
	RentPrice       int32
	IsActive        bool
	Type            string
	Owner           *Owner
	Passenger       int
	Model           int
}

type Owner struct {
	Id        uint64
	FirstName string
	LastName  string
	Email     string
}
