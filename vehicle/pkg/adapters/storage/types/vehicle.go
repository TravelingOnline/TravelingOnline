package types

type Vehicle struct {
	Id              string
	Unicode         string
	RequiredExperts int32
	Speed           int32
	RentPrice       int32
	IsActive        bool
	Type            string
	Owner           *Owner
}

type Owner struct {
	Id        uint64
	FirstName string
	LastName  string
	Email     string
}
