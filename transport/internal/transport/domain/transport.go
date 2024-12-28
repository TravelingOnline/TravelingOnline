package domain

type CompanyID string

type Company struct {
	Id    string
	Name  string
	Owner *Owner
}

type Owner struct {
	Id        uint64
	FirstName string
	LastName  string
	Email     string
}
