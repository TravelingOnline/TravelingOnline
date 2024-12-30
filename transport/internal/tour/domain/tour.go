package domain

type TourID string

type Tour struct {
	Id             string
	Source         string
	Destination    string
	StartDate      string
	EndDate        string
	Type           string
	Price          int32
	VehicleUnicode string
	TechnicalTeam  []*TechnicalTeam
}

type TechnicalTeam struct {
	Id        uint64
	FirstName string
	LastName  string
	Age       int32
	Expertise string
}
