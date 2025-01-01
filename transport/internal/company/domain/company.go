package domain

type CompanyID string

type Company struct {
	Id    string
	Name  string
	Owner *Owner
}

type Owner struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
}
